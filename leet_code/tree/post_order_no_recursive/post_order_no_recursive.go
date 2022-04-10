package post_order_no_recursive

import "github.com/aozeahj/go_algorithm/leet_code/tree"

/**
145. 二叉树的后序遍历

给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
*/

func postOrderTraversal(root *tree.BinaryTreeNode) []int {
	if root == nil {
		return nil
	}

	var prev *tree.BinaryTreeNode
	stack := []*tree.BinaryTreeNode{root}
	ret := []int{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node != nil {
			stack = append(stack, node.Left)
			continue
		}

		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			continue
		}

		node = stack[len(stack)-1]
		if node.Right == nil || node.Right == prev {
			ret = append(ret, node.Val)
			stack = stack[:len(stack)-1]
			prev = node
			stack = append(stack, nil)
		} else {
			stack = append(stack, node.Right)
		}
	}

	return ret
}
