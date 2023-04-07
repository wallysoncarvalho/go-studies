package djikstra

import "fmt"

type Node struct {
	Value byte
	Edges []Edge
}

type Edge struct {
	To   byte
	Cost int
}

type Graph struct {
	Nodes []Node
}

func (g Graph) getNodeByValue(value byte) *Node {

	for _, node := range g.Nodes {

		if node.Value == value {
			return &node
		}

	}

	return nil
}

type Stack []byte

func (s *Stack) Push(v byte) {
    *s = append(*s, v)
}

func (s *Stack) Pop() byte {
    if s.IsEmpty() {
        return 0
    }

    lastIndex := len(*s) - 1
    result := (*s)[lastIndex]
    *s = (*s)[:lastIndex]
    return result
}

func (s *Stack) IsEmpty() bool {
    return len(*s) == 0
}

func dijkstra(graph Graph) {

	var visited = make(map[byte]bool)

	unvisited := Stack{'S'}

	var costs = make(map[byte]int)

	var costFrom = make(map[byte]byte)
	
	for !unvisited.IsEmpty() {
		currentNodeValue := unvisited.Pop()

		currentNode := graph.getNodeByValue(currentNodeValue)

		for _, edge := range currentNode.Edges {

			value, ok := costs[edge.To]


			if !ok || value > edge.Cost {

				costs[edge.To] = edge.Cost
				
				costFrom[edge.To] = currentNodeValue
			}

			if _, wasVisited := visited[edge.To]; !wasVisited {
				unvisited.Push(edge.To)
			}
			
		}

		visited[currentNodeValue] = true

	}

	for key, value := range costs {
		fmt.Printf("%c: %d ", key, value)
	}

	fmt.Println()

	for key, value := range costFrom {
		fmt.Printf("%c: %c ", key, value)
	}

}

