package _5_zui_chang_hui_wen_zi_chuan

/**
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。

回文子串，从左读和从右读 得到的字符串是一样的
*/

func longestPalindrome(s string) string {
	l := len(s)
	dp := make([][]bool, l)
	maxLen := s[0:1]
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
	}

	for i := 0; i < l; i++ {
		dp[i][i] = true
	}

	for d := 2; d <= l; d++ {
		//i 为左起始位置， j 为右起始位置
		for i := 0; i < l; i++ {
			j := i + d - 1

			if j >= l {
				break
			}

			if j-i == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}

			if dp[i][j] && (j-i+1) > len(maxLen) {
				maxLen = s[i : j+1]
			}
		}
	}

	return maxLen
}
