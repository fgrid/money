package money_test

import (
	"github.com/fgrid/money"
	"testing"
)

func TestMoneyCents(t *testing.T) {
	m := money.USD(1)

	a := m.Cents()
	e := uint64(1)

	if a != e {
		t.Errorf("unexpected cents value %d. expected %d.", a, e)
	}
}

func TestMoneyString(t *testing.T) {
	m := money.New(100, money.NewCurrency("EUR"))

	a := m.String()
	e := "EUR 1.00"

	if a != e {
		t.Errorf("unexpected string representation %q. expected %q.", a, e)
	}
}

func TestMoneyAdd(t *testing.T) {
	m1 := money.EUR(1)
	m2 := money.EUR(100)

	a := m1.Add(m2).String()
	e := "EUR 1.01"

	if a != e {
		t.Errorf("unexpected sum %q. expected %q.", a, e)
	}
}

func TestMoneyAddDebits(t *testing.T) {
	m1 := money.EUR(1).Debit()

	a := m1.Add(m1).String()
	e := "EUR -0.02"

	if a != e {
		t.Errorf("unexpected sum of debits %q. expected %q.", a, e)
	}
}

func TestMoneySub(t *testing.T) {
	m1 := money.GBP(100)
	m2 := money.GBP(1)

	a1 := m1.Sub(m2).String()
	e1 := "GBP 0.99"

	if a1 != e1 {
		t.Errorf("unexpected difference %q. expected %q.", a1, e1)
	}

	a2 := m2.Sub(m1).String()
	e2 := "GBP -0.99"
	if a2 != e2 {
		t.Errorf("unexpected difference %q. expected %q.", a2, e2)
	}
}

func TestMoneyInv(t *testing.T) {
	m := money.CHF(1)
	m1 := m.Inv()
	m2 := m1.Inv()

	a1 := m1.String()
	a2 := m2.String()
	e1 := "CHF -0.01"
	e2 := "CHF 0.01"

	if a1 != e1 {
		t.Errorf("unexpected inverse %q. expected %q", a1, e1)
	}
	if a2 != e2 {
		t.Errorf("unexpected inverse %q. expected %q", a2, e2)
	}
}

func TestMoneyMul(t *testing.T) {
	m1 := money.JPY(51)

	a := m1.Mul(3).String()
	e := "JPY 1.53"

	if a != e {
		t.Errorf("unexpected product %q. expected %q.", a, e)
	}
}

func TestMoneyPercent(t *testing.T) {
	m1 := money.CNY(100)

	a := m1.Percent(33).String()
	e := "CNY 0.33"

	if a != e {
		t.Errorf("unexpected percentage %q. expected %q.", a, e)
	}
}

func TestMoneyDebit(t *testing.T) {
	m1 := money.AUD(1)
	m2 := money.AUD(2)

	a := m1.Sub(m2).String()
	e := "AUD -0.01"

	if a != e {
		t.Errorf("unexpected difference %q. expected %q.", a, e)
	}
}

func TestMoneyEquals(t *testing.T) {
	m1 := money.INR(101)
	m2 := m1.Add(money.INR(1)).Sub(money.INR(1))

	if !m1.Equals(m2) {
		t.Error("failed equality check")
	}
}

func TestMoneyEqualsDebits(t *testing.T) {
	m1 := money.RUB(101).Debit()
	m2 := money.RUB(101).Credit()

	if m1.Equals(m2) {
		t.Errorf("failed equality check by debits")
	}
}

func TestMoneyDebitZero(t *testing.T) {
	m1 := money.HKD(0).Debit()

	a := m1.String()
	e := "HKD 0.00"

	if a != e {
		t.Errorf("got zero debit %q. expected %q", a, e)
	}
}

func TestMoneyDebitZeroEqualsCreditZero(t *testing.T) {
	c := money.CAD(0)
	d := money.CAD(0).Debit()

	if !c.Equals(d) {
		t.Errorf("credit/debit zero failed equality check")
	}

}
