/*
* @Author: Ins
* @Date:   2019-04-01 13:02:43
* @Last Modified by:   Ins
* @Last Modified time: 2019-04-01 16:57:12
*/
package main

func maxProduct(nums []int) int {
    t, max := 1, nums[0]
    for i := 0; i < len(nums); i++ {
        t *= nums[i]
        if t > max {
            max = t
        }
        if nums[i] == 0 {
            t = 1
        }
    }
    t = 1
    for i := len(nums) - 1; i >= 0; i-- {
        t *= nums[i]
        if t > max {
            max = t
        }
        if nums[i] == 0 {
            t = 1
        }
    }
    return max
}
