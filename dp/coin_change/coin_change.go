package coin_change

import "math"

/**
322. 零钱兑换

给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/coin-change
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**

 */
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt64
	}

	for i := 1; i <= amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				if dp[i] >= dp[i-coins[j]]+1 {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}

	if dp[amount] > math.MaxInt64 {
		return -1
	}

	return dp[amount]
}
