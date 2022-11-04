package main

import (
	"testing"
)

func TestNetworkDelayTime(t *testing.T) {
	mapping := [][]int { {2,1,1}, {2,3,1}, {3,4,1} }
	var n int = 3
	var k int = 2
	var expected int = 2

	if result := networkDelayTime(mapping, n, k ); result != expected {
		t.Errorf("Was expecting '%d' but the result was '%d'", expected, result)
	}
}
