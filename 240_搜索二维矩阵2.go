func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    i, j := 0, len(matrix[0]) - 1
    for i < len(matrix) && j > -1 {
        if matrix[i][j] < target {
            i++
        } else if matrix[i][j] > target {
            j--
        } else {
            return true
        }
    }
    return false
}