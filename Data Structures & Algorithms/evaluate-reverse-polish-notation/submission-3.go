func evalRPN(tokens []string) int {
	i := len(tokens) - 1

	var dfs func() int
	dfs = func() int {
		token := tokens[i]
		i--
		
		if isNumber(token) {
			num, _ := strconv.Atoi(token)
			return num
		}

		rightOperand := dfs()
		leftOperand := dfs()

		switch token {
		case "+":
			return leftOperand + rightOperand
		case "-":
			return leftOperand - rightOperand
		case "*":
			return leftOperand * rightOperand
		default:
			return leftOperand / rightOperand
		}
	}

	return dfs()
}

func isNumber(s string) bool {
	return s != "+" && s != "-" && s != "*" && s != "/"
}