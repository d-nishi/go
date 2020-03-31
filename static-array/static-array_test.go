package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name string
	out  []int
	want bool
}{
	{
		name: "equals",
		out:  []int{1, 2, 5, 7, 9},
		want: true,
	},
	{
		name: "not equals",
		out:  []int{1, 1, 4, 7, 9},
		want: false,
	},
}

func TestCreateArr(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateArr()
			equals := reflect.DeepEqual(got, tt.out)
			if equals != tt.want {
				t.Errorf("test failed")
			}
		})
	}
}
