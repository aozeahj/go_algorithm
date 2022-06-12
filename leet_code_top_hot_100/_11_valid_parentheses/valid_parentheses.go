package _11_valid_parentheses

/**
20. 有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

输入：s = "()[]{}"
输出：true

输入：s = "(]"
输出：false

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var bruketMap = map[byte]byte{'(': ')', '{': '}', '[': ']'}

func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if _, ok := bruketMap[ch]; ok {
			stack = append(stack, ch)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		t := stack[(len(stack) - 1)]
		stack = stack[:len(stack)-1]
		if bruketMap[t] != ch {
			return false
		}
	}

	return len(stack) == 0
}
