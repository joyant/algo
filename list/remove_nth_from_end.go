package list

/**
19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/

// 重点是如何找到倒数第 n 个元素
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    a, b := dummy, dummy
    // 这里很关键 i<=n
    for i := 0; i <= n; i++ {
        a = a.Next
    }
    for a != nil {
        a = a.Next
        b = b.Next
    }
    // 这里才真的修改了 dummy 的值
    // 前面对 a 和 b 的赋值都不会修改 dummy
    b.Next = b.Next.Next
    return dummy.Next
}
