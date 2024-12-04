package binary

/**
173. 点名
level: easy
某班级 n 位同学的学号为 0 ~ n-1。点名结果记录于升序数组 records。假定仅有一位同学缺席，请返回他的学号。

示例 1:

输入: records = [0,1,2,3,5]
输出: 4
示例 2:

输入: records = [0, 1, 2, 3, 4, 5, 6, 8]
输出: 7
*/

// 这道题的解法有很多，但二分法可能还是出题人最希望的意图
func takeAttendance(records []int) int {
    i, j := 0, len(records)-1
    for i <= j {
        m := (i + j) / 2 // 找中间 index 的公式
        if records[m] == m {
            i = m + 1
        } else {
            j = m - 1
        }
    }
    return i
}
