package list

import "testing"

func Test_deleteNodes(t *testing.T) {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
    head := &ListNode{
        Val:  arr[0],
        Next: nil,
    }
    h := head
    for _, a := range arr[1:] {
        h.Next = &ListNode{
            Val: a,
        }
        h = h.Next
    }
    x := deleteNodes(head, 3, 2)
    t.Log(x)
}
