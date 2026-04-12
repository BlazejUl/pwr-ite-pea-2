package atsp

import (
	"math"
	"time"

	"github.com/BlazejUl/pwr-ite-pea-2/graph"
)

type BranchAndBoundBreadthFirstSolver struct {
	graph      graph.Graph
	UpperBound int
	BestPath   []int
}

func NewBranchAndBoundBreadthFirstSolver(g graph.Graph) *BranchAndBoundBreadthFirstSolver {
	bestPath := make([]int, g.GetVerticesNum())
	return &BranchAndBoundBreadthFirstSolver{
		graph:      g,
		UpperBound: math.MaxInt,
		BestPath:   bestPath,
	}
}

func (atsp *BranchAndBoundBreadthFirstSolver) GetGraph() graph.Graph {
	return atsp.graph
}

func (atsp *BranchAndBoundBreadthFirstSolver) Solve(startVertex int, useTimer bool) (int, []int) {
	return atsp.BranchAndBoundBreadthFirstSolver(startVertex, useTimer)
}

// BFSNode reprezentuje węzeł w kolejce BFS
type BFSNode struct {
	Vertex     int
	LowerBound int
	Visited    []bool
	Path       []int
}

func (atsp *BranchAndBoundBreadthFirstSolver) BranchAndBoundBreadthFirstSolver(startVertex int, useTimer bool) (int, []int) {
	var timer time.Time
	if useTimer {
		timer = time.Now()
	}

	n := atsp.graph.GetVerticesNum()
	matrix := atsp.graph.GetMatrix()

	// Inicjalizacja kolejki FIFO
	startVisited := make([]bool, n)
	startVisited[startVertex] = true

	startNode := BFSNode{
		Vertex:     startVertex,
		LowerBound: 0,
		Visited:    startVisited,
		Path:       []int{startVertex},
	}

	queue := []BFSNode{startNode}

	for len(queue) > 0 {

		if useTimer {
			if time.Since(timer).Seconds() >= 180 {
				return -1, nil
			}
		}
		// Pobierz węzeł z przodu kolejki (FIFO)
		current := queue[0]
		queue = queue[1:]

		// Sprawdź czy odwiedziliśmy wszystkie wierzchołki
		allVisited := true
		for _, v := range current.Visited {
			if !v {
				allVisited = false
				break
			}
		}

		if allVisited {
			// Dodaj koszt powrotu do wierzchołka startowego
			returnCost := current.LowerBound + matrix[current.Vertex][startVertex]
			if returnCost < atsp.UpperBound {
				atsp.UpperBound = returnCost
				copy(atsp.BestPath, current.Path)
			}
			continue
		}

		// Rozwiń węzeł — dodaj wszystkich nieodwiedzonych sąsiadów do kolejki
		for i := 0; i < n; i++ {
			if !current.Visited[i] {
				newBound := current.LowerBound + matrix[current.Vertex][i]

				// Przytnij gałąź jeśli dolne ograniczenie >= górne ograniczenie
				if newBound >= atsp.UpperBound {
					continue
				}

				// Skopiuj stan odwiedzonych wierzchołków
				newVisited := make([]bool, n)
				copy(newVisited, current.Visited)
				newVisited[i] = true

				// Skopiuj ścieżkę
				newPath := make([]int, len(current.Path)+1)
				copy(newPath, current.Path)
				newPath[len(current.Path)] = i

				queue = append(queue, BFSNode{
					Vertex:     i,
					LowerBound: newBound,
					Visited:    newVisited,
					Path:       newPath,
				})
			}
		}
	}

	return atsp.UpperBound, atsp.BestPath
}
