package leetcode100

// https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	n := len(strs[0])
	for {
		if n == 0 {
			return ""
		}

		prefix := strs[0][:n]
		flag := true
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) < n || strs[i][:n] != prefix {
				flag = false
				break
			}
		}
		if flag {
			return prefix
		}
		n--
	}
}
