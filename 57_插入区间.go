func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {
        return [][]int{newInterval}
    }
    res := [][]int{}
    flag := false
    for i := 0; i < len(intervals); i++ {
        if intervals[i][0] > newInterval[1] && !flag {
            res = append(res, newInterval)
            flag = true
        }
        if intervals[i][1] < newInterval[0] || intervals[i][0] > newInterval[1] {
            res = append(res, intervals[i])
        } else {
            if intervals[i][0] < newInterval[0] {
                newInterval[0] = intervals[i][0]
            }
            if intervals[i][1] > newInterval[1] {
                newInterval[1] = intervals[i][1]
            }
        }
    }
    if !flag {
        res = append(res, newInterval)
    }
    return res
}