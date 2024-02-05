package calculator

import "errors"

func Divide(val1 float64, val2 float64) (float64, error) { // exported function since using PascalCase or UpperCase
	if isZero(val2) {
		return 0.0, errors.New("Divisor by zero")
	}
	return val1 / val2, nil
}


func isZero(val float64) bool { // unexported function since using camelcase
	return val == 0.0
}
