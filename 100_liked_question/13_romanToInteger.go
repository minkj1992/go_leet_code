func romanToInt(s string) int {
	n := len(s)
	if n < 1 {
		return 0
	}

	var romanMap = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var result = romanMap[s[n-1]]

	// n이 2일 때도 동작
	for i := n - 2; i >= 0; i-- {
		if romanMap[s[i]] < romanMap[s[i+1]] {
			result -= romanMap[s[i]]
		} else {
			result += romanMap[s[i]]
		}
	}
	return result
}