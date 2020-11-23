func isValid(s string) bool {
	if s == "" {
		return true
	}

	if len(s) == 1 {
		return false
	}

	pair := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]rune, 0, 1)
	for _, ch := range s {
		n := len(stack)
		switch ch {
		case '{', '(', '[':
			stack = append(stack, ch) // push
		case '}', ')', ']':
			if n == 0 || stack[n-1] != pair[ch] {
				return false
			} else {
				stack = stack[:n-1] // pop
			}
		}
	}
	return len(stack) == 0
}