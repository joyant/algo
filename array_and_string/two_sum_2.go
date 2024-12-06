package array_and_string

/**
leetcode: 167 两数之和 II - 输入有序数组
level: medium
给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列，
请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，
则 1 <= index1 < index2 <= numbers.length 。

以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。

你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。

你所设计的解决方案必须只使用常量级的额外空间。
*/

/**
标记虽然是中等难度，但其实是简单难度
看着好像复杂，但不要忘记返回的结果的长度是 2，也就是说，找到两个数就行了
快慢指针法
因为数据是有序的，所以根据 sum 的结果控制指针的移动
另外注意返回的值，不是索引而是序号，所以需要索引+1
*/

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	i, j := 0, n-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else if sum < target {
			i++
		}
	}
	return []int{-1, -1}
}
