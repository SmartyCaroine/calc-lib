package main

import (
	"io"
	"os"

	"training/calc-lib/calc"
)

type Calculator interface{ Calculate(a, b int) int }

func main() {
	var (
		inputs     []string   = os.Args[1:]
		calculator Calculator = calc.Addition{}
		output     io.Writer  = os.Stdout
	)
	handler := NewHandler(calculator, output)

	handler.handle(inputs)
}
