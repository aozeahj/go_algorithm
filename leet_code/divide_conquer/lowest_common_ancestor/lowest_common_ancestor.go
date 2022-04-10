package lowest_common_ancestor

import "github.com/aozeahj/go_algorithm/leet_code/tree"

/**
236. 二叉树的最近公共祖先

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
*/

/**
分治解法
*/
func lowestCommonAncestor(root, p, q *tree.BinaryTreeNode) *tree.BinaryTreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	return r
}
