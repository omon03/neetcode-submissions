func isPalindrome(s string) bool {
	runes := []rune{}
	for _, r := range []rune(s) {
		if isLetter(r) {
			runes = append(runes, r)
		}
	}
	if len(runes) <= 1 {
		return true
	}
	l := len(runes) / 2
	if len(runes) % 2 == 1 {
		l++
	}
	i := 0
	j := len(runes) - 1
	fmt.Println("l:", l)
	fmt.Println()

	for i < l && j >= 0 {
		for !isLetter(runes[i]) && i <= l {
			fmt.Printf("1i: %c - %c\n", runes[i], runes[j])
			i++
			fmt.Printf("2i: %c - %c\n", runes[i], runes[j])
		}

		for !isLetter(runes[j]) && (j >= l || j != i) {
			fmt.Printf("1j: %c - %c\n", runes[i], runes[j])
			j--
			fmt.Printf("2j: %c - %c\n", unicode.ToLower(runes[i]), unicode.ToLower(runes[j]))
		}
		if unicode.ToLower(runes[i]) != unicode.ToLower(runes[j]) {
			fmt.Printf("%c != %c\n", unicode.ToLower(runes[i]), unicode.ToLower(runes[j]))
			return false
		}
		fmt.Printf("%c == %c\n", unicode.ToLower(runes[i]), unicode.ToLower(runes[j]))

		j--
		i++

		if i >= j {
			break
		}
	}
	return true
}

func isLetter(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}
