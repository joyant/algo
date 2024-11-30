package array_and_string

import "sort"

/**
252. 会议室
level: easy
给定一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，请你判断一个人是否能够参加这里面的全部会议。

示例 1：

输入：intervals = [[0,30],[5,10],[15,20]]
输出：false
示例 2：

输入：intervals = [[7,10],[2,4]]
输出：true

*/

// 先排序，再判断后面元素的开始时间是否大于前面元素的结束时间
func canAttendMeetings(intervals [][]int) bool {
    n := len(intervals)
    if n < 2 {
        return true
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    for i := 1; i < n; i++ {
        prev := intervals[i-1]
        cur := intervals[i]
        if cur[0] < prev[1] {
            return false
        }
    }
    return true
}
