package test

import (
	"LeetCode/HappyNumber"
	"testing"
)

func TestHappyNumber(t *testing.T) {
	example1 := 19
	expected1 := true
	example2 := 2
	expected2 := false
	example3 := 7
	expected3 := true

	result1 := HappyNumber.IsHappy(example1)
	if result1 != expected1 {
		t.Errorf("HappyNumber(example1) = %t; want %t", result1, expected1)
	}

	result2 := HappyNumber.IsHappy(example2)
	if result2 != expected2 {
		t.Errorf("HappyNumber(example2) = %t; want %t", result2, expected2)
	}

	result3 := HappyNumber.IsHappy(example3)
	if result3 != expected3 {
		t.Errorf("HappyNumber(example3) = %t; want %t", result3, expected3)
	}
}
