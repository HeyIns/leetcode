func canJump(nums []int) bool {
    length := len(nums)
    index := length - 1
    for i := length - 2; i >= 0; i-- {
        if index - i <= nums[i] {
            index = i
        }
    }
    if index == 0 {
        return true
    }
    return false
}