package all_paths

import (
	"github.com/aozeahj/go_algorithm/leet_code/tree"
	"strconv"
)

/**
257. 二叉树的所有路径

输入：root = [1,2,3,null,5]
输出：["1->2->5","1->3"]
*/

func binaryTreePaths(root *tree.BinaryTreeNode) []string {
	paths := []string{}
	stack := []*tree.BinaryTreeNode{root}
	pathStack := []string{""}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		path := pathStack[len(pathStack)-1]

		if node != nil {
			path += strconv.Itoa(root.Val) + "->"
			pathStack = append(pathStack, path)
			stack = append(stack, node.Left)
			continue
		}

		stack = stack[:len(stack)-1]
		if len(stack) < 1 {
			continue
		}

		node = stack[len(stack)-1]
		path = pathStack[len(pathStack)-1]

		if node.Left == nil && node.Right == nil {
			path += strconv.Itoa(node.Val)
			paths = append(paths, path)
			pathStack = pathStack[:len(pathStack)-1]
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, node.Right)
	}

	return paths
}
