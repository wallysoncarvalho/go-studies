package main

import (
	"fmt"
	// "time"
	// "log"
	log2 "github.com/sirupsen/logrus"
	//"math"
)

func main() {
	log2.SetFormatter(&log2.JSONFormatter{})

	fmt.Println(findMedianSortedArrays([]int{0, 2, 4, 6}, []int{1, 3, 5, 7}))

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// mergedSize := len(nums1) + len(nums2)
	var merged []int

	nums1Index := 0
	nums2Index := 0

	for nums1Index < len(nums1) && nums2Index < len(nums2) {
		var num int

		if nums1[nums1Index] < nums2[nums2Index] {
			num = nums1[nums1Index]
			nums1Index++
		}else {
			num = nums2[nums2Index]
			nums2Index++
		}

		merged = append(merged, num)
	}

	merged = append(merged, nums1[nums1Index:]...)
	merged = append(merged, nums2[nums2Index:]...)

	return getArrayMedian(merged)
}

func getArrayMedian(arr []int) float64 {
	arrSize := len(arr)

	if arrSize % 2 != 0 {
		return float64(arr[arrSize/2])
	}

	leftMiddle := float64(arr[(arrSize / 2) -1 ])
	rightMiddle := float64(arr[arrSize / 2])

	return (leftMiddle + rightMiddle) / 2
}
