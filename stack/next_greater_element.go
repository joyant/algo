package stack

/**
496. 下一个更大元素 I
level: easy
nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。

给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。

对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。

返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。

示例 1：
输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
输出：[-1,3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
- 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
- 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。

示例 2：
输入：nums1 = [2,4], nums2 = [1,2,3,4].
输出：[3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
- 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。
*/

/**
预处理 nums2：我们可以使用一个单调栈来预处理 nums2，找到每个元素的下一个更大元素。单调栈是一种特殊的栈结构，它可以帮助我们高效地找到下一个更大元素。

构建映射关系：在预处理 nums2 时，我们可以使用一个哈希表（map）来存储每个元素及其下一个更大元素的映射关系。

查询结果：对于 nums1 中的每个元素，我们只需要在哈希表中查找其对应的下一个更大元素即可。
**/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    nextGreater := map[int]int{}
    var stack []int
    for _, num := range nums2 {
        for len(stack) > 0 && stack[len(stack)-1] < num {
            nextGreater[stack[len(stack)-1]] = num
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, num)
    }
    for _, s := range stack {
        nextGreater[s] = -1
    }
    result := make([]int, 0, len(nums1))
    for _, num := range nums1 {
        result = append(result, nextGreater[num])
    }
    return result
}
