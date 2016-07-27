package money

import "errors"

var (
	ErrDiffCurrencies = errors.New("differenct currencies not allowed")
	ErrDivZero        = errors.New("zero not allowed")
	ErrMissingParam   = errors.New("missing parameter")
)
