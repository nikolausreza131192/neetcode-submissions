func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, 0) // store days that haven't found the warmer temperature
	for i := 0; i < len(temperatures); i++ {
		current := temperatures[i]
		for len(stack) > 0 && 
			current > temperatures[stack[len(stack)-1]] {
			// pop index from stack
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// compute days diff from current temp and store to result
			result[index] = i - index
		}

		// push index to stack
		stack = append(stack, i)
	}

	return result
}
