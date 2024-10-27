package main

import (
	"testing"
)

// Функция для сравнения вычисленного значения с ожидаемым (с допуском для float64)
func floatEquals(a, b, tolerance float64) bool {
	return (a-b) < tolerance && (b-a) < tolerance
}

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		shouldFail bool
	}{
		// Базовые арифметические операции
		{"3 + 4", 7, false},
		{"10 - 3", 7, false},
		{"6 * 3", 18, false},
		{"8 / 4", 2, false},

		// Проверка приоритета операций
		{"3 + 4 * 2", 11, false},
		{"(3 + 4) * 2", 14, false},
		{"10 - 2 * 3", 4, false},
		{"(10 - 2) * 3", 24, false},

		// Проверка на деление на ноль
		{"10 / 0", 0, true},

		// Сложные выражения
		{"(3 + 4) * 2 / (1 - 5) * 2", -7, false},
		{"5 * (6 + 2) - 8 / 4", 38, false},

		// Неверные выражения
		{"1+1*", 0, true},
		{"3 + (4 * 2", 0, true},
		{"(3 + 4) * 2)", 0, true},
		{"abc + 4", 0, true},
		{"", 0, true},
	}

	for _, tt := range tests {
		result, err := Calc(tt.expression)
		if tt.shouldFail {
			if err == nil {
				t.Errorf("ожидалась ошибка для выражения %q, но её не произошло", tt.expression)
			}
		} else {
			if err != nil {
				t.Errorf("не ожидалось ошибки для выражения %q, но получила ошибку %v", tt.expression, err)
			} else if !floatEquals(result, tt.expected, 1e-9) {
				t.Errorf("выражение %q: ожидалось %v, но получено %v", tt.expression, tt.expected, result)
			}
		}
	}
}
