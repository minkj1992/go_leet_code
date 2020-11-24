// https://leetcode.com/problems/climbing-stairs/
// nPr = n! / (n-r)!
// 이상하게 35에서 에러가 나서, memo추가하였지만, n부분을 수정해야 한다.
var memo = make(map[int]int)

func getFactorial(i int) int {
	fact := 1
	for j := 2; j <= i; j++ {
		fact *= j
	}
	return fact
}

func getCombination(n, r int) int {
	if r == 0 || n == r {
		return 1
	}

	if num, ok := memo[n]; ok {
		return num
	}

	memo[n] = getFactorial(n-r) / (getFactorial(n-2*r) * getFactorial(r))
	return memo[n]
}

func climbStairs(n int) int {
	var answer int
	if n <= 3 {
		return n
	}
	r := n / 2
	for i := 0; i <= r; i++ {
		answer += getCombination(n, i)
	}
	return answer
}
