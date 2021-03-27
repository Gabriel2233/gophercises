package main

import (
	"fmt"
	"testing"
)

func TestNormalize(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"(123)456-7892", "1234567892"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test # %d\n", i), func(t *testing.T) {
			normalized := normalize(tc.input)

			if normalized != tc.want {
				t.Errorf("got %s want %s", normalized, tc.want)
			}
		})
	}
}
