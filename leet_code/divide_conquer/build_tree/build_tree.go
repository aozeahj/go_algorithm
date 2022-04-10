package build_tree

import "github.com/aozeahj/go_algorithm/leet_code/tree"

/**
105. 从前序与中序遍历序列构造二叉树

给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
*/

func buildTree(preorder []int, inorder []int) *tree.BinaryTreeNode {
	if len(preorder) < 1 {
		return nil
	}

	root := new(tree.BinaryTreeNode)
	root.Val = preorder[0]
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}
	left := inorder[:index]
	right := inorder[index+1:]

	root.Left = buildTree(preorder[1:len(left)+1], left)
	root.Right = buildTree(preorder[len(left)+1:], right)
	return root
}
