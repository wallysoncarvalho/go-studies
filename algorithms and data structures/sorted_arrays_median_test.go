package main

import (
	"testing"
)

func TestFindMedian(t *testing.T) {
	nums1 := []int{0, 2, 4, 6}
	nums2 := []int{1, 3, 5, 7}

	result := findMedianSortedArrays(nums1, nums2)
	expected := 3.5

	if result != expected {
		t.Errorf("got %f but was expecting %f", result, expected)
	}
}

type testFindMedianInput struct {
	nums1 []int
	nums2 []int
	expected float64
}

var testInputs = []testFindMedianInput{
	{[]int{0, 2, 4, 6}, []int{1, 3, 5, 7}, 3.5},
	{[]int{0, 2}, []int{1}, 1},
	{[]int{8, 10}, []int{1, 3, 5, 7}, 6},
}

func TestBatchTestFindMedian(t *testing.T) {

	for _, test := range testInputs {

		if output := findMedianSortedArrays(test.nums1, test.nums2); output != test.expected {
			t.Errorf("got %f but was expecting %f", output, test.expected)
		}
	}

}
