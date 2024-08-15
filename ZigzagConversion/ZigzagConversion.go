package ZigzagConversion

func ZigzagConversion(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	if numRows >= len(s) {
		return s
	}

	explanationStr := ""
	for i := 0; i < numRows; i++ {
		explanationStr += s[i : i+1]
		for j := i + (2 * (numRows - 1)); j-(2*i) < len(s); j += (2 * (numRows - 1)) {
			if i != 0 && i != numRows-1 {
				explanationStr += s[j-(2*i) : j-(2*i)+1]
			}

			if j < len(s) {
				explanationStr += s[j : j+1]
			}
		}
	}

	return explanationStr
}
