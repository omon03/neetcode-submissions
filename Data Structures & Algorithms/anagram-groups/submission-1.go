func groupAnagrams(strs []string) [][]string {
	mRez := map[string][]string{}
	// m := map[string]map[rune]int{}
	for _, word := range strs {
		if len(mRez) == 0 {
			mRez[word] = []string{word}
			// fmt.Println("len(mRez) == 0:", mRez)
			continue
		}
		exist := false
		for k, v := range mRez {
			if equel(k, word) {
				mRez[k] = append(v, word)
				// fmt.Println("equel:", mRez)
				exist = true
				break
			}
		}
		if !exist {
			mRez[word] = []string{word}
			// fmt.Println("!exist:", mRez)
		}
	}
	rez := make([][]string, len(mRez))
	i := 0
	for _, v := range mRez {
		rez[i] = v
		i++
		// fmt.Println("rez:", rez)
	}
	return rez
}

func equel(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	mCount1 := count(word1)
	mCount2 := count(word2)

	for ch1, count1 := range mCount1 {
		if count2, ok := mCount2[ch1]; !ok || count1 != count2 {
			return false
		}
	}

	return true
}

func count(word string) map[rune]int {
	m := make(map[rune]int, len([]rune(word)))
	for _, r := range []rune(word) {
		m[r] = m[r] + 1
	}
	return m
}
