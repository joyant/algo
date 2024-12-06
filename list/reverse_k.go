package list

func reverseK(n *ListNode, k int) *ListNode {
	dummy := n
	for k > 2 {
		n = n.Next
		k--
	}
	n.Next = reverse(n.Next)
	return dummy
}

func reverse(n *ListNode) *ListNode {
	var prev *ListNode
	var cur = n
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
