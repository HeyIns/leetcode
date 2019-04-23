func spiralOrder(matrix [][]int) []int {
    res := []int{}
    for len(matrix) > 0 && len(matrix[0]) > 0 {
        res = append(res, matrix[0]...)
        matrix = matrix[1:]
        for k, v := range(matrix) {
            res = append(res, v[len(v) - 1])
            matrix[k] = matrix[k][:len(v) - 1]
        }
        reverse(matrix)
        for k, _ := range(matrix) {
            reverse(matrix[k])
        }
    }
    return res
}
func reverse(itf interface{}) {
    if s, ok := itf.([][]int); ok {
        for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
            s[i], s[j] = s[j], s[i]
        }
    } else if s, ok := itf.([]int); ok {
        for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
            s[i], s[j] = s[j], s[i]
        }
    }
}