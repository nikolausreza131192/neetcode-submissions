func search(nums []int, target int) int {
	if len(nums) == 1 && target == nums[0] {
		return 0
	}
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i] < target {
			i++
		}
		if nums[j] > target {
			j--
		}
		if nums[i] == target {
			return i
		}
		if nums[j] == target {
			return j
		}
	}
	return -1
}
