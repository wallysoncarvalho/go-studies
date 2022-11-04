package main

func networkDelayTime(times [][]int, n int, k int) int {
	adjacencyList := adjacencyList(times, n)

	stack := []int{k}

	visited := make([]int, n)

	

	for len(stack) > 0 {

		source := stack[0]

		for dest, cost := range adjacencyList[source] {

			// if not visited
			if visited[dest] == 0 {

			}

			stack = append([]int{dest}, stack[1:]...)

			visited[dest] = cost
			
					
			

		}


	}



	return 0
}

func adjacencyList(times [][]int, n int) [][]int {
	adjacency := make([][]int, n)

	for _, tuple := range times {
		source := tuple[0]
		dest := tuple[1]
		cost := tuple[2]


		if adjacency[source] == nil {
			adjacency[source] = make([]int, n)
		}

		adjacency[source][dest] = cost
	}

	return adjacency
}
