package reverse_array

/**
旋转数组
*/

func reverse(nums []int) {
	var nLen = len(nums)
	for i := 0; i < nLen/2; i++ {
		nums[i], nums[nLen-i-1] = nums[nLen-i-1], nums[i]
	}
}
