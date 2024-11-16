package list

/**
82. 删除排序链表中的重复元素 II
给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
*/

// 注意这里的要求是把所有重复的都删除，而不是重复的元素只留下一个
func deleteDuplicates(head *ListNode) *ListNode {
    dummy := &ListNode{Next: head}
    cur := dummy
    for cur.Next != nil && cur.Next.Next != nil {
        if cur.Next.Val == cur.Next.Next.Val {
            val := cur.Next.Val
            // 像消消乐一样把重复的元素都消掉
            for cur.Next != nil && cur.Next.Val == val {
                cur.Next = cur.Next.Next
            }
        } else {
            cur = cur.Next
        }
    }
    return dummy.Next
}
