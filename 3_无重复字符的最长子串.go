func lengthOfLongestSubstring(s string) int {
    if s == "" {
        return 0
    }
    max_length := 0
    maps := map[byte]int{}
    for i, j := 0, 0; j < len(s); j++ {
        if v, ok := maps[s[j]]; ok {
            if i < v {
                i = v
            }
        }
        if j + 1 - i > max_length {
            max_length = j + 1 - i
        }
        maps[s[j]] = j + 1
    }
    return max_length
}
