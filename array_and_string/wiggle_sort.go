package array_and_string

import "sort"

/**
给你一个的整数数组 nums, 将该数组重新排序后使 nums[0] <= nums[1] >= nums[2] <= nums[3]...

输入数组总是有一个有效的答案。



示例 1:

输入：nums = [3,5,2,1,6,4]
输出：[3,5,1,6,2,4]
解释：[1,6,2,5,3,4]也是有效的答案
示例 2:

输入：nums = [6,6,5,6,3,8]
输出：[6,6,5,6,3,8]
*/

// 先排序，然后从索引为1的元素开始，第 i 个和第 i+1个元素交换，这样就能造成题干的效果
// 代码量看的话是简单类型，但想不出来的话就完全不会做
func wiggleSort(nums []int) {
    sort.Ints(nums)
    for i := 1; i < len(nums)-1; i += 2 {
        nums[i], nums[i+1] = nums[i+1], nums[i]
    }
}
