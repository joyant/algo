package list

/**
141. 环形链表
给你一个链表的头节点 head ，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。

如果链表中存在环 ，则返回 true 。 否则，返回 false 。
*/

type ListNode struct {
    Val  int
    Next *ListNode
}

// 快慢指针
// 快慢指针的特点是：初始化时，慢的指向第一个，快的指向下一个，移动的时候，慢的每次一步，快的每次两步，
// 移动过程中，记住判断有没有下一步
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    slow, fast := head, head.Next
    for slow != fast {
        // 只判断 fast 就可以，因为它走的快
        if fast == nil || fast.Next == nil {
            return false
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    return true
}
