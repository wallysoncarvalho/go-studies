package djikstra

import (
	"fmt"
	"testing"
)

func TestMainStructure(t *testing.T) {

	nodeA := Node{Value: 'A', Edges: []Edge{{To: 'B', Cost: 5}, {To: 'C', Cost: 3}}}
	nodeB := Node{Value: 'B', Edges: []Edge{{To: 'D', Cost: 2}, {To: 'E', Cost: 6}, {To: 'G', Cost: 4}}}
	nodeC := Node{Value: 'C', Edges: []Edge{{To: 'D', Cost: 6}, {To: 'E', Cost: 4}}}
	nodeD := Node{Value: 'D', Edges: []Edge{{To: 'E', Cost: 1}, {To: 'F', Cost: 3}}}
	nodeE := Node{Value: 'E', Edges: []Edge{{To: 'G', Cost: 2}}}
	nodeF := Node{Value: 'F', Edges: []Edge{{To: 'G', Cost: 2}, {To: 'J', Cost: 4}}}
	nodeG := Node{Value: 'G', Edges: []Edge{{To: 'H', Cost: 5}, {To: 'I', Cost: 3}}}
	nodeH := Node{Value: 'H', Edges: []Edge{{To: 'J', Cost: 3}}}
	nodeI := Node{Value: 'I', Edges: []Edge{{To: 'J', Cost: 6}}}
	nodeJ := Node{Value: 'J', Edges: []Edge{}}

	graph := Graph{Nodes: []Node{nodeA, nodeB, nodeC, nodeD, nodeE, nodeF, nodeG, nodeH, nodeI, nodeJ}}

	// A -> C -> E -> G -> I

	fmt.Println(graph)


}