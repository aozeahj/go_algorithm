package quick

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{a: []int{10, 1, 2, 3, 2, 3, 5, 6, 7, 8, 9, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.a)
			fmt.Println(tt.args.a)
		})
	}
}

func TestSortByStack(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{a: []int{10, 1, 2, 3, 2, 3, 5, 6, 7, 8, 9, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortByStack(tt.args.a)
			fmt.Println(tt.args.a)
		})
	}
}
