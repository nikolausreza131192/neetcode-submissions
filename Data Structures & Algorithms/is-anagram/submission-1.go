func isAnagram(s string, t string) bool {
    if len(s) != len(t) { return false }
    sMap := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        sMap[s[i]]++
    }

    for i := 0; i < len(t); i++ {
        if sMap[t[i]] <= 0 {
            return false
        }
        sMap[t[i]]--
    }
    return true
}
