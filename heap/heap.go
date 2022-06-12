package heap

/*
接下来堆特指二叉堆
堆的特点
1.对于结点i而言，其父节点，左子节点，右子结点
	parent(i) = i/2
	left(i) = i*2
	right(i) = i*2+1
2.大小堆的定义
	A[parent(i)] >= A[i] 父节点的值均大于等于子节点的值则为最大堆
3.堆高
	一般我们使用数组来存储堆元素，长度为n的数组的堆高为lgn，相当于完全二叉树
4.堆元素数量
	高为h的堆的元素数量小于等于 2^h -1
4.堆最底层元素数量
	高位h的堆最底层元素数量小于等于 2 ^(h-1)
5.堆排序步骤
	（1）维护堆 adjustHeap
	（2）构建堆 buildHeap
	（3）排序   heapSort
*/

/*
维护堆 ,大堆
*/

func AdjustHeap(a []int, index int) {
	left := index * 2
	right := left + 1
	l := len(a)
	largest := index
	if left < l && a[left] > a[largest] {
		largest = left
	}

	if right < l && a[right] > a[largest] {
		largest = right
	}

	if largest == index {
		return
	}

	a[index], a[largest] = a[largest], a[index]
	AdjustHeap(a, largest)
}

/**
构建堆
*/
func BuildHeap(a []int) {
	for i := len(a) / 2; i > 0; i-- {
		AdjustHeap(a, i)
	}
}

func HeapSort(a []int) {
	BuildHeap(a)
	l := len(a)

	for i := 1; i < l-1; i++ {
		a[l-i], a[1] = a[1], a[l-i]
		AdjustHeap(a[:l-i], 1)
	}
}
