package maximum_subarray

import "math"

/**
53. 最大子数组和

给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。

*/

/**
贪心解法
*/
func maxSubArray(nums []int) int {
	cur := 0
	max := math.MinInt64

	for i := 0; i < len(nums); i++ {
		if cur < 0 {
			cur = 0
		}
		cur += nums[i]

		if cur > max {
			max = cur
		}
	}

	return max
}

/**
动态规划
*/
func maxSubArrayByDp(nums []int) int {
	max := math.MinInt64
	n := len(nums)
	if n < 1 {
		return 0
	}

	//初始化
	dp := make([]int, n)
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
