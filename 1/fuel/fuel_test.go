package fuel

import (
	"fmt"
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		mass     float64
		expected int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for index, test := range tests {
		testName := fmt.Sprintf("Calculator (%d)", index)
		t.Run(testName, func(t *testing.T) {
			result := Calculator(test.mass)
			if result != test.expected {
				t.Errorf("result: %v, expected %v", result, test.expected)
			}
		})
	}
}
