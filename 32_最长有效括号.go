//用栈检测是否配对，然后遍历找栈中相邻元素最大差值
func longestValidParentheses(s string) int {
    stack := []int{}
    for k, v := range(s) {
        if len(stack) == 0 {
            stack = append(stack, k)
        } else {
            if v == ')' && s[stack[len(stack) - 1]] == '(' {
                stack = stack[:len(stack) - 1]
            } else {
                stack = append(stack, k)
            }
        }
    }
    if len(stack) == 0 {
        return len(s)
    }
    max := len(s) - 1 - stack[len(stack) - 1]
    for i := len(stack) - 1; i > 0; i-- {
        if stack[i] - stack[i - 1] - 1 > max {
            max = stack[i] - stack[i - 1] - 1
        }
    }
    if stack[0] - 0 > max {
        max = stack[0] - 0
    }
    return max
}