package money

// AUD creates an australian dollar money object with given value in cents.
func AUD(cents uint64) *Money {
	return &Money{subs: cents, cur: currencies["AUD"]}
}

// CAD creates a canadian dollar money object with given value in cents.
func CAD(cents uint64) *Money {
	return &Money{subs: cents, cur: currencies["CAD"]}
}

func CHF(centimes uint64) *Money {
	return &Money{subs: centimes, cur: currencies["CHF"]}
}

func CNY(分 uint64) *Money {
	return &Money{subs: 分, cur: currencies["CNY"]}
}

// EUR creates an Euro money object with given value in cents
func EUR(cents uint64) *Money {
	return &Money{subs: cents, cur: currencies["EUR"]}
}

func GBP(pennies uint64) *Money {
	return &Money{subs: pennies, cur: currencies["GBP"]}
}

func HKD(cents uint64) *Money {
	return &Money{subs: cents, cur: currencies["HKD"]}
}

func JPY(sen uint64) *Money {
	return &Money{subs: sen, cur: currencies["JPY"]}
}

func USD(cents uint64) *Money {
	return &Money{subs: cents, cur: currencies["USD"]}
}

func INR(paisa uint64) *Money {
	return &Money{subs: paisa, cur: currencies["INR"]}
}

func RUB(копейка uint64) *Money {
	return &Money{subs: копейка, cur: currencies["RUB"]}
}
