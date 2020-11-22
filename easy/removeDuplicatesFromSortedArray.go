// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	cur, nxt := 0, 1
	for nxt < n {
		if nums[nxt] == nums[cur] {
			nxt++
		} else {
			cur++
			nums[cur] = nums[nxt]
		}
	}
	return cur + 1
}