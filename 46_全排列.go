func permute(nums []int) [][]int {
    res := [][]int{}
    sort.Ints(nums)
    var rec func(l []int, index int)
    rec = func(l []int, index int) {
        if index == len(nums) - 1 {
            res = append(res, append([]int{}, l...))
        }
        for i := index; i < len(l); i++ {
            l[i], l[index] = l[index], l[i]
            rec(l, index + 1)
            l[i], l[index] = l[index], l[i]
        }
    }
    rec(nums, 0)
    return res
}