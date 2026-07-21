func isValid(s string) bool {
	stack := make([]string, 0)
    for i := 0; i < len(s); i++ {
		current := string(s[i])
		if !isCloseBracket(current) {
			stack = append(stack, current)
			continue
		} else if len(stack) == 0 {
			return false
		}

		openBracket := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !isValidPair(openBracket, current) {
			return false
		}
	}
	return true && len(stack) == 0
}

func isValidPair(left, right string) bool {
	return left == "{" && right == "}" ||
		left == "(" && right == ")" ||
		left == "[" && right == "]"
}

func isCloseBracket(s string) bool {
	return s == ")" || s == "}" || s== "]"
}