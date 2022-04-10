package priority_queue

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	type args struct {
		a *[]int
		e int
	}
	tests := []struct {
		name string
		args args
	}{

		{name: "", args: args{a: &[]int{0, 10, 9, 7, 8, 5, 6, 3, 1, 4, 2}, e: 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Insert(tt.args.a, tt.args.e)
			Insert(tt.args.a, 88)
			fmt.Println(tt.args.a)
		})
	}
}

func TestIncreaseToK(t *testing.T) {
	type args struct {
		a     []int
		index int
		k     int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{a: []int{0, 10, 9, 7, 8, 5, 6, 3, 1, 4, 2}, index: 3, k: 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IncreaseToK(tt.args.a, tt.args.index, tt.args.k)
			fmt.Println(tt.args.a)
		})
	}
}

func TestMaximum(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{name: "", args: args{a: []int{0, 10, 9, 7, 8, 5, 6, 3, 1, 4, 2}}, want: 10, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Maximum(tt.args.a)
			if got != tt.want {
				t.Errorf("Maximum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Maximum() got1 = %v, want %v", got1, tt.want1)
			}
			fmt.Println(tt.args.a)
		})
	}
}

func TestExtractMaximum(t *testing.T) {
	type args struct {
		a *[]int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{name: "", args: args{a: &[]int{0, 20, 16, 18, 8, 5, 6, 3, 1, 4, 2}}, want: 20, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ExtractMaximum(tt.args.a)
			if got != tt.want {
				t.Errorf("ExtractMaximum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExtractMaximum() got1 = %v, want %v", got1, tt.want1)
			}
			fmt.Println(tt.args.a)
		})
	}
}
