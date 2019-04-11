func combinationSum2(candidates []int, target int) [][]int {
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
            // if i > num && candidates[i] == candidates[i - 1] {
            //     continue
            // }
            // rec(append(list, candidates[i]), tar - candidates[i], i + 1)
            if i == num || candidates[i] != candidates[i - 1] {
                rec(append(list, candidates[i]), tar - candidates[i], i + 1)
            }
        }
    }
    rec([]int{}, target, 0)
    return res
}
