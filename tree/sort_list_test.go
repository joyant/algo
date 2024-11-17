package tree

import "testing"

func Test_sortList(t *testing.T) {
    head := &ListNode{
        Val: 2,
        Next: &ListNode{
            Val: 1,
            Next: &ListNode{
                Val: 3,
                Next: &ListNode{
                    Val: 4,
                },
            },
        },
    }
    x := sortList(head)
    t.Log(x)
}
