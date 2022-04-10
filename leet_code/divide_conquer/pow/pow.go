package pow

/**
50. Pow(x, n)

实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn ）

输入：x = 2.00000, n = 10
输出：1024.00000
*/

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}

	return 1 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {

	r := 1.0
	for n > 0 {
		if n%2 == 1 {
			r *= x
		}

		x *= x
		n /= 2
	}

	return r
}
