package heap

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1", args: args{a: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}},
		{name: "test1", args: args{a: []int{0, 8, 9, 10, 4, 5, 6, 7, 1, 2, 3}}},
		{name: "test1", args: args{a: []int{0, 1, 1, 3, 4, 5, 9, 7, 8, 6, 10}}},
		{name: "test1", args: args{a: []int{0, 1, 2, 2, 4, 5, 3, 7, 8, 2, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.a)
			fmt.Println(tt.args.a)
		})
	}
}
