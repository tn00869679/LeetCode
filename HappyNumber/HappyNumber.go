package HappyNumber

import (
	"strconv"
	"strings"
)

func IsHappy(n int) bool {
	result := calculateHappy(n)

	if result == 1 {
		return true
	}

	return false
}

func calculateHappy(n int) int {
	totalResult := []int{}
Restart:
	nStr := strconv.Itoa(n)
	strSlice := strings.Split(nStr, "")

	intSlice := []int{}
	resultInt := 0

	for _, str := range strSlice {
		strInt, _ := strconv.Atoi(str)
		intSlice = append(intSlice, strInt)
	}

	for _, num := range intSlice {
		resultInt = resultInt + num*num
	}

	if resultInt != 1 {
		if checkRepeat(totalResult, resultInt) {
			return resultInt
		}

		n = resultInt
		totalResult = append(totalResult, resultInt)
		goto Restart
	}

	return resultInt
}

func checkRepeat(total []int, value int) bool {
	for _, val := range total {
		if val == value {
			return true
		}
	}
	return false
}
