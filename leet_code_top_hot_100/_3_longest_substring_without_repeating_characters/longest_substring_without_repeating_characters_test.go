package _3_longest_substring_without_repeating_characters

import (
	"fmt"
	"testing"
)

func Test_test(t *testing.T) {
	fmt.Println(test())
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{s: "abcabcbb"}, 3},
		{"", args{s: " "}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
