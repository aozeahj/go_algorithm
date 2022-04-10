package combine

/**
77. 组合

给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。

输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

func combine(n int, k int) [][]int {
	var res [][]int

	if n <= 0 || k <= 0 || k > n {
		return res
	}

	var backtrack func(n, k, index int, track []int)
	backtrack = func(n, k, index int, track []int) {
		if len(track) == k {
			tmp := make([]int, k)
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		//减枝操作
		if len(track)+n-index+1 < k {
			return
		}

		for i := index; i <= n; i++ {
			track = append(track, i)
			backtrack(n, k, i+1, track)
			track = track[:len(track)-1]
		}
	}

	backtrack(n, k, 1, []int{})
	return res
}
