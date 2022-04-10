package jump_game_2

/**
45. 跳跃游戏 II

给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

*/

/**
逆向推导，最后一笔条最远
*/
func jump(nums []int) int {
	pos := len(nums) - 1
	step := 0

	for pos > 0 {
		for i := 0; i < len(nums)-1; i++ {
			if i+nums[i] >= pos {
				step++
				pos = i
				break
			}
		}
	}
	return step
}

/**
正向推导，每一步都跳最远
*/
func jump2(nums []int) int {
	end := 0
	maxDistance := 0
	step := 0

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > maxDistance {
			maxDistance = i + nums[i]
		}
		if i == end {
			step++
			end = maxDistance
		}
	}

	return step
}
