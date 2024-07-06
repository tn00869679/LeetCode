package MedianofTwoSortedArrays

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergeSlice := append(nums1, nums2...)
	sortedSlice := quickSort(mergeSlice)

	if len(sortedSlice)%2 != 0 {
		return float64(sortedSlice[(len(sortedSlice) / 2)])
	} else {
		return (float64(sortedSlice[(len(sortedSlice)/2)-1]) + float64(sortedSlice[(len(sortedSlice)/2)])) / 2
	}
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivotIndex := len(arr) / 2
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := 0; i < len(arr); i++ {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}
