package main

import (
	"reflect"
	"testing"
)

func TestSlidingWindow(t *testing.T) {
	var nums = []int {1,3,-1,-3,5,3,6,7}
	expected := []int{3,3,5,5,6,7}

	if result := maxSlidingWindow(nums, 3); reflect.DeepEqual(result, expected) {
		t.Errorf("Was expecting slice %d but got %d", expected, result)
	}
}
