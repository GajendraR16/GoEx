package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	if graph[from] == nil {
		graph[from] = make(map[string]bool)
	}
	graph[from][to] = true
}

func hasEdges(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("a", "b")
	addEdge("a", "c")
	addEdge("b", "c")
	fmt.Println(hasEdges("a", "c"))

	fmt.Println(graph)
}
