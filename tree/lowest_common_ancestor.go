package tree

/**
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
*/

/**
 解题思路：

我们可以通过递归的方式，从根节点开始搜索节点 p 和 q。如果当前节点 cur 等于 null，或者等于 p 或者 q，那么就返回当前节点。

如果 cur 不等于 p 或 q，那么就递归地在 cur 的左子树和右子树中分别搜索 p 和 q。如果左子树和右子树都找到了结果，那么说明 p 和 q 分别在 cur 的左子树和右子树中，因此 cur 就是 p 和 q 的最近公共祖先。如果只有左子树找到了结果，那么说明 p 和 q 都在 cur 的左子树中，因此左子树的结果就是 p 和 q 的最近公共祖先。右子树的情况同理。
*/

// 代码里的递归并不好理解，需要画一个二叉树，把第三层的两个相邻节点作为 p和 q，根节点就是 root
// 往下推导几步，就知道该如何递归了，或者能验证以下递归的正确性了
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root.Val == p.Val || root.Val == q.Val {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    if left != nil && right != nil {
        return root
    }
    if left != nil {
        return left
    }
    return right
}
