func search(nums []int, target int) int {
	i := 0 // 0
	j := len(nums) - 1 // 1

	for {
		k := i + ((j - i) / 2) // 0
		if nums[k] == target {
			return k
		}
		if i >= j { // 0 < 1
			break
		}
		if nums[k] > target { // 2 < 5
			j = k
		} else {
			i = k // <=
		}
		if j - i <= 1 {
			if nums[i] == target {
				return i
			}
			if nums[j] == target {
				return j
			}
			break
		}
	}
	return -1
}
