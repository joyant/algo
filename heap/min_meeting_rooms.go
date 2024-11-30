package heap

import (
    "container/heap"
    "sort"
)

/**
253. 会议室 II
level: medium

给你一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，返回 所需会议室的最小数量 。

示例 1：

输入：intervals = [[0,30],[5,10],[15,20]]
输出：2
示例 2：

输入：intervals = [[7,10],[2,4]]
输出：1

*/

/**
存储结束时间的最小堆
这个问题可以使用优先队列（也称为堆）解决。我们首先需要将所有的会议按照开始时间排序，然后使用一个优先队列来跟踪当前正在进行的会议。

优先队列中的元素是会议的结束时间，优先队列的顶部元素（最小元素）是当前最早结束的会议。当我们考虑下一个会议时，如果它的开始时间在当前最早结束的会议之后，那么我们不需要一个新的会议室，我们可以在最早结束的会议结束后开始这个新的会议。在这种情况下，我们可以将优先队列的顶部元素（旧会议的结束时间）替换为新会议的结束时间。

如果新会议的开始时间在当前最早结束的会议之前，那么我们需要一个新的会议室。在这种情况下，我们将新会议的结束时间添加到优先队列中。

在处理完所有的会议后，优先队列中的元素数量就是所需的最小会议室数量。
*/

type minHeap []int

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *minHeap) Push(x any) {
    *m = append(*m, x.(int))
}

func (m *minHeap) Pop() any {
    old := *m
    n := len(old)
    last := old[n-1]
    old = old[:n-1]
    *m = old
    return last
}

func minMeetingRooms(intervals [][]int) int {
    n := len(intervals)
    if n < 2 {
        return n
    }
    // 按开始时间排序
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    // 不断更新最小堆
    m := &minHeap{}
    heap.Push(m, intervals[0][1])

    for _, interval := range intervals[1:] {
        if m.Len() > 0 && (*m)[0] <= interval[0] {
            heap.Pop(m)
        }
        heap.Push(m, interval[1])
    }
    return m.Len()
}
