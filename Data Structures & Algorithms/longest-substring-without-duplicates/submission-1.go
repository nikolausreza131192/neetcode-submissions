func lengthOfLongestSubstring(s string) int {
	hash := make(map[string]bool)
	longest := 0
	for left, right := 0, 0; right < len(s); right++ {
		for hash[string(s[right])] {
			delete(hash, string(s[left]))
			left++
		}

		hash[string(s[right])] = true
		currLength := right - left + 1
		if currLength > longest {
			longest = currLength
		}
	}
	return longest
}
