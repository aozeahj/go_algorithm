package valid_bracket

/**
20. 有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

输入：s = "()[]{}"
输出：true
*/

func isValid(s string) bool {
	l := len(s)
	//是否成对存在
	if l%2 != 0 {
		return false
	}
	bracketMap := map[byte]byte{'(': ')', '{': '}', '[': ']'}
	var stack []byte

	for i := 0; i < l; i++ {
		_, ok := bracketMap[s[i]]
		if ok {
			stack = append(stack, s[i])
			continue
		}

		if len(stack) == 0 {
			return false
		}

		b := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if bracketMap[b] != s[i] {
			return false
		}
	}

	return len(stack) == 0
}
