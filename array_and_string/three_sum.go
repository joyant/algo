package array_and_string

import "sort"

/**
leetcode: 15 三数之和

给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/

/**
是标准的中等难度
因为不能有重复的三元组，所以必须要先排序，然后在遍历的过程中，跳过重复元素（判断重复就简单了，不和前一个相等就好）
定住一个元素，双指针按需移动
*/
func threeSum(nums []int) [][]int {
    n := len(nums)
    var res [][]int
    sort.Ints(nums)
    for i := 0; i < n-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        j, k := i+1, n-1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if sum > 0 {
                k--
            } else if sum < 0 {
                j++
            } else {
                res = append(res, []int{nums[i], nums[j], nums[k]})
                // skip左边重复的元素
                for j < k && nums[j] == nums[j+1] {
                    j++
                }
                // skip右边重复元素
                for k > j && nums[k] == nums[k-1] {
                    k--
                }
                // 跳到不重复的新元素
                j++
                k--
            }
        }
    }
    return res
}
