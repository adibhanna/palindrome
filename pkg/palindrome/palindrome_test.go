package palindrome

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"babad", "bab"}, // or "aba", as both are correct
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"}, // or "c", as both are correct
		{"racecar", "racecar"},
		{"", ""},
		{"abb", "bb"},
		{"bananas", "anana"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := LongestPalindrome(tc.input)
			if actual != tc.expected {
				t.Errorf("longestPalindrome(%v) = %v; expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}
