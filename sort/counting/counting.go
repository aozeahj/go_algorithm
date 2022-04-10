package counting

/**
计数排序，在输入数组A中元素大小在 0~k 之间，可以通过空间换时间的方式，将其排序复杂度从 o(nlgn) 降低到 o(n)
*/
func Sort(a []int) []int {
	l := len(a)
	result := make([]int, l)
	counter := make([]int, l)

	for _, v := range a {
		counter[v]++
	}

	for i := 1; i < l; i++ {
		counter[i] = counter[i] + counter[i-1]
	}

	for _, v := range a {
		result[counter[v]-1] = v
		counter[v]--
	}

	return result
}
