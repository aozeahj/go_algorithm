package test

import (
	"testing"
)

type Iter func() (int, bool)

func tk(n int) Iter {
	data := []int{}
	for i := 0; i < n; i++ {
		data = append(data, i)
	}

	cnt := 0
	l := len(data)
	return func() (int, bool) {
		if cnt >= l {
			return 0, false
		}
		e := data[cnt]
		cnt++
		return e, true
	}
}

func TestA_sayHi(t *testing.T) {

}

func test() (ret int) {
	i := 100
	defer func(a int) {
		ret = i + 100
	}(i)

	return 300
}
