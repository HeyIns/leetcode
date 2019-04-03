package main

func twoSum(nums []int, target int) []int {
    maps := map[int]int{}
    for k, v := range nums {
        member := target - v
        if index, ok := maps[member]; ok {
            return []int{index,k}
        }else {
            maps[v] = k
        }
    }
    return nil
}

func main() {
    
}