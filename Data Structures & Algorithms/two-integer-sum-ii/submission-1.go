func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		if numbers[i] + numbers[j] < target {
			i++
		} else if numbers[i] + numbers[j] > target {
			j--
		} else {
			return []int{i+1, j+1}
		}
	}
	return nil
}
