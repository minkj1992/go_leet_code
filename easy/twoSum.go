func twoSum(nums []int, target int) []int {
	indexes := make(map[int]int, len(nums))
	for i, v := range nums {
		k := target - v
		if j, ok := indexes[k]; ok {
			return []int{j, i}
		}
		indexes[v] = i
	}
	return nil
}