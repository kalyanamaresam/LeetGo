package two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		value2 := target - nums[i]
		if _, ok := m[value2]; ok {
			return []int{m[value2], i}
		}
		m[nums[i]] = i
	}
	return nil
}
