package atsp

import (
	"math/rand"

	"github.com/BlazejUl/pwr-ite-pea-1/graph"
)

type Random struct {
	graph graph.Graph
}

func (ra *Random) GetGraph() graph.Graph {
	return ra.graph
}

func NewRandom(G graph.Graph) *Random {
	return &Random{graph: G}
}

// Funkcja sprawdzająca wylosowane rozwiązania i zapisująca najlepszy na podstawie kosztu
func (ra *Random) Solve(times int, startVertex int) (int, []int) {
	path := make([]int, 0, ra.graph.GetVerticesNum())
	bestPath := make([]int, ra.graph.GetVerticesNum())
	bestCost := 2147483644
	cost := 0

	for range times * ra.graph.GetVerticesNum() {
		cost, path = ra.getRandom(startVertex)
		if bestCost > cost {
			bestCost = cost
			copy(bestPath, path)
		}
	}

	return bestCost, bestPath
}

// Funkcja losowo wyznaczająca drogę oddaje też koszt tej drogi
func (ra *Random) getRandom(startVertex int) (int, []int) {
	path := make([]int, 0, ra.graph.GetVerticesNum())
	visited := make([]bool, ra.graph.GetVerticesNum())
	visited[0] = true
	path = append(path, startVertex)
	cost := 0
	currentVertex := startVertex
	for i := 0; i < ra.graph.GetVerticesNum()-1; i++ {
		r := rand.Intn(ra.graph.GetVerticesNum())
		// jeżeli dane misto zostało już odwiedzone to przechodzi po kolei aż natrafi na poprawne
		if visited[r] {
			for {
				r = (r + 1) % ra.graph.GetVerticesNum()
				if !visited[r] {
					break
				}
			}
		}
		visited[r] = true
		path = append(path, r)
		cst, _ := ra.graph.GetPath(currentVertex, r)
		cost = cost + cst
		currentVertex = r
	}
	//dodaje drogę spowrotem do punktu startowego
	cst, _ := ra.graph.GetPath(currentVertex, startVertex)
	cost = cost + cst

	return cost, path
}
