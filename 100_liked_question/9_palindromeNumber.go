func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else if x%10 == 0 {
		return false
	}

	var num int
	tmp := x
	for tmp != 0 {
		num = (num * 10) + tmp%10
		tmp /= 10
	}

	return num == x

}