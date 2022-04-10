package bst_validate

import (
	"github.com/aozeahj/go_algorithm/leet_code/tree"
	"math"
)

/**
98. 验证二叉搜索树

给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/

/**
递归方式
*/
func isValidBST(root *tree.BinaryTreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(root *tree.BinaryTreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return validate(root.Left, lower, root.Val) && validate(root.Right, root.Val, upper)
}

/**
非递归
*/

func isValidBST2(root *tree.BinaryTreeNode) bool {
	stack := []*tree.BinaryTreeNode{}
	max := math.MinInt64
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= max {
			return false
		}
		max = root.Val
		root = root.Left
	}

	return true
}
