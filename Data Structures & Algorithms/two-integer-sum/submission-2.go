func twoSum(nums []int, target int) []int {
	for i1, v1 := range nums {
		for i2 := i1 + 1; i2 < len(nums); i2++ {
			if v1 + nums[i2] == target {
				return []int{i1, i2}
			}
		}
	}
	return []int{}
}
