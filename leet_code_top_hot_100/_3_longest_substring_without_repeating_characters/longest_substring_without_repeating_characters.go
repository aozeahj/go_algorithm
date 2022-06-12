package _3_longest_substring_without_repeating_characters

/**
3. 无重复字符的最长子串

给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]bool{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := 0, 0
	for i := 0; i < n && rk < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk < n && !m[s[rk]] {
			// 不断地移动右指针
			m[s[rk]] = true
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func test() (ret int) {
	i := 100
	defer func(a int) {
		ret = a + 100
	}(i)

	return 300
}
