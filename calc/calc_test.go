package calc

import "testing"

func TestAddition(t *testing.T) {
	var calculator Addition
	assertEqual(t, 2, calculator.Calculate(1, 1))
	assertEqual(t, 0, calculator.Calculate(-1, 1))
	assertEqual(t, -2, calculator.Calculate(-3, 1))
	assertEqual(t, 0, calculator.Calculate(0, 0))
}

func TestSubtraction(t *testing.T) {
	var calculator Subtraction
	assertEqual(t, 0, calculator.Calculate(1, 1))
	assertEqual(t, -2, calculator.Calculate(-1, 1))
	assertEqual(t, -4, calculator.Calculate(-3, 1))
	assertEqual(t, 0, calculator.Calculate(0, 0))
}

func TestDivision(t *testing.T) {
	var calculator Division
	assertEqual(t, 1, calculator.Calculate(1, 1))
	assertEqual(t, -1, calculator.Calculate(-1, 1))
	assertEqual(t, -6, calculator.Calculate(-6, 1))
	assertEqual(t, 0, calculator.Calculate(0, 0))
}

func TestMultiplication(t *testing.T) {
	var calculator Multiplication
	assertEqual(t, 1, calculator.Calculate(1, 1))
	assertEqual(t, -1, calculator.Calculate(-1, 1))
	assertEqual(t, -6, calculator.Calculate(-3, 2))
	assertEqual(t, 0, calculator.Calculate(0, 0))
}

func assertEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
