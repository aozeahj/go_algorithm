package priority_queue

import "github.com/aozeahj/go_algorithm/heap"

/**
优先队列是堆的一种具体应用，所以可以先参考堆，优先队列常用语操作系统的任务调度

优先队列的特点是根据元素的关键字进行排序管理

优先队列也分为 最大优先队列 最小优先队列

以最大优先队列举例，包含一下方法
（1）插入 Insert
（2）查找拥有最大关键字的元素 Maximum
（3）查找并且取出拥有最大关键字的元素 ExtractMaximum
（4）增加指定元素的关键字大小 Increase

*/

/**
插入 Insert
*/
func Insert(a *[]int, e int) {
	*a = append(*a, 0)
	IncreaseToK(*a, len(*a)-1, e)
}

/**
增加指定元素的关键字大小 Increase, 注意此处 k 默认为大于零的整数
*/
func IncreaseToK(a []int, index int, k int) {
	if index < 1 || index > len(a)-1 {
		return
	}

	a[index] = k
	largest := index
	parent := index / 2
	for parent > 0 && a[parent] < a[largest] {
		a[parent], a[largest] = a[largest], a[parent]
		largest = parent
		parent = parent / 2
	}
}

/**
查找拥有最大关键字的元素 Maximum
*/
func Maximum(a []int) (int, bool) {
	if len(a) > 0 {
		return a[1], true
	}
	return 0, false
}

/**
查找并且取出拥有最大关键字的元素 ExtractMaximum
*/
func ExtractMaximum(a *[]int) (int, bool) {
	arr := *a
	if len(arr) < 1 {
		return 0, false
	}
	max := arr[1]
	arr[1] = arr[len(arr)-1]
	heap.AdjustHeap(arr[:len(arr)-1], 1)
	*a = arr[:len(arr)-1]
	return max, true
}
