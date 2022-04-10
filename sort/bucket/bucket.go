package bucket

import "sort"

/**
桶排序，如果输入数组A中的元素A[i]的值可以近视满足 0~1均匀分布 则可以将将数组A的n个元素先分成n个桶中，
每个桶单独排序后，再按照桶的顺序合并，最后的结果即为有序的
*/

//假设a 数组中的元素满足 0~1 均匀分布
func Sort(a []float64) []float64 {
	l := len(a)
	buckets := make([][]float64, l)

	for _, v := range a {
		bucketIndex := int(v * float64(l))
		buckets[bucketIndex] = append(buckets[bucketIndex], v)
	}

	for _, bucket := range buckets {
		sort.Float64s(bucket)
	}

	result := make([]float64, 0, len(a))
	for _, bucket := range buckets {
		result = append(result, bucket...)
	}

	return result
}
