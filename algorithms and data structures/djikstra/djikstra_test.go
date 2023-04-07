package djikstra

import (
	
	"testing"
)

func TestMainStructure(t *testing.T) {

	// nodeA := Node{Value: 'A', Edges: []Edge{{To: 'B', Cost: 5}, {To: 'C', Cost: 3}}}
	// nodeB := Node{Value: 'B', Edges: []Edge{{To: 'D', Cost: 2}, {To: 'E', Cost: 6}, {To: 'G', Cost: 4}}}
	// nodeC := Node{Value: 'C', Edges: []Edge{{To: 'D', Cost: 6}, {To: 'E', Cost: 4}}}
	// nodeD := Node{Value: 'D', Edges: []Edge{{To: 'E', Cost: 1}, {To: 'F', Cost: 3}}}
	// nodeE := Node{Value: 'E', Edges: []Edge{{To: 'G', Cost: 2}}}
	// nodeF := Node{Value: 'F', Edges: []Edge{{To: 'G', Cost: 2}, {To: 'J', Cost: 4}}}
	// nodeG := Node{Value: 'G', Edges: []Edge{{To: 'H', Cost: 5}, {To: 'I', Cost: 3}}}
	// nodeH := Node{Value: 'H', Edges: []Edge{{To: 'J', Cost: 3}}}
	// nodeI := Node{Value: 'I', Edges: []Edge{{To: 'J', Cost: 6}}}
	// nodeJ := Node{Value: 'J', Edges: []Edge{}}

	// graph := Graph{Nodes: []Node{nodeA, nodeB, nodeC, nodeD, nodeE, nodeF, nodeG, nodeH, nodeI, nodeJ}}

	// A -> C -> E -> G -> I

	node1 := Node{Value: 'S', Edges: []Edge{{To: 'C', Cost: 5}, {To: 'A', Cost: 3}, {To: 'B', Cost: 2}}}
	node2 := Node{Value: 'A', Edges: []Edge{{To: 'D', Cost: 3}}}
	node3 := Node{Value: 'B', Edges: []Edge{{To: 'D', Cost: 1}, {To: 'E', Cost: 6}}}
	node4 := Node{Value: 'C', Edges: []Edge{{To: 'E', Cost: 2}}}
	node5 := Node{Value: 'D', Edges: []Edge{{To: 'F', Cost: 4}}}
	node6 := Node{Value: 'E', Edges: []Edge{{To: 'G', Cost: 4}, {To: 'F', Cost: 1}}}
	node7 := Node{Value: 'F', Edges: []Edge{{To: 'G', Cost: 2}}}
	node8 := Node{Value: 'G', Edges: []Edge{}}

	// result: 
	graph2 := Graph{Nodes: []Node{node1, node2, node3, node4, node5, node6, node7, node8}}




	dijkstra(graph2)

}