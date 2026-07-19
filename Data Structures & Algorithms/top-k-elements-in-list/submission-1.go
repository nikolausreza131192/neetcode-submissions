func topKFrequent(nums []int, k int) []int {
    result := make([]int, 0, k)
    hash := make(map[int]int)
    for _, num := range nums {
        hash[num]++
    }

    // create temporary array to store frequency for each number
    tmp := make([][]int, 0, len(hash))
    for num, freq := range hash {
        tmp = append(tmp, []int{freq, num})
    }

    // Sort the array based on the frequencies (index 0)
    sort.Slice(tmp, func(i, j int) bool {
        return tmp[i][0] > tmp[j][0]
    })

    for _, t := range tmp[:k] {
        result = append(result, t[1])
    }

    return result
}
