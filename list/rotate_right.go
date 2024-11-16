package list

/**
61. 旋转链表
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
*/

/**
1. 使得链表变成一个圈，计算长度
2. 找到这个圈的新头尾
3. 断开连接
*/
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {
        return head
    }
    n := 1
    tail := head
    for tail.Next != nil {
        tail = tail.Next
        n++
    }
    // 变成有一个圈
    tail.Next = head
    // 找新的 head 和 tail
    step := n - k%n
    newTail := tail
    for i := 0; i < step; i++ {
        newTail = newTail.Next
    }
    // 举个例子：1 -> 2 -> 3 -> 4，如果 k 的值是 2
    // 此时 newTail 就是 2
    // 一旦涉及到链表移动的问题，要不先遍历一遍计算长度
    // 要不就是两个指针移动，一个先移动 k，然后两个一起移动，移动 k 的移动到最后，另一个移动到当前
    newHead := newTail.Next
    newTail.Next = nil
    return newHead
}
