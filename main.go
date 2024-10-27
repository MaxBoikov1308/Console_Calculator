package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

var ErrZeroDiv = errors.New("division by zero")
var ErrParentheses = errors.New("error in parentheses")
var ErrSyntax = errors.New("syntax error")

func Calc(expression string) (float64, error) {
	if expression == "" {
		return 0, ErrSyntax
	}

	postfix, err := infixToPostfix(expression)
	if err != nil {
		return 0, err
	}

	result, err := evalPostfix(postfix)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func infixToPostfix(expression string) ([]string, error) {
	var output []string
	var stack []rune
	precedence := map[rune]int{
		'+': 1, '-': 1,
		'*': 2, '/': 2,
		'(': 0, ')': 0,
	}

	for i := 0; i < len(expression); i++ {
		ch := rune(expression[i])

		switch {
		case unicode.IsDigit(ch) || ch == '.':
			number := string(ch)
			for i+1 < len(expression) && (unicode.IsDigit(rune(expression[i+1])) || rune(expression[i+1]) == '.') {
				i++
				number += string(expression[i])
			}
			output = append(output, number)
		case ch == '(':
			stack = append(stack, ch)
		case ch == ')':
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				output = append(output, string(stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, ErrParentheses
			}
			stack = stack[:len(stack)-1]
		case ch == '+' || ch == '-' || ch == '*' || ch == '/':
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[ch] {
				output = append(output, string(stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, ch)
		default:
			if !unicode.IsSpace(ch) {
				return nil, ErrSyntax
			}
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == '(' {
			return nil, ErrParentheses
		}
		output = append(output, string(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
	}

	return output, nil
}

func evalPostfix(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, ErrSyntax
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, ErrSyntax
				}
				stack = append(stack, a/b)
			default:
				return 0, ErrSyntax
			}
		}
	}

	if len(stack) != 1 {
		return 0, nil
	}

	return stack[0], nil
}

func main() {
	expression := "(3 + 4) * 2 / (1 - 5) * 2"
	result, err := Calc(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}
