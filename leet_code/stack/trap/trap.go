package trap

/**
42. 接雨水

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
*/

func trap(height []int) int {
	area := 0
	stack := []int{0}
	for i := 0; i < len(height); i++ {
		for height[i] > height[stack[len(stack)-1]] {
			curHeight := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			//弹出元素后变为空栈，则表示左侧没有元素了，即做边无限小，无法储存雨水
			if len(stack) == 0 {
				break
			}

			h := height[stack[len(stack)-1]]
			if h > height[i] {
				h = height[i]
			}
			area += (h - curHeight) * (i - stack[len(stack)-1] - 1)
		}
		stack = append(stack, i)
	}

	return area
}
