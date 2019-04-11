func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}
    sort.Ints(candidates)
    var rec func(list []int, tar, num int)
    rec = func(list []int, tar, num int) {
        if tar == 0 {
            l := make([]int, len(list))
            copy(l, list)
            res = append(res, l)
        }
        for i := num; i < len(candidates) && candidates[i] <= tar; i++ {
            rec(append(list, candidates[i]), tar - candidates[i], i)
        }
    }
    rec([]int{}, target, 0)
    return res
}
