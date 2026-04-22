func lengthOfLongestSubstring(s string) int {
	rez := 0
	buf := 0
	r := []rune(s)
	m := make(map[rune]int, len(r))

	for i := 0; i < len(r); i++ {
		for j := i; j < len(r); j++ {
			v := r[j]
			if _, ok := m[v]; !ok {
				m[v] = j
				buf++
				if buf > rez {
					rez = buf
				}
			} else {
				i = m[v]
				m = make(map[rune]int, len(r) - i)
				buf = 0
				break
			}
		}
	}
	return rez
}
// thequickbrownf oxjumpsoverthelazydogthequickbrownfoxjumpsovert
// thequickbro wnfoxjumps overthelazydogthequickbrownfoxjumpsovert
// thequickbrownfo xjumpsoverth elazydogthequickbrownfoxjumpsovert
// thequickbrownfoxjumpsove rthelazydog thequickbrownfoxjumpsovert
// thequickbrownfoxjumpsovert helazydogt hequickbrownfoxjumpsovert
// thequickbrownfoxjumpsoverth elazydogth equickbrownfoxjumpsovert
// thequickbrownfoxjumpsoverthe lazydogthequickbr ownfoxjumpsovert
//             --------------|||