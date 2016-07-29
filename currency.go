package money

type Currency struct {
	code string
	prec uint
}

var currencies = map[string]*Currency{
	"AUD": {code: "AUD", prec: 2},
	"CAD": {code: "CAD", prec: 2},
	"CHF": {code: "CHF", prec: 2},
	"CNY": {code: "CNY", prec: 2},
	"EUR": {code: "EUR", prec: 2},
	"GBP": {code: "GBP", prec: 2},
	"HKD": {code: "HKD", prec: 2},
	"ILS": {code: "ILS", prec: 2},
	"INR": {code: "INR", prec: 2},
	"JPY": {code: "JPY", prec: 2},
	"MRO": {code: "MRO", prec: 1},
	"PYG": {code: "PYG", prec: 0},
	"RUB": {code: "RUB", prec: 2},
	"TND": {code: "TND", prec: 3},
	"USD": {code: "USD", prec: 2},
}

func NewCurrency(code string) *Currency {
	return currencies[code]
}

func (c *Currency) Equals(o *Currency) bool {
	return c.code == o.code
}
