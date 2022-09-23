package main

func twoSum(nums []int, target int) []int {
	lookedUp := map[int]int{}
	result := make([]int, 2)

	for index, number := range nums {
		complement := target - number

		seenIndex, exists := lookedUp[complement]

		if !exists {
			lookedUp[number] = index
			continue
		}

		result[0] = seenIndex
		result[1] = index
		break
	}

	return result
}
