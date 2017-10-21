package money

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// Min returns the smallest value out of the given entries. Every debt is smaller than any credit.
// nil instances will be ignored.
func Min(m ...*Money) (*Money, error) {
	var result *Money
	for _, o := range m {
		if o == nil {
			continue
		}
		if result == nil {
			result = o
			continue
		}
		gt, err := result.GT(o)
		if err != nil {
			return nil, err
		}
		if gt {
			result = o
		}
	}
	return result, nil
}

// Max returns the money with the biggest value out of the given entries. Every credit is bigger than any debt.
// nil instances will be ignored.
func Max(m ...*Money) (*Money, error) {
	var result *Money
	for _, o := range m {
		if o == nil {
			continue
		}
		if result == nil {
			result = o
			continue
		}
		gt, err := o.GT(result)
		if err != nil {
			return nil, err
		}
		if gt {
			result = o
		}
	}
	return result, nil
}

// Sum calculates the sum of the given amounts given a common currency
func Sum(op ...*Money) (result *Money, err error) {
	if len(op) < 1 {
		err = ErrMissingParam
		return
	}
	for i, o := range op {
		if i == 0 {
			result = o
			continue
		}
		result, err = result.Add(o)
		if err != nil {
			break
		}
	}
	return
}

func Parse(value, cur string) (*Money, error) {
	c := NewCurrency(cur)
	if c == nil {
		return nil, ErrUnknownCurrency
	}
	r := new(big.Rat)
	if _, err := fmt.Sscan(value, r); err != nil {
		return nil, ErrInvalidSyntax
	}
	div := int64(1)
	for i := uint(0); i < c.prec; i++ {
		div = div * 10
	}
	r = r.Mul(r, new(big.Rat).SetFrac64(div, 1))
	s := r.FloatString(0)
	signedSubs, _ := strconv.ParseInt(s, 10, 64)
	if signedSubs >= 0 {
		return &Money{subs: uint64(signedSubs), cur: c}, nil
	}
	m := Money{subs: uint64(-signedSubs), cur: c}
	return m.Debit(), nil
}

func ParseString(value string) (*Money, error) {
	parts := strings.Split(value, " ")
	if len(parts) != 2 {
		return nil, ErrInvalidSyntax
	}
	return Parse(parts[1], parts[0])
}
