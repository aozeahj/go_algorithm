package move_zero

/**
移动零 283
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
*/

func moveZeroes(nums []int) {
	index := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[index] = nums[j]
			index++
		}
	}

	for j := index; j < len(nums); j++ {
		nums[j] = 0
	}
}
