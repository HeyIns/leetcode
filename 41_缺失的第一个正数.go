//遍历一次数组把大于等于1的和小于数组大小的值放到原数组对应位置，然后再遍历一次数组查当前下标是否和值对应，如果不对应那这个下标就是答案，否则遍历完都没出现那么答案就是数组长度加1。
func firstMissingPositive(nums []int) int {
    if len(nums) == 0 {
        return 1
    }
    for i := 0; i < len(nums); i++ {
        for(nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i] - 1]) {
            nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
        }
    }
    fmt.Println(nums)
    for i := 0; i < len(nums); i++ {
        if i + 1 != nums[i] {
            return i + 1
        }
    }
    return len(nums) + 1
}