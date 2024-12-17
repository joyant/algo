package bit

/**
1486. 数组异或操作
level: easy
给你两个整数，n 和 start 。

数组 nums 定义为：nums[i] = start + 2*i（下标从 0 开始）且 n == nums.length 。

请返回 nums 中所有元素按位异或（XOR）后得到的结果。

我自己总是记不住异或，两个bit相同计算结果是0，不同计算结果是1
记住它偏好不同就可以了，不同的结果是1
*/

func xorOperation(n int, start int) int {
	// 任意数与0异或运算的结果都是任意数自己
	var ans int
	for i := 0; i < n; i++ {
		ans ^= start + i*2
	}
	return ans
}
