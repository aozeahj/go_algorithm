package level_order

import "github.com/aozeahj/go_algorithm/leet_code/tree"

func LevelOrder(root *tree.BinaryTreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	queue := []*tree.BinaryTreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var levelElem []int
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			levelElem = append(levelElem, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, levelElem)
	}

	return ret
}
