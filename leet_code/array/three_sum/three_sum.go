package three_sum

import "sort"

/**
15. 三数之和

给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

*/

func threeSum(nums []int) [][]int {
	var ret [][]int

	if len(nums) < 3 {
		return ret
	}

	sort.Ints(nums)
	if nums[0] > 0 {
		return ret
	}

	for index := 0; index < len(nums)-2; index++ {
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}

		target := -nums[index]
		l := index + 1
		r := len(nums) - 1

		for l < r {
			if l > index+1 && nums[l] == nums[l-1] {
				l++
				continue
			}

			if r < len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue
			}

			sum := nums[l] + nums[r]
			if sum < target {
				l++
			} else if sum > target {
				r--
			} else {
				ret = append(ret, []int{nums[index], nums[l], nums[r]})
				l++
				r--
			}
		}
	}
	return ret
}
