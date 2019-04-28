type list [][]int
func (l list) Len() int {
    return len(l)
}
func (l list) Swap(i, j int) {
    l[i], l[j] = l[j], l[i]
}
func (l list) Less(i, j int) bool {
    return l[i][0] < l[j][0]
}
func merge(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }
    sort.Sort(list(intervals))
    // fmt.Println(intervals)
    res := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] <= res[len(res) - 1][1] {
            if intervals[i][1] > res[len(res) - 1][1] {
                res[len(res) - 1][1] = intervals[i][1]
            }
        } else {
            res = append(res, intervals[i])
        }
    }
    return res
}