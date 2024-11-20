package list

/**
LCR 024. 反转链表
给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。
*/

// 凭记忆容易写出来，但过一段时间就忘记，还是没有真正理解，需要找到其中的关键点
func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    return pre
}
