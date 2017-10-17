package money

import "fmt"

type Money struct {
	cur   *Currency
	subs  uint64
	debit bool
}

func New(subs uint64, cur *Currency) *Money {
	return &Money{subs: subs, cur: cur}
}

func (m *Money) Add(o *Money) (*Money, error) {
	if o == nil || m == nil {
		return nil, ErrMissingParam
	}
	if !m.cur.Equals(o.cur) {
		return nil, ErrDiffCurrencies
	}
	if m.debit == o.debit {
		return &Money{subs: m.subs + o.subs, cur: m.cur, debit: m.debit}, nil
	}
	if m.subs > o.subs {
		return &Money{subs: m.subs - o.subs, cur: m.cur, debit: m.debit}, nil
	}
	return &Money{subs: o.subs - m.subs, cur: m.cur, debit: o.debit}, nil
}

func (m *Money) Subunits() uint64 {
	if m == nil {
		return uint64(0)
	}
	return m.subs
}

func (m *Money) Cents() uint64 {
	if m == nil {
		return uint64(0)
	}
	r := m.subs
	d := int(m.cur.prec) - 2
	if d > 0 {
		for i := 0; i < d; i++ {
			r = r / 10
		}
	}
	if d < 0 {
		for i := 0; i > d; i-- {
			r = r * 10
		}
	}
	return r
}

func (m *Money) Credit() *Money {
	return &Money{subs: m.subs, cur: m.cur, debit: false}
}

func (m *Money) Debit() *Money {
	return &Money{subs: m.subs, cur: m.cur, debit: true}
}

func (m *Money) IsDebit() bool {
	return m != nil && m.debit
}

func (m *Money) Div(o uint64) (*Money, error) {
	if o == 0 {
		return nil, ErrDivZero
	}
	v := (m.subs * 10) / o
	q := v / 10
	r := v % 10
	if r < 5 {
		v = q
	} else {
		v = q + 1
	}
	return &Money{subs: v, cur: m.cur, debit: m.debit}, nil
}

func (m *Money) Equals(o *Money) bool {
	if m == o {
		return true
	}
	if m == nil || o == nil {
		return false
	}
	return m.cur.Equals(o.cur) &&
		m.subs == o.subs &&
		(m.debit == o.debit || m.subs == 0)
}

func (m *Money) Mul(o uint64) (*Money, error) {
	return &Money{subs: m.subs * o, cur: m.cur, debit: m.debit}, nil
}

// MulFloat64 multiplies the amount with the given float64
func (m *Money) MulFloat64(f float64) *Money {
	var newDebit bool
	if f < 0 {
		newDebit = !m.debit
		f *= -1
	} else {
		newDebit = m.debit
	}
	return &Money{subs: uint64(float64(m.subs) * f), cur: m.cur, debit: newDebit}
}

func (m *Money) Inv() *Money {
	return &Money{subs: m.subs, cur: m.cur, debit: !m.debit}
}

func (m *Money) Percent(p uint64) (*Money, error) {
	if m == nil {
		return nil, ErrMissingParam
	}
	tmp, _ := m.Mul(p)
	return tmp.Div(100)
}

func (m *Money) String() string {
	s := ""
	if m.debit && m.subs != 0 {
		s = "-"
	}
	return fmt.Sprintf("%s %s%s", m.CurrencyCode(), s, m.Value())
}

func (m *Money) Sub(o *Money) (*Money, error) {
	if m == nil || o == nil {
		return nil, ErrMissingParam
	}
	return m.Add(o.Inv())
}

// Value returns the absolute amount in readable form
func (m *Money) Value() string {
	div := uint64(1)
	for i := uint(0); i < m.cur.prec; i++ {
		div = div * 10
	}
	q := m.subs / div
	r := m.subs % div
	result := fmt.Sprintf("%d", q)
	if m.cur.prec > 0 {
		result += fmt.Sprintf(".%0[1]*[2]d", m.cur.prec, r)
	}
	return result
}

// CurrencyCode returns the currency code
func (m *Money) CurrencyCode() string {
	return m.cur.code
}

// GT greater than - returns true if the money is greater than the given parameter value
func (m *Money) GT(o *Money) (bool, error) {
	if m == nil || o == nil {
		return false, ErrMissingParam
	}
	if m.CurrencyCode() != o.CurrencyCode() {
		return false, ErrDiffCurrencies
	}
	if m.debit {
		return o.debit && m.subs < o.subs, nil
	}
	return o.debit || m.subs > o.subs, nil
}

// LE less or equal - returns true if the money is less or equal to the given parameter value
func (m *Money) LE(o *Money) (bool, error) {
	result, err := m.GT(o)
	return !result, err
}

// LT less than - returns true if the money is less than the given parameter value
func (m *Money) LT(o *Money) (bool, error) {
	if m == nil || o == nil {
		return false, ErrMissingParam
	}
	if m.CurrencyCode() != o.CurrencyCode() {
		return false, ErrDiffCurrencies
	}
	if m.debit {
		return !o.debit || m.subs > o.subs, nil
	}
	return !o.debit && m.subs < o.subs, nil
}

// GE greater or equal - returns tru if the money is greater or equal to the given parameter value
func (m *Money) GE(o *Money) (bool, error) {
	result, err := m.LT(o)
	return !result, err
}
