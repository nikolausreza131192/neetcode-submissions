func maxArea(heights []int) int {
	maxArea := 0
	for i, j := 0, len(heights)-1; i < j; {
		maxHeight := heights[i]
		if heights[i] > heights[j] {
			maxHeight = heights[j]
		}
		area := maxHeight * (j-i)
		if area > maxArea {
			maxArea = area
		}
		if heights[j] > heights[i] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}
