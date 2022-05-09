package two_sum

func twoSum(nums []int, target int) []int {

	for i, value1 := range nums {
		for j, value2 := range nums[i+1:] {
			if value1+value2 == target {
				return []int{i, i + j + 1}
			}
		}
	}
	return nil
}
