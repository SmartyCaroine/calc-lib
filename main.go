package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"training/calc-lib/calc"
)

type Calculator interface{ Calculate(a, b int) int }

func main() {
	var (
		inputs     []string   = os.Args[1:]
		calculator Calculator = calc.Addition{}
		output     io.Writer  = os.Stdout
	)

	handle(inputs, calculator, output)
}

func handle(inputs []string, calculator Calculator, output io.Writer) {
	if len(inputs) < 2 {
		panic("Usage : <a>, <b>")
	}

	a, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic(err)
	}

	c := calculator.Calculate(a, b)
	_, err = fmt.Fprintln(output, c)

	if err != nil {
		panic(err)
	}
}
