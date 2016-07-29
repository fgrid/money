package money_test

import (
	"fmt"
	"testing"

	"github.com/fgrid/money"
)

func ExampleNew() {
	m := money.New(123, money.NewCurrency("USD"))
	fmt.Println(m)
	// Output: USD 1.23
}

func ExampleMoney_Subs() {
	fmt.Println(money.EUR(1).Subs())
	// Output: 1
}

func ExampleMoney_Subs_nil() {
	var m *money.Money
	fmt.Println(m.Subs())
	// Output: 0
}

func ExampleMoney_Cents() {
	fmt.Println(money.AUD(1).Cents())
	// Output:
	// 1
}

func ExampleMoney_Cents_noSubunits() {
	fmt.Println(money.New(1, money.NewCurrency("PYG")).Cents())
	// Output: 100
}

func ExampleMoney_Cents_finerSubunits() {
	TND := money.NewCurrency("TND")

	fmt.Println(money.New(1, TND).Cents())
	fmt.Println(money.New(1000, TND).Cents())
	// Output:
	// 0
	// 100
}

func ExampleMoney_String() {
	fmt.Println(money.EUR(1))
	fmt.Println(money.EUR(100))
	fmt.Println(money.EUR(100).Debit())
	// Output:
	// EUR 0.01
	// EUR 1.00
	// EUR -1.00
}

func ExampleMoney_String_subunits() {
	fmt.Println(money.New(123, money.NewCurrency("PYG")))
	fmt.Println(money.New(123, money.NewCurrency("MRO")))
	fmt.Println(money.New(123, money.NewCurrency("ILS")))
	fmt.Println(money.New(123, money.NewCurrency("TND")))
	// Output:
	// PYG 123
	// MRO 12.3
	// ILS 1.23
	// TND 0.123
}

func ExampleMoney_Add() {
	m1 := money.EUR(1)
	m2 := money.EUR(100)
	m3, _ := m1.Add(m2)
	fmt.Println(m3)
	// Output: EUR 1.01
}

func ExampleMoney_Add_debits() {
	m1 := money.EUR(1).Debit()
	m2, _ := m1.Add(m1)
	fmt.Println(m2)
	// Output: EUR -0.02
}

func ExampleMoney_Add_failWithDifferentCurrencies() {
	m1 := money.EUR(1)
	m2 := money.USD(1)
	_, err := m1.Add(m2)
	fmt.Println(err.Error())
	// Output: different currencies not allowed
}

func ExampleMoney_Sub() {
	m1 := money.GBP(1)
	m2 := money.GBP(100)

	m3, _ := m1.Sub(m2)
	m4, _ := m2.Sub(m1)

	fmt.Println(m3)
	fmt.Println(m4)
	// Output:
	// GBP -0.99
	// GBP 0.99
}

func ExampleMoney_Inv() {
	m1 := money.CHF(1)
	m2 := m1.Inv()
	m3 := m2.Inv()

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	// Output:
	// CHF 0.01
	// CHF -0.01
	// CHF 0.01
}

func ExampleMoney_Mul() {
	m1 := money.JPY(51)
	m2, _ := m1.Mul(3)

	fmt.Println(m2)
	// Output: JPY 1.53
}

func ExampleMoney_Percent() {
	m1 := money.CNY(100)
	m2, _ := m1.Percent(33)

	fmt.Println(m2)
	// Output: CNY 0.33
}

func ExampleMoney_Div_roundingUp() {
	u1, _ := money.EUR(105).Div(10)
	u2, _ := money.EUR(106).Div(10)

	fmt.Println(u1)
	fmt.Println(u2)
	// Output:
	// EUR 0.11
	// EUR 0.11
}

func ExampleMoney_Div_roundingDown() {
	d1, _ := money.EUR(104).Div(10)
	d2, _ := money.EUR(100).Div(10)

	fmt.Println(d1)
	fmt.Println(d2)
	// Output:
	// EUR 0.10
	// EUR 0.10
}

func ExampleMoney_Div_roundingHalfAwayFromZero() {
	d1, _ := money.EUR(105).Debit().Div(10)
	u2, _ := money.EUR(105).Div(10)

	fmt.Println(d1)
	fmt.Println(u2)
	// Output:
	// EUR -0.11
	// EUR 0.11
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
