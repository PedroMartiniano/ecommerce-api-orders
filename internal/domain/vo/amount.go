package vo

import "errors"

type Amount struct {
	value float64
}

func NewAmount(value float64) (Amount, error) {

	if value < 0 {
		return Amount{}, errors.New("Amount cannot be negative")
	}
	return Amount{
		value: value,
	}, nil
}