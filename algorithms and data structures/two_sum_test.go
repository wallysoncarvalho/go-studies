package main

import "testing"

type testInput struct {
	nums     []int
	target   int
	expected []int
}

func TestFindTwoSum(t *testing.T) {
	var inputs = []testInput{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for _, test := range inputs {

		output := twoSum(test.nums, test.target)

		if output[0] != test.expected[0] || output[1] != test.expected[1] {
			t.Errorf("got %v but was expecting %v", output, test.expected)
		}
	}
}
