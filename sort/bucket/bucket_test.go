package bucket

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		a []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"", args{a: []float64{0.1, 0, 0.3, 0.2, 0.7, 0.3, 0.4}}, []float64{0, 0.1, 0.2, 0.3, 0.3, 0.4, 0.7}},
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
