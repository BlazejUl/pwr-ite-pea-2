package atsp

import (
	"github.com/BlazejUl/pwr-ite-pea-1/graph"
)

// pomocnicza struktur zawierająca graf
type BruteForce struct {
	graph graph.Graph
}

// funkcja zwracająca graf na którym pracuje bruteforce
func (bf *BruteForce) GetGraph() graph.Graph {
	return bf.graph
}

// funkcja tworząca nowy BruteForce
func NewBruteforce(g graph.Graph) *BruteForce {
	return &BruteForce{graph: g}
}

// funkcja rozwiązująca problem atsp dla danego miasta startowego
// funkcja rozwiązująca problem atsp dla danego miasta startowego
func (bf *BruteForce) Solve(startVertex int) (int, []int) {
	visited := make([]bool, bf.graph.GetVerticesNum())
	path := make([]int, 0, bf.graph.GetVerticesNum())
	bestPath := make([]int, bf.graph.GetVerticesNum())
	bestCost := 2147483644

	visited[startVertex] = true
	path = append(path, startVertex)

	bestCost, bestPath = bf.BFRec(startVertex, visited, startVertex, 0, bestCost, bestPath, path)

	return bestCost, bestPath
}

// funkcja rekurencyjnie sprawdza wszystkie ścieżki od zadanego wierzchołka startowego
func (bf *BruteForce) BFRec(startVertex int, visited []bool, currentVertex int, currentCost int, bestCost int, bestPath []int, path []int) (int, []int) {
	// jeżeli skończy wraca do wierzchołka startowego czyli dodaje koszt przebycia drogi do początku
	if len(path) == bf.graph.GetVerticesNum() {
		cost, _ := bf.graph.GetPath(currentVertex, startVertex)
		currentCost += cost

		if currentCost < bestCost {
			bestCost = currentCost
			copy(bestPath, path)
		}

		return bestCost, bestPath
	}

	// Rekurencyjnie sprawdzi wszystkie ścieżki od wierzchołka startowego
	for i := 0; i < bf.graph.GetVerticesNum(); i++ {
		if !visited[i] {
			visited[i] = true
			cost, _ := bf.graph.GetPath(currentVertex, i)
			currentCost += cost
			path = append(path, i)

			bestCost, bestPath = bf.BFRec(startVertex, visited, i, currentCost, bestCost, bestPath, path)

			visited[i] = false
			currentCost -= cost
			path = path[:len(path)-1]
		}
	}

	return bestCost, bestPath
}
