package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name string
	in   [][]int
	out  []int
}{
	{
		name: "equals",
		in: [][]int{
			[]int{1, 30, 51},
			[]int{2, 15, 20, 30},
		},
		out: []int{1, 2, 15, 20, 30, 30, 51},
	},
	{
		name: "not equals",
		in: [][]int{
			[]int{2, 8, 15},
			[]int{12},
		},
		out: []int{2, 8, 12, 15},
	},
	{
		name: "3 array test",
		in: [][]int{
			[]int{1, 2, 2},
			[]int{3, 7},
			[]int{8, 9, 14, 31},
		},
		out: []int{1, 2, 2, 3, 7, 8, 9, 14, 31},
	},
	{
		name: "4 array test",
		in: [][]int{
			[]int{1, 2, 2},
			[]int{3, 7},
			[]int{8, 9, 14, 31},
			[]int{10, 19, 24},
		},
		out: []int{1, 2, 2, 3, 7, 8, 9, 10, 14, 19, 24, 31},
	},
	{
		name: "5 array test",
		in: [][]int{
			[]int{1, 2, 6},
			[]int{3, 3},
			[]int{8, 9, 24, 31},
			[]int{10, 19, 44},
			[]int{15, 16, 34},
		},
		out: []int{1, 2, 3, 3, 6, 8, 9, 10, 15, 16, 19, 24, 31, 34, 44},
	},
}

func TestMergeKSortedArrs(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeKSortedArrs(tt.in)
			want := tt.out
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v and want: %v\n", got, want)
			}
		})
	}
}
