package pre_order_traver

import "github.com/aozeahj/go_algorithm/leet_code/tree"

/**
二叉树前序遍历
*/

func preOrderTraver(root *tree.BinaryTreeNode) []int {
	var ret []int
	if root.Left != nil {
		ret = append(ret, preOrderTraver(root.Left)...)
	}

	ret = append(ret, root.Val)

	if root.Right != nil {
		ret = append(ret, preOrderTraver(root.Right)...)
	}

	return ret
}
