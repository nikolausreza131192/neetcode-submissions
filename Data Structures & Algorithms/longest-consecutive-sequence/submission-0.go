func longestConsecutive(nums []int) int {
	hash := make(map[int]bool)
	for _, num := range nums {
		hash[num] = true
	}

	longest := 0
	for _, num := range nums {
		// we should start from lowest num
		_, found := hash[num-1]
		if found { continue }

		// check consecutive sequence
		x := num + 1
		length := 1
		for hash[x] {
			x++
			length++
		}
		if length > longest {
			longest = length
		}
	}

	return longest
}
