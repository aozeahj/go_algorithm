package climb_stairs

/**
70. 爬楼梯

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶

思路：不要纠结于多少种方法，通过观察归纳，本质上这是一个数列
n=1，1
n=2，2
n=3，3
n=4，5


n项 = n-1项 + n-2项
*/

func climbStairs(n int) int {
	if n < 1 {
		return 0
	}

	p, q, r := 0, 0, 1

	for j := 0; j < n; j++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
