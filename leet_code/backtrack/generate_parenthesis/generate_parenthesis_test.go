package generate_parenthesis

import (
	"fmt"
	"testing"
)

var res [][]int

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"t1", args{n: 3}, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	res := [][]int{}
	for _, tt := range tests {
		/*t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})*/
		_ = tt.args
		k()
		fmt.Println(res)
	}
}

func k() {
	res = append(res, []int{1})
}
