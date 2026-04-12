package atsp

import (
	"math"
	"time"

	"github.com/BlazejUl/pwr-ite-pea-2/graph"
	"github.com/BlazejUl/pwr-ite-pea-2/utils"
)

type BranchAndBoundBestFirstSolver struct {
	graph      graph.Graph
	UpperBound int
	BestPath   []int
}

func NewBranchAndBoundBestFirstSolver(g graph.Graph) *BranchAndBoundBestFirstSolver {
	bestPath := make([]int, g.GetVerticesNum())
	return &BranchAndBoundBestFirstSolver{
		graph: g, UpperBound: math.MaxInt, BestPath: bestPath,
	}
}

func (atsp *BranchAndBoundBestFirstSolver) GetGraph() graph.Graph {
	return atsp.graph
}

func (atsp *BranchAndBoundBestFirstSolver) Solve(startVertex int, useTimer bool) (int, []int) {
	return atsp.BranchAndBoundBestFirst(startVertex, useTimer)
}

type BestFirstNode struct {
	Vertex     int
	LowerBound int
	Visited    []bool
	Path       []int
	PathCost   int
}

func (atsp *BranchAndBoundBestFirstSolver) minOutgoingEdge(v int, visited []bool) int {
	matrix := atsp.graph.GetMatrix()
	n := atsp.graph.GetVerticesNum()
	minCost := math.MaxInt

	for i := 0; i < n; i++ {
		//pomiń odwiedzone i siebie
		if visited[i] || matrix[v][i] == -1 {
			continue
		}
		if matrix[v][i] < minCost {
			minCost = matrix[v][i]
		}
	}

	if minCost == math.MaxInt {
		return 0
	}
	return minCost
}

// computeBound oblicza pathCost + drogę z tego do następnego +
// minimum z następnego do jeszcze nie odwiedzonych
func (atsp *BranchAndBoundBestFirstSolver) computeBound(
	pathCost int,
	currentVertex int,
	nextVertex int,
	visited []bool,
	startVertex int,
) int {
	matrix := atsp.graph.GetMatrix()
	n := atsp.graph.GetVerticesNum()

	edgeCost := matrix[currentVertex][nextVertex]
	if edgeCost == -1 {
		return math.MaxInt
	}

	newPathCost := pathCost + edgeCost

	// stwórz kopie odwiedzonych
	newVisited := make([]bool, n)
	copy(newVisited, visited)
	newVisited[nextVertex] = true

	// sprawdź ile nie odwiedzonych
	remaining := 0
	for i := 0; i < n; i++ {
		if !newVisited[i] {
			remaining++
		}
	}

	// jeżeli niema nie odwiedzonych to koszt drogi + koszt powrotu do startu
	if remaining == 0 {
		ret := matrix[nextVertex][startVertex]
		if ret == -1 {
			return math.MaxInt
		}
		return newPathCost + ret
	}

	bound := newPathCost

	// minimum wychodzące z następnego wierzchołka do nie odwiedzonych
	bound += atsp.minOutgoingEdge(nextVertex, newVisited)

	//minimum wychodzące z każdego następnego wierzchołka
	for i := 0; i < n; i++ {
		if !newVisited[i] {
			bound += atsp.minOutgoingEdge(i, newVisited)
		}
	}

	return bound
}

func (atsp *BranchAndBoundBestFirstSolver) BranchAndBoundBestFirst(startVertex int, useTimer bool) (int, []int) {
	n := atsp.graph.GetVerticesNum()
	matrix := atsp.graph.GetMatrix()

	var timer time.Time
	if useTimer {
		timer = time.Now()
	}

	atsp.UpperBound = math.MaxInt

	globalQueue := utils.NewPriorityQueue(func(a, b BestFirstNode) bool {
		return a.LowerBound < b.LowerBound
	})

	startVisited := make([]bool, n)
	startVisited[startVertex] = true

	// obliczenie początkowego ograniczenia dolnego
	initialBound := 0
	initialBound += atsp.minOutgoingEdge(startVertex, startVisited)
	for i := 0; i < n; i++ {
		if !startVisited[i] {
			initialBound += atsp.minOutgoingEdge(i, startVisited)
		}
	}

	startNode := BestFirstNode{
		Vertex:     startVertex,
		LowerBound: initialBound,
		Visited:    startVisited,
		Path:       []int{startVertex},
		PathCost:   0,
	}
	globalQueue.Push(startNode)

	for !globalQueue.IsEmpty() {
		if useTimer && time.Since(timer).Seconds() >= 180 {
			return -1, nil
		}

		current := globalQueue.Pop()

		if current.LowerBound >= atsp.UpperBound {
			continue
		}

		// sprawdź czy wszystkie odwiedzone
		allVisited := true
		for _, v := range current.Visited {
			if !v {
				allVisited = false
				break
			}
		}

		if allVisited {
			ret := matrix[current.Vertex][startVertex]
			if ret == -1 {
				continue
			}
			totalCost := current.PathCost + ret
			if totalCost < atsp.UpperBound {
				atsp.UpperBound = totalCost
				copy(atsp.BestPath, current.Path)
			}
			continue
		}

		// dodaj dzieci do kolejki
		for i := 0; i < n; i++ {
			if current.Visited[i] {
				continue
			}

			edgeCost := matrix[current.Vertex][i]
			if edgeCost == -1 {
				continue
			}

			newBound := atsp.computeBound(
				current.PathCost,
				current.Vertex,
				i,
				current.Visited,
				startVertex,
			)
			//jeżeli heurystycznie wyszło więcej niż upperBound to nie dodawaj
			if newBound >= atsp.UpperBound {
				continue
			}

			newVisited := make([]bool, n)
			copy(newVisited, current.Visited)
			newVisited[i] = true

			newPath := make([]int, len(current.Path)+1)
			copy(newPath, current.Path)
			newPath[len(current.Path)] = i

			globalQueue.Push(BestFirstNode{
				Vertex:     i,
				LowerBound: newBound,
				Visited:    newVisited,
				Path:       newPath,
				PathCost:   current.PathCost + edgeCost,
			})
		}
	}

	return atsp.UpperBound, atsp.BestPath
}
