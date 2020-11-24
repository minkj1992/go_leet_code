var memo = make(map[int]int)

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	if num, ok := memo[n]; ok {
		return num
	}

	memo[n] = self.climbStairs(n-1) + self.climbStairs(n-2)
}



