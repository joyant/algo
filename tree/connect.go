package tree

/**
117. 填充每个节点的下一个右侧节点指针 II

给定一个二叉树：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。

初始状态下，所有 next 指针都被设置为 NULL 。
*/

func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
    queue := []*Node{root}
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            // 此时 queue 的前 size 个元素放的就是当前层的元素，用指针将它们连接起来
            e := queue[0]
            queue = queue[1:]
            // 不是最后一个元素才指向
            if i < size-1 {
                e.Next = queue[0]
            }
            // 将下一层的元素放进 queue, 这一步可以放心，因为 size 是上一层的size，所以新放进去的元素会在
            // 下一个 for len(queue) > 0 遍历
            if e.Left != nil {
                queue = append(queue, e.Left)
            }
            if e.Right != nil {
                queue = append(queue, e.Right)
            }
        }
    }
    return root
}
