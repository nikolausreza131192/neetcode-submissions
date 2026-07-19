func isAnagram(s string, t string) bool {
    if len(s) != len(t) { return false }
    sMap := make(map[byte]int)
    tMap := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        sMap[s[i]]++
        tMap[t[i]]++
    }

    for k, _ := range sMap {
        if sMap[k] != tMap[k] {
            return false
        }
    }
    return true
}
