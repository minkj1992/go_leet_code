// https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]
	// range는 index가 초과되더라도 에러를 발생시키지 않는다.
	for _, v := range nums[1:] {
		if v < sum+v {
			sum += v
		} else {
			sum = v
		}

		if max < sum {
			max = sum
		}
	}
	return max
}
