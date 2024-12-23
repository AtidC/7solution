package main

import "testing"

func TestDecodeAndFindMin(t *testing.T) {
	tests := []struct {
		encoded  string
		expected string
	}{
		{"LLRR=", "210122"},
		{"==RLL", "000210"},
		{"=LLRR", "221012"},
		{"RRL=R", "012001"},
	}

	for _, test := range tests {
		result := Decode(test.encoded)
		if result != test.expected {
			t.Errorf("For encoded '%s', expected '%s', but got '%s'", test.encoded, test.expected, result)
		}
	}
}
