package money

import "errors"

var (
	ErrDiffCurrencies  = errors.New("money: different currencies not allowed")
	ErrDivZero         = errors.New("money: zero not allowed")
	ErrInvalidSyntax   = errors.New("money: invalid syntax")
	ErrMissingParam    = errors.New("money: missing parameter")
	ErrUnknownCurrency = errors.New("money: unknown currency")
)
