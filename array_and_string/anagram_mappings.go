package array_and_string

/**
760. 找出变位映射

level: easy

给你两个整数数组 nums1 和 nums2，其中 nums2 是 nums1 的一个 变位词 。两个数组都可能包含重复元素。

返回一个下标映射数组 mapping，它将 nums1 映射到 nums2，其中 mapping[i] = j 表示 nums1 中的第 i 个元素出现在 nums2 的第 j 个下标上。如果有多个答案，返回 任意一个 。

数组 a 是数组 b 的一个 变位词 意味着 b 是通过将 a 中元素的顺序随机打乱生成的。

示例 1：

输入：nums1 = [12,28,46,32,50], nums2 = [50,12,32,46,28]
输出：[1,4,3,2,0]
解释：因为 nums1 中的第 0 个元素出现在 nums2[1] 上，所以 mapping[0] = 1，而 nums1 中的第 1 个元素出现在 nums2[4] 上，所以 mapping[1] = 4，以此类推。
示例 2：

输入：nums1 = [84,46], nums2 = [84,46]
输出：[0,1]

*/

// 能读懂题就能写出来
func anagramMappings(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for idx, num := range nums2 {
		m[num] = idx
	}
	res := make([]int, 0, len(nums1))
	for _, num := range nums1 {
		res = append(res, m[num])
	}
	return res
}
