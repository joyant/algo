package list

/**
143. 重排链表
level: medium

给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

// 可以用三个点来总结题干
// 1. 找中间接点
// 2. 返回后半部分链表
// 3. 合并链表，但不是一般题目的那种合并，而是你一个我一个的那种合并
// 据说这是经常出现的经典面试题，可能是因为这一个题包含了3个知识点，这3个点都是链表的常见操作
// 所以可以说一题顶三道题
func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }
    // 1. 找中间节点
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // 现在 slow 是前半部分的最后一个节点，slow.Next 指向下半部分的第一个节点
    lowerPart := slow.Next
    slow.Next = nil
    // 反转后半链表
    var prev *ListNode
    cur := lowerPart
    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    // 反转之后的链表第一个节点是 prev
    // 合并
    first, second := head, prev
    for second != nil {
        fNext := first.Next
        sNext := second.Next
        first.Next = second
        second.Next = fNext
        first = fNext
        second = sNext
    }
}
