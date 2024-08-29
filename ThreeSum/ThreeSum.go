package ThreeSum

import "sort"

func ThreeSum(nums []int) [][]int {
	resultSlice := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			addition := nums[i] + nums[left] + nums[right]
			if addition == 0 {
				resultSlice = append(resultSlice, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if addition < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return resultSlice
}
