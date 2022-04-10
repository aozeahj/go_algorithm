package dispatch_candy_3

/**
135. 分发糖果

n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目。

输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
*/

/**
解法一，两次遍历
*/
func candy(ratings []int) int {

	candy := 0
	left := make([]int, len(ratings))
	left[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	right := 0
	for j := len(ratings) - 1; j >= 0; j-- {
		if j < len(ratings)-1 && ratings[j] > ratings[j+1] {
			right++
		} else {
			right = 1
		}

		max := right
		if max < left[j] {
			max = left[j]
		}
		candy += max
	}

	return candy
}

/**
解法二，常数空间遍历
*/

func candy2(ratings []int) int {
	candy := 1
	prev := 1
	dec := 0
	inc := 1

	for i := 1; i < len(ratings); i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				prev = 1
			} else {
				prev++
			}
			candy += prev
			inc = prev
		} else {
			dec++

			if dec == inc {
				dec++
			}
			candy += dec
			prev = 1
		}
	}

	return candy
}
