package main

import (
	"testing"
	"reflect"
)

func TestCreateArr(t *testing.T) {
	got := CreateArr()
	want:= []int{1,2,5,7,9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got,want)
	}
}

