package list

/**
给定一个链表数组，每个链表都已经按升序排列。

请将所有链表合并到一个升序链表中，返回合并后的链表。



示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]
*/

// 遍历+分治
// 用小函数作为帮助，先实现合并两个
// 然后遍历整个 lists，逐个合并
func mergeKLists(lists []*ListNode) *ListNode {
    var prev *ListNode
    for _, n := range lists {
        prev = mergeTwoLists2(prev, n)
    }
    return prev
}
