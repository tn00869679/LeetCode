package GenerateParentheses

type questionInfo struct {
	n              int
	resultSlice    *[]string
	combinationStr string
	left           int
	right          int
}

func GenerateParentheses(n int) []string {
	resultSlice := []string{}

	combinateString(&questionInfo{
		n:              n,
		resultSlice:    &resultSlice,
		combinationStr: "",
		left:           0,
		right:          0,
	})

	return resultSlice
}

func combinateString(info *questionInfo) {
	if len(info.combinationStr) == 2*info.n {
		*info.resultSlice = append(*info.resultSlice, info.combinationStr)
		return
	}

	if info.left < info.n {
		combinateString(&questionInfo{
			n:              info.n,
			resultSlice:    info.resultSlice,
			combinationStr: info.combinationStr + "(",
			left:           info.left + 1,
			right:          info.right,
		})
	}

	if info.right < info.left {
		combinateString(&questionInfo{
			n:              info.n,
			resultSlice:    info.resultSlice,
			combinationStr: info.combinationStr + ")",
			left:           info.left,
			right:          info.right + 1,
		})
	}
}
