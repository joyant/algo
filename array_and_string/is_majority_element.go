package array_and_string

/**
1150. 检查一个数是否在数组中占绝大多数
level: easy

给出一个按 非递减 顺序排列的数组 nums，和一个目标数值 target。假如数组 nums 中绝大多数元素的数值都等于 target，则返回 True，否则请返回 False。

所谓占绝大多数，是指在长度为 N 的数组中出现必须 超过 N/2 次。
*/

// 这道题不是在耍我吧，千万不要想复杂了
func isMajorityElement(nums []int, target int) bool {
    count := 0
    for _, num := range nums {
        if num == target {
            count++
        }
    }
    return count > len(nums)/2
}
