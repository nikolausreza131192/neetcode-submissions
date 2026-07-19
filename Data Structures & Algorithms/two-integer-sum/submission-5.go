func twoSum(nums []int, target int) []int {
    indexMap := make(map[int]int)
    for k, v := range nums {
        indexMap[v] = k
    }

    for k, num := range nums {
        diff := target - num
        if v, found := indexMap[diff]; found && v != k {
            return []int{k, v}
        }
    }
    return nil
}
