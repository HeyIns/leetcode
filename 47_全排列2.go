func permuteUnique(nums []int) [][]int {
    res := [][]int{}
    sort.Ints(nums)
    length := len(nums)
    var rec func(l []int, index int)
    rec = func(l []int, index int) {
        if index == length - 1 {
            res = append(res, append([]int{}, l...))
        } else {
            m := map[int]int{}
            for i := index; i < length; i++ {
                if _,ok := m[l[i]];ok {
                    continue
                }
                m[l[i]] = 1
                l[i], l[index] = l[index], l[i]
                rec(l, index + 1)
                l[i], l[index] = l[index], l[i]
            }
        }
    }
    rec(nums, 0)
    return res
}