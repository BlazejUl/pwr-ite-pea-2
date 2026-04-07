package atsp

import (
	"github.com/BlazejUl/pwr-ite-pea-1/graph"
)

// pomocnicza struktur zawierająca graf
type NN struct {
	graph graph.Graph
}

// funkcja zwracająca graf na którym pracuje NN
func (nn *NN) GetGraph() graph.Graph {
	return nn.graph
}

// funkcja tworząca nowy NN
func NewNN(g graph.Graph) *NN {
	return &NN{graph: g}
}

// funkcja rozwiązująca problem atsp dla danego miasta startowego
func (nn *NN) Solve(startVertex int) (int, []int) {
	visited := make([]bool, nn.graph.GetVerticesNum())
	path := make([]int, 0, nn.graph.GetVerticesNum())
	cost := 0
	currentVertex := startVertex
	visited[startVertex] = true
	path = append(path, startVertex)

	for i := 0; i < nn.graph.GetVerticesNum()-1; i++ {
		lCost := 2147483644
		lVert := 0
		for j := 0; j < nn.graph.GetVerticesNum(); j++ {
			if !visited[j] {
				if currentCost, _ := nn.graph.GetPath(currentVertex, j); lCost > currentCost {
					lCost = currentCost
					lVert = j
				}
			}
		}
		currentVertex = lVert
		visited[currentVertex] = true
		path = append(path, currentVertex)
		cost = cost + lCost
	}
	//dodaje drogę powrotną do całkowitego kosztu
	vr, _ := nn.graph.GetPath(currentVertex, startVertex)
	cost = cost + vr

	return cost, path
}
