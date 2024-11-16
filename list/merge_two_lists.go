package list

/**
21. 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/

// 两种解法，一种是递归，一种的遍历，遍历的时候，是组成一个新的链表，不是在原来的基础上修改
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var dummy = &ListNode{}
    pre := dummy
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            pre.Next = list1
            list1 = list1.Next
        } else {
            pre.Next = list2
            list2 = list2.Next
        }
        pre = pre.Next
    }
    if list1 != nil {
        pre.Next = list1
    }
    if list2 != nil {
        pre.Next = list2
    }
    return dummy.Next
}
