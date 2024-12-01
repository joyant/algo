package list

/**
1474. 删除链表 M 个节点之后的 N 个节点
level: easy
给定链表 head 和两个整数 m 和 n. 遍历该链表并按照如下方式删除节点:

开始时以头节点作为当前节点.
保留以当前节点开始的前 m 个节点.
删除接下来的 n 个节点.
重复步骤 2 和 3, 直到到达链表结尾.
在删除了指定结点之后, 返回修改过后的链表的头节点.
*/

// 主要把 countN 函数写好
func deleteNodes(head *ListNode, m int, n int) *ListNode {
    h := head
    for {
        mNode := countN(h, m)
        if mNode == nil {
            break
        }
        nNode := countN(mNode.Next, n)
        if nNode == nil {
            // 这里很关键，可能 mNode 之后不足 n 个，但 Next 也需要置为空
            mNode.Next = nil
            break
        }
        mNode.Next = nNode.Next
        h = nNode.Next // mNode.Next也可以
    }
    return head
}

func countN(head *ListNode, n int) *ListNode {
    if head == nil {
        return nil
    }
    if n <= 0 {
        return head
    }
    count := 1
    for head != nil && count < n {
        count++
        head = head.Next
    }
    return head
}
