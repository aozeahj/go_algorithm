package triangle

import "math"

/**
120. 三角形最小路径和

给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minimumTotal(triangle [][]int) int {
	min := math.MaxInt64
	m := len(triangle)

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		//每一行，最左边的第一个元素值
		dp[i][0] = dp[i-1][0] + triangle[i][0]

		//中间元素
		for j := 1; j < i; j++ {
			if dp[i-1][j-1] > dp[i-1][j] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-1]
			}

			dp[i][j] += triangle[i][j]
		}

		//每一行，最右边最后一个元素
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	for i := 0; i < m; i++ {
		if min > dp[m-1][i] {
			min = dp[m-1][i]
		}
	}

	return min
}
