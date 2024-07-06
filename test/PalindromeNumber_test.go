package test

import (
	"LeetCode/PalindromeNumber"
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	example1 := 121
	expected1 := true
	example2 := 10
	expected2 := false

	result1 := PalindromeNumber.IsPalindrome(example1)
	if result1 != expected1 {
		t.Errorf("PalindromeNumber(example1) = %t; want %t", result1, expected1)
	}

	result2 := PalindromeNumber.IsPalindrome(example2)
	if result2 != expected2 {
		t.Errorf("PalindromeNumber(example2) = %t; want %t", result2, expected2)
	}
}
