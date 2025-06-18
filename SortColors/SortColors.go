package SortColors

func SortColors(nums []int) {
	low, i := 0, 0
	high := len(nums) - 1
	for i <= high {
		if nums[i] == 0 {
			nums[low], nums[i] = nums[i], nums[low]
			low++
			i++
			continue
		}

		if nums[i] == 2 {
			nums[i], nums[high] = nums[high], nums[i]
			high--
			continue
		}

		i++
	}
}
