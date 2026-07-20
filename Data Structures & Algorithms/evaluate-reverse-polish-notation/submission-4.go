func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		if isNumber(token) {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
			continue
		}

		// pop for right & left operand
		rightOperand := stack[len(stack)-1]
		leftOperand := stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		// push new value based on the operation result
		switch token {
		case "+":
			num := leftOperand + rightOperand
			stack = append(stack, num)
		case "-":
			num := leftOperand - rightOperand
			stack = append(stack, num)
		case "*":
			num := leftOperand * rightOperand
			stack = append(stack, num)
		case "/":
			num := leftOperand / rightOperand
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func isNumber(s string) bool {
	return s != "+" && s != "-" && s != "*" && s != "/"
}