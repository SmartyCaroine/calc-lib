package calc

import "testing"

func TestAddition(t *testing.T) {
	var calculator Addition
	assertEqual(t, calculator.Calculate(1, 1), 2)
	assertEqual(t, calculator.Calculate(-1, 1), 0)
	assertEqual(t, calculator.Calculate(-3, 1), -2)
	assertEqual(t, calculator.Calculate(0, 0), 0)
}

func assertEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
