package main

func maxSubArray(nums []int) int {
    res, sum := nums[0], 0
    for _, num := range(nums) {
        if sum > 0 {
            sum += num
        } else {
            sum = num
        }
        if sum > res {
            res = sum
        }
    }
    return res
}

func main() {
    
}