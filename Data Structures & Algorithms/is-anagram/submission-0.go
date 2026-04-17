func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mS := toRunes(s)
	mT := toRunes(t)

	for k, v := range mS {
		if v != mT[k] {
			return false
		}
	}

	return true
}

func toRunes(s string) map[rune]int {
	rez := make(map[rune]int, len([]rune(s)))
	for _, ch := range []rune(s) {
		rez[ch] = rez[ch] + 1
	}
	return rez
}
