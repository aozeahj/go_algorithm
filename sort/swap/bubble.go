package swap

/**
交换排序：
（1）冒泡排序   稳定
（2）快速排序	  不稳定

*/
func Swap(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
