package main

import (
	"bufio"
	"bytes"
	"errors"
	"reflect"
	"testing"

	"training/calc-lib/calc"
)

func assertEqual(t *testing.T, expected, actual any) {
	if !reflect.DeepEqual(expected, actual) {
		t.Helper() // this filters out the unhelpful steps in the stack trace
		t.Errorf("\n Expected: %v "+
			"\n Actual:   %v", expected, actual)
	}
}
func assertError(t *testing.T, expected, actual error) {
	if !errors.Is(actual, expected) {
		t.Helper() // this filters out the unhelpful steps in the stack trace
		t.Errorf("\n Expected: %v "+
			"\n Actual:   %v", expected, actual)
	}
}

func TestTooFewInputs(t *testing.T) {
	output := &bufio.Writer{}
	handler := NewHandler(calc.Addition{}, output)

	err := handler.handle(nil)

	assertError(t, errTooFewArguments, err)
}

func TestMalformedInput(t *testing.T) {
	output := &bufio.Writer{}
	handler := NewHandler(calc.Addition{}, output)

	err := handler.handle([]string{"NaN", "1"})

	assertError(t, errMalformedInput, err)

	err = handler.handle([]string{"1", "NaN"})

	assertError(t, errMalformedInput, err)
}

func TestErringWriter(t *testing.T) {
	smile := errors.New("smile")
	output := &ErringWriter{err: smile}
	handler := NewHandler(calc.Addition{}, output)

	err := handler.handle([]string{"1", "1"})

	assertError(t, errFailToWrite, err)
}

func TestHappyPath(t *testing.T) {
	output := bytes.Buffer{}
	handler := *NewHandler(calc.Addition{}, &output)

	err := handler.handle([]string{"1", "1"})

	assertError(t, nil, err)
}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write(p []byte) (n int, err error) {
	return 0, this.err
}
