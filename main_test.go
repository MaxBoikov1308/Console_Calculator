package main

import (
	"errors"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		err      error
	}{
		{},
		{},
		{},
		{},
		{},
		{},
	}

	for _, test := range tests {
		got, err := Calc(test.input)
		if got != test.expected || !errors.Is(err, test.err) {
			t.Errorf("Calc(%q) = (%b, %v); want (%b, %v)", test.input, got, err, test.expected, test.err)
		}
	}
}
