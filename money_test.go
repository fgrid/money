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

	m3, _ := m1.Add(m2)
	a := m3.String()
	e := "EUR 1.01"

	if a != e {
		t.Errorf("unexpected sum %q. expected %q.", a, e)
	}
}

func TestMoneyAddDebits(t *testing.T) {
	m1 := money.EUR(1).Debit()

	m2, _ := m1.Add(m1)
	a := m2.String()
	e := "EUR -0.02"

	if a != e {
		t.Errorf("unexpected sum of debits %q. expected %q.", a, e)
	}
}

func TestMoneySub(t *testing.T) {
	m1 := money.GBP(100)
	m2 := money.GBP(1)

	m3, _ := m1.Sub(m2)
	m4, _ := m2.Sub(m1)

	a1 := m3.String()
	e1 := "GBP 0.99"

	if a1 != e1 {
		t.Errorf("unexpected difference %q. expected %q.", a1, e1)
	}

	a2 := m4.String()
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

	m2, _ := m1.Mul(3)
	a := m2.String()
	e := "JPY 1.53"

	if a != e {
		t.Errorf("unexpected product %q. expected %q.", a, e)
	}
}

func TestMoneyPercent(t *testing.T) {
	m1 := money.CNY(100)

	m2, _ := m1.Percent(33)
	a := m2.String()
	e := "CNY 0.33"

	if a != e {
		t.Errorf("unexpected percentage %q. expected %q.", a, e)
	}
}

func TestMoneyRoundingTwoThirds(t *testing.T) {
	m1 := money.EUR(100).Debit()
	tmp, _ := m1.Mul(2)
	m2, _ := tmp.Div(3)

	a := m2.String()
	e := "EUR -0.67"

	if a != e {
		t.Errorf("unexpected rounding of 2/3 of 100: %q. expected %q", a, e)
	}
}

func TestMoneyRoundingUp(t *testing.T) {
	m1 := money.EUR(105)
	m2, _ := m1.Div(10)
	a := m2.String()
	e := "EUR 0.11"
	if a != e {
		t.Errorf("unexpected rounding up of 10.5 cents: %q. expected %q", a, e)
	}

	m3 := money.EUR(106)
	m4, _ := m3.Div(10)
	a = m4.String()
	e = "EUR 0.11"
	if a != e {
		t.Errorf("unexpected rounding up of 10.6 cents: %q. expected %q", a, e)
	}
}

func TestMoneyRoundingDown(t *testing.T) {
	m1 := money.EUR(104)
	m2, _ := m1.Div(10)
	a := m2.String()
	e := "EUR 0.10"
	if a != e {
		t.Errorf("unexpected rounding down of 10.4 cents: %q. expected %q", a, e)
	}

	m3 := money.EUR(103)
	m4, _ := m3.Div(10)
	a = m4.String()
	e = "EUR 0.10"
	if a != e {
		t.Errorf("unexpected rounding down of 10.3 cents: %q. expected %q", a, e)
	}

	m5 := money.EUR(100)
	m6, _ := m5.Div(10)
	a = m6.String()
	e = "EUR 0.10"
	if a != e {
		t.Errorf("unexpected div result of 10.0 cents: %q. expected %q", a, e)
	}
}

func TestMoneyRoundAwayFromZero(t *testing.T) {
	m1 := money.EUR(105).Debit()
	m2, _ := m1.Div(10)
	a := m2.String()
	e := "EUR -0.11"
	if a != e {
		t.Errorf("unexpected round up of a debit amount (away from zero): %q. expected %q", a, e)
	}
}

func TestMoneyDebit(t *testing.T) {
	m1 := money.AUD(1)
	m2 := money.AUD(2)

	m3, _ := m1.Sub(m2)
	a := m3.String()
	e := "AUD -0.01"

	if a != e {
		t.Errorf("unexpected difference %q. expected %q.", a, e)
	}
}

func TestMoneyEquals(t *testing.T) {
	m1 := money.INR(101)
	m2, _ := m1.Add(money.INR(1))
	m3, _ := m2.Sub(money.INR(1))

	if !m1.Equals(m3) {
		t.Error("failed equality check")
	}
}

func TestMoneyEqualsDebits(t *testing.T) {
	m1 := money.RUB(101).Debit()
	m2 := money.RUB(101).Credit()

	if m1.Equals(m2) {
		t.Error("failed equality check by debits")
	}
}

func TestMoneyNilEqualsNil(t *testing.T) {
	var m1, m2 *money.Money
	if !m1.Equals(m2) {
		t.Error("nil equality check failed")
	}
}

func TestMoneyEqualsSelf(t *testing.T) {
	m1 := money.EUR(123)
	if !m1.Equals(m1) {
		t.Error("self equality check failed")
	}
}

func TestMoneyNotEqualsNil(t *testing.T) {
	m1 := money.EUR(123)
	if m1.Equals(nil) {
		t.Error("nil equality check did not fail")
	}
}

func TestMoneyNilNotEquals(t *testing.T) {
	var m1 *money.Money
	if m1.Equals(money.EUR(1)) {
		t.Error("nil equality check did not fail")
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

func TestDivideByZero(t *testing.T) {
	m1 := money.EUR(100)
	_, err := m1.Div(0)
	if err != money.ErrDivZero {
		t.Errorf("divide by zero returned with unexpected error %q. expected %q", err, money.ErrDivZero)
	}
}

func TestAddDiffCurr(t *testing.T) {
	m1 := money.EUR(100)
	m2 := money.USD(100)

	_, err := m1.Add(m2)
	if err != money.ErrDiffCurrencies {
		t.Errorf("adding money of different currencies with unexpected error %q. expected %q", err, money.ErrDiffCurrencies)
	}
}

func TestAddNil(t *testing.T) {
	var m1, m2 *money.Money
	m1 = money.EUR(100)
	if _, err := m1.Add(m2); err != money.ErrMissingParam {
		t.Errorf("add nil with unexpected error %q. expected %q", err, money.ErrMissingParam)
	}

	if _, err := m2.Add(m1); err != money.ErrMissingParam {
		t.Errorf("add to a nil with unexpected error %q. expected %q", err, money.ErrMissingParam)
	}
}

func TestSubNil(t *testing.T) {
	var m1, m2 *money.Money
	m1 = money.EUR(100)

	if _, err := m1.Sub(m2); err != money.ErrMissingParam {
		t.Errorf("subtract nil with unexpected error %q. expected %q", err, money.ErrMissingParam)
	}

	if _, err := m2.Sub(m1); err != money.ErrMissingParam {
		t.Errorf("sub from a nil with unexpected error %q. expected %q", err, money.ErrMissingParam)
	}
}

func TestPercentNil(t *testing.T) {
	var m1 *money.Money

	if _, err := m1.Percent(1); err != money.ErrMissingParam {
		t.Errorf("percentage of nil with unexpected error %q. expected %q", err, money.ErrMissingParam)
	}
}

func TestCentsOfNil(t *testing.T) {
	var m1 *money.Money
	a := m1.Cents()
	e := uint64(0)
	if a != e {
		t.Errorf("unexpected count of cents from a nil money %d. expected %q", a, e)
	}
}

func TestMoneyIsDebit(t *testing.T) {
	var m1 *money.Money

	if m1.IsDebit() {
		t.Error("nil money should not be debit")
	}

	m1 = money.EUR(1)
	if m1.IsDebit() {
		t.Error("credit should not be debit")
	}

	m1 = m1.Debit()
	if !m1.IsDebit() {
		t.Errorf("credit should be debit")
	}
}
