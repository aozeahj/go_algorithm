package sliding_window_maximun

import "math"

/**
239. 滑动窗口最大值

给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。

示例
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
*/

func maxSlidingWindow(nums []int, k int) []int {
	newNums := append([]int{math.MaxInt64}, nums...)
	maxIndex := 0
	var result []int

	for i := 1; i < len(newNums)-k+1; i++ {
		j := i + k - 1

		if newNums[j] > newNums[maxIndex] {
			result = append(result, newNums[j])
			maxIndex = j
			continue
		}

		if maxIndex >= i {
			result = append(result, newNums[maxIndex])
			continue
		}

		tmpIndex := i
		for l := i; l <= j; l++ {
			if newNums[l] > newNums[tmpIndex] {
				tmpIndex = l
			}
		}
		maxIndex = tmpIndex
		result = append(result, newNums[maxIndex])
	}

	return result
}
