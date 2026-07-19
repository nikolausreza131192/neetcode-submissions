func hasDuplicate(nums []int) bool {
    hash := make(map[int]bool)
    for _, v := range nums {
        if hash[v] {
            return true
        }
        hash[v] = true
    }
    return false
}
