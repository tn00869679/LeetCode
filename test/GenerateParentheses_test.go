package test

import (
	"LeetCode/GenerateParentheses"
	"reflect"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	example1 := 3
	expected1 := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	example2 := 1
	expected2 := []string{"()"}

	result1 := GenerateParentheses.GenerateParentheses(example1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("GenerateParentheses(example1) = %v; want %v", result1, expected1)
	}

	result2 := GenerateParentheses.GenerateParentheses(example2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("GenerateParentheses(example2) = %v; want %v", result2, expected2)
	}
}
