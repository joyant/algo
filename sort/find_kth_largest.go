package sort

/**
215. 数组中的第K个最大元素
level: medium
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

*/

/**
在快速排序中，每执行一次 quickSelect 函数，就会产生一个 index，它的左边都是比它小的值，而它的右边都是比它大的值
我们可以利用这个原理，当 index 的值等于 len(nums)-k 的时候，就找到了第 k 大的数
*/
func findKthLargest(nums []int, k int) int {
    return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, low, high, index int) int {
    m := quickSelectPartition(nums, low, high)
    if m == index {
        return nums[m]
    } else if m > index {
        // m 作为 middle 把小于它的值和大于它的值分开
        // 如果 m 大于 index，说明我们要找的值在m 的左边
        return quickSelect(nums, low, m-1, index)
    } else {
        // 否则在右边
        return quickSelect(nums, m+1, high, index)
    }
}

func quickSelectPartition(nums []int, low, high int) int {
    pivot := nums[high]
    i := low - 1
    for j := low; j <= high; j++ {
        if nums[j] < pivot {
            i++
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    i++
    nums[i], nums[high] = nums[high], nums[i]
    return i
}
