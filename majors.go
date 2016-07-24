package money

func AUD(cents uint64) *Money {
	return &Money{cents: cents, cur: currencies["AUD"]}
}

func CAD(cents uint64) *Money {
	return &Money{cents: cents, cur: currencies["CAD"]}
}

func CHF(centimes uint64) *Money {
	return &Money{cents: centimes, cur: currencies["CHF"]}
}

func CNY(分 uint64) *Money {
	return &Money{cents: 分, cur: currencies["CNY"]}
}

func EUR(cents uint64) *Money {
	return &Money{cents: cents, cur: currencies["EUR"]}
}

func GBP(pennies uint64) *Money {
	return &Money{cents: pennies, cur: currencies["GBP"]}
}

func HKD(cents uint64) *Money {
	return &Money{cents: cents, cur: currencies["HKD"]}
}

func JPY(sen uint64) *Money {
	return &Money{cents: sen, cur: currencies["JPY"]}
}

func USD(cents uint64) *Money {
	return &Money{cents: cents, cur: currencies["USD"]}
}

func INR(paisa uint64) *Money {
	return &Money{cents: paisa, cur: currencies["INR"]}
}

func RUB(копейка uint64) *Money {
	return &Money{cents: копейка, cur: currencies["RUB"]}
}
