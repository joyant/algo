package binary

/**
137. 只出现一次的数字 II

给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法且使用常数级空间来解决此问题。
*/

/**
我们可以使用一个 32位的计数器来统计每一位上1出现的次数，如果某一位上1出现的次数是3的倍数，那么在那个只出现一次的数字的二进制表示中，那一位肯定是0；否则就是1。
因为定义 i 的最大值是 32，所以 ans 应该定义成 int32 类型
这里默认是 32bit 的数，但如果是 64bit 呢？
**/

func singleNumber2(nums []int) int {
    var ans int32
    for i := 0; i < 32; i++ {
        count := 0
        for _, num := range nums {
            if ((num >> i) & 1) == 1 {
                count++
            }
        }
        if count%3 != 0 {
            ans |= 1 << i
        }
    }
    return int(ans)
}
