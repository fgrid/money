package money

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
