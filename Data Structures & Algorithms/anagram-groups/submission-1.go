func groupAnagrams(strs []string) [][]string {
    hash := make(map[string][]string)
    for _, str := range strs {
        sortedStr := sortString(str)
        hash[sortedStr] = append(hash[sortedStr], str)
    }

    result := make([][]string, 0)
    for _, h := range hash {
        result = append(result, h)
    }
    return result
}

func sortString(s string) string {
    runes := []rune(s)
    sort.Slice(runes, func(i, j int) bool {
        return runes[i] < runes[j]
    })

    return string(runes[:len(runes)])
}