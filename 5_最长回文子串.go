func longestPalindrome(s string) string {
    ss := "#"
    for _, v := range s {
        ss += string(v) + "#"
    }
    index := 0; radius := 0
    for i := 0; i < len(ss); i++ {
        j := 1
        for i - j >= 0 && i + j < len(ss) && ss[i - j] == ss[i + j] {
            if j > radius {
                index = i
                radius = j
            }
            j++
        }
    }
    return s[(index - radius) / 2 : (index + radius) / 2]
}
