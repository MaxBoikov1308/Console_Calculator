package main

import (
	"errors"
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int
		err      error
	}{
		{[]byte("hello"), 5, nil},
		{[]byte("こんにちは"), 5, nil},
		{[]byte("你好"), 2, nil},
		{[]byte(""), 0, nil},
		{[]byte{0xFF}, 0, ErrInvalidUTF8},       // Неверная UTF-8 последовательность
		{[]byte{0xE2, 0x82}, 0, ErrInvalidUTF8}, // Неполная последовательность
	}

	for _, test := range tests {
		got, err := GetUTFLength(test.input)
		if got != test.expected || !errors.Is(err, test.err) {
			t.Errorf("GetUTFLength(%q) = (%d, %v); want (%d, %v)", test.input, got, err, test.expected, test.err)
		}
	}
}
