package generate_parenthesis

/**
22. 括号生成

数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
*/

func generateParenthesis(n int) []string {
	ret := []string{}
	b := make([]byte, n*2)

	BackTrack(&b, &ret, n*2, 0)
	return ret
}

func BackTrack(b *[]byte, str *[]string, n, pos int) {
	if n == pos {
		if Valid(*b) {
			*str = append(*str, string(*b))
		}
		return
	}

	(*b)[pos] = '('
	BackTrack(b, str, n, pos+1)
	(*b)[pos] = ')'
	BackTrack(b, str, n, pos+1)
}

func Valid(str []byte) bool {
	cnt := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			cnt++
		} else {
			cnt--
		}

		if cnt < 0 {
			return false
		}
	}

	return cnt == 0
}
