// package calculator_test

// package calculator_test make this test BLACK BOX TESTING -> which only can access exported function
// exported function means variabel / function / struct using PascalCase or Uppercase

package calculator

// package calculator make this test WHITE BOX TESTING -> which can access all function

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// V1
// func TestDivide(t *testing.T) {
// 	expected := 2.0

// 	actual, err := calculator.Divide(10.0, 5.0)

// 	if err != nil {
// 		t.Errorf("expected %.1f, got %.1f caused %s", expected, actual, err.Error())
// 	}
// 	if expected != actual {
// 		t.Errorf("expected %.1f, got %.1f", expected, actual)
// 	}
// }

// func TestDivideNegative(t *testing.T) {
// 	expected := -2.0

// 	actual, err := calculator.Divide(10.0, -5.0)

// 	if err != nil {
// 		t.Errorf("expected %.1f, got %.1f caused %s", expected, actual, err.Error())
// 	}
// 	if expected != actual {
// 		t.Errorf("expected %.1f, got %.1f", expected, actual)
// 	}
// }

// func TestDivideZero(t *testing.T) {
// 	expected := 0.0

// 	actual, err := calculator.Divide(10.0, 0.0)

// 	if expected != actual {
// 		if err.Error() != "Divisor by zero" {
// 			t.Errorf("expected %.1f, got %.1f caused %s", expected, actual, err.Error())
// 		}
// 	}
// }

// t. Fatal vs t.Error
// Rest of the test code (not executed if t.Fatal is called)
// Rest of the test code (executed even if t.Error is called)

// REFACTOR
// V2
type DivideCalculatorStructTest struct {
	name          string
	expected      float64
	expectedError error
	dividend      float64
	divisor       float64
}

var testCasesDivide = []DivideCalculatorStructTest{
	{"TestDivide", 2.0, nil, 10.0, 5.0},
	{"TestDivideNegative", -2.0, nil, 10.0, -5.0},
	{"TestDivideZero", 0.0, errors.New("Divisor by zero"), 10.0, 0.0},
}


func TestDivide(t *testing.T) {
	for _, ts := range testCasesDivide {
		t.Run(ts.name, func(t *testing.T) {
			//testify
			assert := assert.New(t)
			actual, err := Divide(ts.dividend, ts.divisor)

			// if err != nil {
			// 	if err.Error() != ts.expectedError.Error() {
			// 		t.Errorf("expected %.1f, got %.1f expected caused %s but was %s", ts.expected, actual, ts.expectedError.Error(), err.Error())
			// 	}
			// }
			// if ts.expected != actual {
			// 	t.Errorf("expected %.1f, got %.1f", ts.expected, actual)
			// }

			// assert using testify
			assert.Equal(ts.expected, actual)
			assert.Equal(ts.expectedError, err)
		})
	}
}

type IsZeroStructTest struct {
	name     string
	expected bool
	arg      float64
}

var isZeroTestCases = []IsZeroStructTest{
	{"TestIsZero", true, 0.0},
	{"TestIsNotZero", false, 1.0},
}

func TestIsZero(t *testing.T) {
	for _, ts := range isZeroTestCases {
		t.Run(ts.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := isZero(ts.arg)

			assert.Equal(ts.expected, actual)
		})
	}
}
