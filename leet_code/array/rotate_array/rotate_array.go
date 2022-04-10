package rotate_array

/**
189. 轮转数组

给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

*/

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	l := len(nums)
	for i := 0; i < l/2; i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
}
