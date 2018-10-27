package gotest

import (
	"errors"
)

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除數不可為0")
	}

	return a / b, nil
}
