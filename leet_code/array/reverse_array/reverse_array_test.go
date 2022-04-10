package reverse_array

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{nums: []int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}
