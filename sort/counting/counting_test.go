package counting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "", args: args{a: []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}}, want: []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}},
		{name: "", args: args{a: []int{1, 9, 3, 7, 4, 5, 6, 3, 8, 2, 10}}, want: []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
			fmt.Println(Sort(tt.args.a))
		})
	}
}
