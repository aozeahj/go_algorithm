package tree_depth

import "github.com/aozeahj/go_algorithm/leet_code/tree"

/**
104. 二叉树的最大深度

给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
*/

/**
分治解法
*/
func maxDepth(root *tree.BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l + 1
	}

	return r + 1
}

/**
广度优先搜索解法
*/
func maxDepth2(root *tree.BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*tree.BinaryTreeNode{root}
	level := 0
	for len(queue) > 0 {
		s := len(queue)
		for s > 0 {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			s--
		}
		level++
	}

	return level
}
