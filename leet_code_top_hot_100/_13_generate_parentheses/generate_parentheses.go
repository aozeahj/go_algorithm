package _13_generate_parentheses

/**

22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
*/

var bracketList = [2]byte{'(', ')'}

func generateParenthesis(n int) []string {

	ret := []string{}
	backtrack([]byte{}, 2*n, &ret)
	return ret
}

func backtrack(path []byte, l int, ret *[]string) {
	if len(path) == l {
		if !validate(path) {
			return
		}

		*ret = append(*ret, string(path))
		//非法校验，合法添加到ret
		return
	}

	for _, ch := range bracketList {
		path = append(path, ch)
		backtrack(path, l, ret)
		path = path[:len(path)-1]
	}
}

func validate(path []byte) bool {
	cnt := 0
	for i := 0; i < len(path); i++ {
		if path[i] == bracketList[0] {
			cnt++
		} else if path[i] == bracketList[1] {
			cnt--
		}

		if cnt < 0 {
			return false
		}
	}

	return cnt == 0
}
