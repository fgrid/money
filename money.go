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

func (m *Money) Add(o *Money) *Money {
	if m.debit == o.debit {
		return &Money{cents: m.cents + o.cents, cur: m.cur, debit: m.debit}
	}
	if m.cents > o.cents {
		return &Money{cents: m.cents - o.cents, cur: m.cur, debit: m.debit}
	}
	return &Money{cents: o.cents - m.cents, cur: m.cur, debit: o.debit}
}

func (m *Money) Cents() uint64 {
	return m.cents
}

func (m *Money) Credit() *Money {
	return &Money{cents: m.cents, cur: m.cur, debit: false}
}

func (m *Money) Debit() *Money {
	return &Money{cents: m.cents, cur: m.cur, debit: true}
}

func (m *Money) Div(o uint64) *Money {
	v := (m.cents * 10) / o
	q := v / 10
	r := v % 10
	if r < 5 {
		v = q
	} else {
		v = q + 1
	}
	return &Money{cents: v, cur: m.cur, debit: m.debit}
}

func (m *Money) Equals(o *Money) bool {
	return m.cur.Equals(o.cur) &&
		m.cents == o.cents &&
		(m.debit == o.debit || m.cents == 0)
}

func (m *Money) Mul(o uint64) *Money {
	return &Money{cents: m.cents * o, cur: m.cur, debit: m.debit}
}

func (m *Money) Inv() *Money {
	return &Money{cents: m.cents, cur: m.cur, debit: !m.debit}
}

func (m *Money) Percent(p uint64) *Money {
	return m.Mul(p).Div(100)
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

func (m *Money) Sub(o *Money) *Money {
	return m.Add(o.Inv())
}
