package binary

import "testing"

func Test_closestValue(t *testing.T) {
    n := &TreeNode{
        Val: 4,
        Left: &TreeNode{
            Val:   2,
            Left:  &TreeNode{Val: 1},
            Right: &TreeNode{Val: 3},
        },
        Right: &TreeNode{Val: 5},
    }
    res := closestValue(n, 3.5)
    t.Log(res)
}
