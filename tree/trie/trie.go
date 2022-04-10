package trie

/**
leetcode 208 702 692
*/

type Trie struct {
	children map[rune]*Trie
	isEnd    bool
	val      string
}

func NewTrie() (t *Trie) {
	t = &Trie{}
	t.children = make(map[rune]*Trie)
	return
}

func (t *Trie) Insert(word string) {
	node := t
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = NewTrie()
		}
		node = node.children[c]
	}
	node.isEnd = true
	node.val = word
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}

	return true
}

func (t *Trie) StartsWith(prefix string) (words []string) {
	node := t
	for _, c := range prefix {
		if _, ok := node.children[c]; !ok {
			return
		}
		node = node.children[c]
	}

	var queue []*Trie
	queue = append(queue, node)
	for len(queue) > 0 {
		levelNodeNum := len(queue)
		for i := 0; i < levelNodeNum; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			//获取已存储的字符串
			if currentNode.isEnd {
				words = append(words, currentNode.val)
			}

			//子节点入队
			if len(currentNode.children) > 0 {
				childrenNodes := make([]*Trie, 0, len(currentNode.children))
				for _, ch := range currentNode.children {
					childrenNodes = append(childrenNodes, ch)
				}
				queue = append(queue, childrenNodes...)
			}

		}
	}
	return
}
