package list

/**
86. 分隔链表
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。
*/
func partition(head *ListNode, x int) *ListNode {
    small := &ListNode{}
    big := &ListNode{}
    smallHead := small
    bigHead := big
    for head != nil {
        val := head.Val
        if val >= x {
            big.Next = head
            big = big.Next
        } else {
            small.Next = head
            small = small.Next
        }
        head = head.Next
    }
    small.Next = bigHead.Next
    // big的 next 可能是一个比 x 小的值，总之就是不属于 big 这个链表了
    // 可以画个例子来理解
    big.Next = nil
    return smallHead.Next
}
