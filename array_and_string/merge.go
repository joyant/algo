package array_and_string

import "sort"

/**
56. 合并区间

以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。


提示：

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/

// 先排序，按每个元素的第一个元素排序
// 再线性扫描，判断重叠
// 重叠的定义是这样，比如[1, 3]和[2, 6]，分别成为 a 和 b，必须 a[1] > b[0]才算重叠
func merge(intervals [][]int) [][]int {
    n := len(intervals)
    if n == 0 {
        return [][]int{}
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    merged := [][]int{intervals[0]}
    for i := 1; i < n; i++ {
        lastIndex := len(merged) - 1
        // 注意这里必须是：>=
        if merged[lastIndex][1] >= intervals[i][0] {
            if merged[lastIndex][1] < intervals[i][1] {
                merged[lastIndex][1] = intervals[i][1]
            }
        } else {
            merged = append(merged, intervals[i])
        }
    }
    return merged
}
