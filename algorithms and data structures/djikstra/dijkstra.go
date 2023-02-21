package djikstra

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

	//visited['A'] = true

	unvisited := Stack{'A'}

	var teste := make([byte]map[byte]int)

	for unvisited.IsEmpty() {

		currentNode := unvisited.Pop()


		for _, edge := range graph.Nodes[currentNode].Edges {

			from, cost := teste[edge.To]
			

			

		}


		
		

	}
	


	

}
