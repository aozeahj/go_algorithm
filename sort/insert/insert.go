package insert

/**
插入排序：
（1）直接插入排序  稳定
（2）希尔排序		不稳定
*/

func DirectInsert(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
