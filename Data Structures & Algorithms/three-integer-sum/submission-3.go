func threeSum(nums []int) [][]int {
	rez := [][]int{}
	if len(nums) < 3 {
		return rez
	}

	m := map[[3]int]int{}
	sort.Ints(nums)

	i := 0
	for i < len(nums) - 2 {
		j := i + 1

		for j < len(nums) - 1 {
			k := j + 1

			for k < len(nums) {
				w := nums[i] + nums[j] + nums[k]
				if w == 0 {
					s := []int{nums[i], nums[j], nums[k]}
					// sort.Ints(s)
					var arr [3]int
					copy(arr[:], s)
					m[arr] = 0
				}
				k++
			}
			j++
		}
		i++
	}

	for k, _ := range m {
		rez = append(rez, k[:])
	}
	
	return rez
}
