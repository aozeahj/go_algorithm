package pre_order_no_recursive

import (
	"github.com/aozeahj/go_algorithm/leet_code/tree"
)

func preOrderTraversal(root *tree.BinaryTreeNode) []int {
	if root == nil {
		return nil
	}

	stack := []*tree.BinaryTreeNode{root}
	ret := []int{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node != nil {
			ret = append(ret, node.Val)
			stack = append(stack, node.Left)
			continue
		}

		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			continue
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		stack = append(stack, node.Right)
	}

	return ret
}
