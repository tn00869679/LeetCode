package test

import (
	"LeetCode/MinimumWindowSubstring"
	"testing"
)

func TestMinimumWindowSubstring(t *testing.T) {
	example1_1 := "ADOBECODEBANC"
	example1_2 := "ABC"
	expected1 := "BANC"

	example2_1 := "a"
	example2_2 := "a"
	expected2 := "a"

	example3_1 := "a"
	example3_2 := "aa"
	expected3 := ""

	result1 := MinimumWindowSubstring.MinWindow(example1_1, example1_2)
	if result1 != expected1 {
		t.Errorf("MinimumWindowSubstring(example1_1, example1_2) = %s; want %s", result1, expected1)
	}

	result2 := MinimumWindowSubstring.MinWindow(example2_1, example2_2)
	if result2 != expected2 {
		t.Errorf("MinimumWindowSubstring(example2_1, example2_2) = %s; want %s", result2, expected2)
	}

	result3 := MinimumWindowSubstring.MinWindow(example3_1, example3_2)
	if result3 != expected3 {
		t.Errorf("MinimumWindowSubstring(example3_1, example3_2) = %s; want %s", result3, expected3)
	}
}
