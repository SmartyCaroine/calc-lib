package calc

type Addition struct{}

func (this Addition) Calculate(a, b int) int {
	return a + b
}

type Subtraction struct{}

func (this Subtraction) Calculate(a, b int) int {
	return a - b
}

type Division struct{}

func (this Division) Calculate(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

type Multiplication struct{}

func (this Multiplication) Calculate(a, b int) int {
	return a * b
}
