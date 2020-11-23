import "math"

func reverse(x int) int {
	var reversed int
	for ; x != 0; x /= 10 {
		reversed *= 10
		reversed += x - ((x / 10) * 10) // ((x/10)*10) = 1의 자리를 제외한 나머지 값
	}

	if math.MaxInt32 < reversed || reversed < math.MinInt32 {
		return 0
	}
	return reversed
}