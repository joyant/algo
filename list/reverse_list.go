package list

/**
LCR 024. 反转链表
给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。
*/

// 凭记忆容易写出来，但过一段时间就忘记，还是没有真正理解，需要找到其中的关键点
/**
本质上是让 cur 不断的指向的 prev，然后 现在的 cur 成了 prev，而 next 成了 cur
*/
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
