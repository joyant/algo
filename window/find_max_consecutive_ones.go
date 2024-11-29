package window

/**
487. 最大连续1的个数 II

给定一个二进制数组 nums ，如果最多可以翻转一个 0 ，则返回数组中连续 1 的最大个数。

示例 1：

输入：nums = [1,0,1,1,0]
输出：4
解释：翻转第一个 0 可以得到最长的连续 1。
     当翻转以后，最大连续 1 的个数为 4。
示例 2:

输入：nums = [1,0,1,1,0,1]
输出：4


提示:

1 <= nums.length <= 105
nums[i] 不是 0 就是 1.
*/

// 滑动窗口
// 最开始我的思路是记录 start 指向0的位置，但发现不好记录，一旦出现这样的时刻
// 该有一个意识，就是思路有问题，果然答案是通过 zeroCount 的增减来实现的
// 这有点像使用一个 map 来辅助计算的时候，有时需要删除 map 里的元素，令代码变得非常精简
// Consecutive 连续的
func findMaxConsecutiveOnes(nums []int) int {
    n := len(nums)
    start, end, res, zeroCount := 0, 0, 0, 0
    for end < n {
        if nums[end] == 0 {
            zeroCount++
        }
        for zeroCount > 1 {
            if nums[start] == 0 {
                zeroCount--
            }
            start++
        }
        res = max(res, end-start+1)
        end++
    }
    return res
}
