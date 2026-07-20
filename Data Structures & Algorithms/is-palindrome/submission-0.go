func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; {
		if !isAlphaNumeric(s[i]) {
			i++
			continue
		}
		if !isAlphaNumeric(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isAlphaNumeric(char byte) bool {
	return char >= '0' && char <= '9' || char >= 'a' && char <= 'z'
}