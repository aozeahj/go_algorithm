package jump_game

/**
55. 跳跃游戏

给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
*/

/**
解题思路：贪心算法，从数组尾部倒数，如果倒数一个位置能到达，倒数第二个位置能到达倒数第一个位置，倒数第三个位置能到达倒数第二个位置...
*/

func canJump(nums []int) bool {
	pos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i]+i >= pos {
			pos = i
		}
	}

	return pos == 0
}
