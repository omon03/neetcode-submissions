func maxProfit(p []int) int {
	min := 0
	max := 0
	rez := 0

	for i, v := range p {
		min = v
		max = v
		j := i + 1
		for j < len(p) {
			if p[j] > max {
				max = p[j]
			}
			j++
		}
		buf := max - min
		if buf > rez {
			rez = buf
		}
	}

	return rez
}
