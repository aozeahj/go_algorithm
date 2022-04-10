package quick

func Partition(a []int, l int, r int) int {
	pivot := a[r]
	index := l - 1
	for j := l; j < r; j++ {
		if a[j] <= pivot {
			index++
			if j > index {
				a[index], a[j] = a[j], a[index]
			}
		}
	}
	a[index+1], a[r] = a[r], a[index+1]
	return index + 1
}

func QSort(a []int, l int, r int) {
	if l >= r {
		return
	}

	index := Partition(a, l, r)

	QSort(a, l, index-1)
	QSort(a, index+1, r)
}

func Sort(a []int) {
	QSort(a, 0, len(a)-1)
}

/**
非递归版本
*/
func SortByStack(a []int) {
	stack := []int{len(a) - 1, 0}
	for len(stack) > 0 {
		l := stack[len(stack)-1]
		r := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		index := Partition(a, l, r)
		if index-1 > l {
			stack = append(stack, []int{index - 1, l}...)
		}
		if index+1 < r {
			stack = append(stack, []int{r, index + 1}...)
		}
	}
}
