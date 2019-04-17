func topKFrequent(nums []int, k int) []int {
    res := []int{}
    nums_map := map[int]int{}
    for _, n := range(nums) {
        nums_map[n]++
    }
    
    for key, value := range(nums_map) {
        if len(res) < k {
            //setup
            res = append(res, key)
            i := len(res) - 1
            j := (i + 1) / 2 - 1
            for i > 0 && nums_map[res[i]] < nums_map[res[j]] {
                res[i], res[j] = res[j], res[i]
                i = j
                j = (i + 1) / 2 - 1
            }
        } else {
            if value > nums_map[res[0]] {
                res[0] = key
                //adjust
                i := 0
                for i < len(res) {
                    child := i
                    if 2 * i + 2 < len(res) {
                        if nums_map[res[2 * i + 1]] < nums_map[res[2 * i + 2]] {
                            child = 2 * i + 1
                        } else {
                            child = 2 * i + 2
                        }
                    } else if 2 * i + 1 < len(res) {
                        child = 2 * i + 1
                    }
                    if nums_map[res[i]] > nums_map[res[child]] {
                        res[i], res[child] = res[child], res[i]
                        i = child
                    } else {
                        break
                    }
                }
            }
        }
    }
    return res
}