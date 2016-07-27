package money

import "fmt"

type Money struct {
	cur   *Currency
	cents uint64
	debit bool
}

func New(cents uint64, cur *Currency) *Money {
	return &Money{cents: cents, cur: cur}
}

func (m *Money) Add(o *Money) (*Money, error) {
	if o == nil || m == nil {
		return nil, ErrMissingParam
	}
	if !m.cur.Equals(o.cur) {
		return nil, ErrDiffCurrencies
	}
	if m.debit == o.debit {
		return &Money{cents: m.cents + o.cents, cur: m.cur, debit: m.debit}, nil
	}
	if m.cents > o.cents {
		return &Money{cents: m.cents - o.cents, cur: m.cur, debit: m.debit}, nil
	}
	return &Money{cents: o.cents - m.cents, cur: m.cur, debit: o.debit}, nil
}

func (m *Money) Cents() uint64 {
	if m == nil {
		return uint64(0)
	}
	return m.cents
}

func (m *Money) Credit() *Money {
	return &Money{cents: m.cents, cur: m.cur, debit: false}
}

func (m *Money) Debit() *Money {
	return &Money{cents: m.cents, cur: m.cur, debit: true}
}

func (m *Money) IsDebit() bool {
	return m != nil && m.debit
}

func (m *Money) Div(o uint64) (*Money, error) {
	if o == 0 {
		return nil, ErrDivZero
	}
	v := (m.cents * 10) / o
	q := v / 10
	r := v % 10
	if r < 5 {
		v = q
	} else {
		v = q + 1
	}
	return &Money{cents: v, cur: m.cur, debit: m.debit}, nil
}

func (m *Money) Equals(o *Money) bool {
	if m == o {
		return true
	}
	if m == nil || o == nil {
		return false
	}
	return m.cur.Equals(o.cur) &&
		m.cents == o.cents &&
		(m.debit == o.debit || m.cents == 0)
}

func (m *Money) Mul(o uint64) (*Money, error) {
	return &Money{cents: m.cents * o, cur: m.cur, debit: m.debit}, nil
}

func (m *Money) Inv() *Money {
	return &Money{cents: m.cents, cur: m.cur, debit: !m.debit}
}

func (m *Money) Percent(p uint64) (*Money, error) {
	if m == nil {
		return nil, ErrMissingParam
	}
	tmp, _ := m.Mul(p)
	return tmp.Div(100)
}

func (m *Money) String() string {
	div := uint64(100)
	q := m.cents / div
	r := m.cents % div
	s := ""
	if m.debit && m.cents != 0 {
		s = "-"
	}
	return fmt.Sprintf("%[1]s %[2]s%[3]d.%0[4]*[5]d", m.cur.code, s, q, m.cur.prec, r)
}

func (m *Money) Sub(o *Money) (*Money, error) {
	if m == nil || o == nil {
		return nil, ErrMissingParam
	}
	return m.Add(o.Inv())
}
