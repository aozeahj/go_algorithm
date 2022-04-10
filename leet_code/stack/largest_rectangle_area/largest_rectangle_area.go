package largest_rectangle_area

/**
84. 柱状图中最大的矩形

给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/
func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := []int{0}
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	for i := 1; i < len(heights); i++ {
		for heights[i] < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i - 1 - stack[len(stack)-1]
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}

	return maxArea
}
