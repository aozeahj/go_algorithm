package dispatch_candy_2

import (
	"reflect"
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	type args struct {
		candies   int
		numPeople int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{
			candies:   7,
			numPeople: 4,
		}, []int{1, 2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistributeCandies(tt.args.candies, tt.args.numPeople); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
