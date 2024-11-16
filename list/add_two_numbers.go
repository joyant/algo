package list

/**
2. 两数相加
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var carry = 0
    // dummy 是神来之笔，跟 yinwang.org 学的
    var dummy = &ListNode{}
    var pre = dummy
    // 这里加一个 carry 的条件，就可以不在 for 结束的时候再对 carry 做额外的判断
    for l1 != nil || l2 != nil || carry > 0 {
        a, b := 0, 0
        if l1 != nil {
            a = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            b = l2.Val
            l2 = l2.Next
        }
        sum := a + b + carry
        carry = sum / 10
        sum = sum % 10
        n := &ListNode{Val: sum}
        pre.Next = n
        pre = n
    }
    return dummy.Next
}
