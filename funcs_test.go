package money_test

import (
	"fmt"
	"github.com/fgrid/money"
)

func ExampleMin() {
	var m0, m1, m2, m3, m4 *money.Money
	m0 = money.USD(400).Debit()
	m1 = money.USD(100)
	m2 = nil
	m3 = money.USD(201)
	m4 = money.USD(500).Debit()

	min, _ := money.Min(m0, m1, m2, m3, m4)
	fmt.Printf("min(%s, %s, %s, %s, %s) = %s\n", m0, m1, m2, m3, m4, min)
	// Output:
	// min(USD -4.00, USD 1.00, <nil>, USD 2.01, USD -5.00) = USD -5.00
}

func ExampleMin_fail() {
	m1 := money.USD(200)
	m2 := money.EUR(201)
	if _, err := money.Min(m1, m2); err != nil {
		fmt.Printf("min(%s, %s) = %s\n", m1, m2, err.Error())
	}
	// Output: min(USD 2.00, EUR 2.01) = money: different currencies not allowed
}

func ExampleMax() {
	var m0, m1, m2, m3, m4 *money.Money
	m0 = money.USD(400).Debit()
	m1 = money.USD(100)
	m2 = nil
	m3 = money.USD(201)
	m4 = money.USD(100)

	max, _ := money.Max(m0, m1, m2, m3, m4)
	fmt.Printf("max(%s, %s, %s, %s, %s) = %s\n", m0, m1, m2, m3, m4, max)
	// Output:
	// max(USD -4.00, USD 1.00, <nil>, USD 2.01, USD 1.00) = USD 2.01
}

func ExampleMax_fail() {
	m1 := money.USD(200)
	m2 := money.EUR(201)
	if _, err := money.Max(m1, m2); err != nil {
		fmt.Printf("max(%s, %s) = %s\n", m1, m2, err.Error())
	}
	// Output: max(USD 2.00, EUR 2.01) = money: different currencies not allowed
}
