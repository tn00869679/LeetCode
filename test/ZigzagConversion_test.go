package test

import (
	"LeetCode/ZigzagConversion"
	"testing"
)

func TestZigzagConversion(t *testing.T) {
	example1_1 := "PAYPALISHIRING"
	example1_2 := 3
	expected1 := "PAHNAPLSIIGYIR"
	example2_1 := "PAYPALISHIRING"
	example2_2 := 4
	expected2 := "PINALSIGYAHRPI"

	result1 := ZigzagConversion.ZigzagConversion(example1_1, example1_2)
	if result1 != expected1 {
		t.Errorf("ZigzagConversion(example1_1, example1_2) = %s; want %s", result1, expected1)
	}

	result2 := ZigzagConversion.ZigzagConversion(example2_1, example2_2)
	if result2 != expected2 {
		t.Errorf("ZigzagConversion(example2_1, example2_2) = %s; want %s", result2, expected2)
	}
}
