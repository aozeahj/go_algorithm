package _9_letter_combinations_of_a_phone_number

/**
17. 电话号码的字母组合

给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

var digitsStrMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	ret := []string{}
	if len(digits) == 0 {
		return ret
	}

	backtrack([]byte{}, digits, &ret)
	return ret
}

func backtrack(path []byte, digits string, ret *[]string) {
	if len(digits) == 0 {
		*ret = append(*ret, string(path))
		return
	}

	digit := digits[:1]
	str := digitsStrMap[digit]

	for i := 0; i < len(str); i++ {
		path = append(path, str[i])
		backtrack(path, digits[1:], ret)
		path = path[:len(path)-1]
	}
}
