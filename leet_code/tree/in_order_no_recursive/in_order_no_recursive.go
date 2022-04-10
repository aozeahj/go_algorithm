package in_order_no_recursive

import "github.com/aozeahj/go_algorithm/leet_code/tree"

/**
94. 二叉树的中序遍历
*/

func inOrderTraversal(root *tree.BinaryTreeNode) []int {
	if root == nil {
		return nil
	}

	ret := []int{}
	stack := []*tree.BinaryTreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node != nil {
			stack = append(stack, node.Left)
			continue
		}

		stack = stack[:len(stack)-1]
		if len(stack) < 1 {
			continue
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		stack = append(stack, node.Right)
	}

	return ret
}
