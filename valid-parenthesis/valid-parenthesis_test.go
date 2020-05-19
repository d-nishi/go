package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name string
	in   string
	out  bool
}{
	{
		name: "equals 1",
		in:   "(){}[]",
		out:  true,
	},
	{
		name: "not equals 1",
		in:   "(){[]",
		out:  true,
	},
	{
		name: "not equals 2",
		in:   "(",
		out:  true,
	},
	{
		name: "equals 2",
		in:   "()",
		out:  true,
	},
	{
		name: "not equals 3",
		in:   "(){(}[]",
		out:  true,
	},
	{
		name: "equals",
		in:   " ",
		out:  true,
	},
}

func TestisValid(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.in)
			if !(reflect.DeepEqual(got, tt.out)) {
				t.Errorf("test failed")
			}
		})
	}
}
