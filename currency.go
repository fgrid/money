package money

type Currency struct {
	code string
	prec uint
}

var currencies = map[string]*Currency{
	"AUD": &Currency{code: "AUD", prec: 2},
	"CAD": &Currency{code: "CAD", prec: 2},
	"CHF": &Currency{code: "CHF", prec: 2},
	"CNY": &Currency{code: "CNY", prec: 2},
	"EUR": &Currency{code: "EUR", prec: 2},
	"GBP": &Currency{code: "GBP", prec: 2},
	"HKD": &Currency{code: "HKD", prec: 2},
	"INR": &Currency{code: "INR", prec: 2},
	"JPY": &Currency{code: "JPY", prec: 2},
	"USD": &Currency{code: "USD", prec: 2},
	"RUB": &Currency{code: "RUB", prec: 2},
}

func NewCurrency(code string) *Currency {
	return currencies[code]
}

func (c *Currency) Equals(o *Currency) bool {
	return c.code == o.code
}
