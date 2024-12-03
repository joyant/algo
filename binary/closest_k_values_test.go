package binary

import "testing"

func Test_closestKValues(t *testing.T) {
    n := &TreeNode{
        Val: 4,
        Left: &TreeNode{
            Val:   2,
            Left:  &TreeNode{Val: 1},
            Right: &TreeNode{Val: 3},
        },
        Right: &TreeNode{Val: 5},
    }
    x := closestKValues(n, 3.714286, 2)
    t.Log(x)
}
