package PalindromeNumber

import (
	"strconv"
	"strings"
)

func IsPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	strSlice := strings.Split(xStr, "")
	left := 0
	right := len(strSlice) - 1

	for left < right {
		strSlice[left], strSlice[right] = strSlice[right], strSlice[left]
		left++
		right--
	}

	tmpStr := strings.Join(strSlice, "")
	if tmpStr != xStr {
		return false
	}

	return true
}
