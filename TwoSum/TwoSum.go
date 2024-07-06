package TwoSum

import (
	"strconv"
	"strings"
)

func TwoSum(nums []int, target int) []int {
	resultArr := []int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				resultArr = append(resultArr, i)
				resultArr = append(resultArr, j)
				break
			}
		}

		if len(resultArr) == 2 {
			break
		}
	}

	return resultArr
}

func GetString(nums []int) string {
	var values []string
	for _, num := range nums {
		values = append(values, strconv.Itoa(num))
	}

	return strings.Join(values, ",")
}
