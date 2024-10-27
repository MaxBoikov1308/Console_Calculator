package main

import (
	"errors"
	"fmt"
)

var ErrDivZero = errors.New("division by zero")

func main() {
	fmt.Println(Calc("1+1"))
}

func Calc(expression string) (float64, error) {
	if len(expression) == 0 {
		return 0, nil
	}
	return 0, nil
}

func Plus(a, b float64) (float64, error) {
	return a + b, nil
}

func Minus(a, b float64) (float64, error) {
	return a - b, nil
}

func Multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivZero
	}
	return a / b, nil
}
