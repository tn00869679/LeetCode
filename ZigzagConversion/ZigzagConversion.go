package ZigzagConversion

import "strings"

func ZigzagConversion(s string, numRows int) string {
	explanationArr := [][]string{}
	explanationStr := ""

	i := 0
	for i < len(s) {
		end := i + numRows
		if end > len(s) {
			end = len(s)
		}

		str := s[i:end]
		strSlice := strings.Split(str, "")
		i += numRows
		explanationArr = append(explanationArr, strSlice)
		if end == len(s) {
			break
		}
		for j := numRows - 1; j > 1; j-- {
			stringsSlice := make([]string, numRows)
			stringsSlice[j-1] = s[i : i+1]
			i += 1
			explanationArr = append(explanationArr, stringsSlice)
			if i > len(s)-1 {
				break
			}
		}
	}

	for n := 0; n < numRows; n++ {
		for _, slice := range explanationArr {
			if n > len(slice)-1 {
				continue
			}
			explanationStr += slice[n]
		}
	}

	return explanationStr
}
