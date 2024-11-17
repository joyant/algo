package tree

/**
148. 排序链表
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
*/

func sortList(head *ListNode) *ListNode {
    // 这里很重要
    if head == nil || head.Next == nil {
        return head
    }
    // 找中间节点，以下是一个技巧
    slow, fast := head, head.Next
    // fast 不为空，slow 一定不空
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    mid := slow.Next
    // 将链表从 mid 断开，形成两个链表
    slow.Next = nil
    // 递归
    left := sortList(head)
    right := sortList(mid)
    // 合并
    return mergeSortList(left, right)
}

func mergeSortList(a, b *ListNode) *ListNode {
    dummy := &ListNode{}
    cur := dummy
    for a != nil && b != nil {
        if a.Val > b.Val {
            cur.Next = b
            b = b.Next
        } else {
            cur.Next = a
            a = a.Next
        }
        cur = cur.Next
    }
    if a != nil {
        cur.Next = a
    }
    if b != nil {
        cur.Next = b
    }
    return dummy.Next
}
