package atsp

import "github.com/BlazejUl/pwr-ite-pea-1/graph"

type RepetitiveNN struct {
	graph graph.Graph
}

func (rnn *RepetitiveNN) GetGraph() graph.Graph {
	return rnn.graph
}

func NewRepetitiveNN(G graph.Graph) *RepetitiveNN {
	return &RepetitiveNN{graph: G}
}

// Funkcja przygotowująca zmienne i zapisująca rozwiązanie z RNN
func (rnn *RepetitiveNN) Solve(startVertex int) (int, []int) {
	visited := make([]bool, rnn.graph.GetVerticesNum())
	path := make([]int, 0, rnn.graph.GetVerticesNum())
	bestPath := make([]int, rnn.graph.GetVerticesNum())
	bestCost := 2147483644

	visited[startVertex] = true
	path = append(path, startVertex)

	bestCost, bestPath = rnn.RNNRec(startVertex, visited, startVertex, 0, bestCost, bestPath, path)

	return bestCost, bestPath
}

// funkcja RNN która bada rekurencyjnie ścieżki ale tylko te które mają najmniejszy koszt
func (rnn *RepetitiveNN) RNNRec(startVertex int, visited []bool, currentVertex int, currentCost int, bestCost int, bestPath []int, path []int) (int, []int) {
	// jeżeli skończy wraca do wierzchołka startowego czyli dodaje koszt przebycia drogi do początku
	if len(path) == rnn.graph.GetVerticesNum() {
		cost, _ := rnn.graph.GetPath(currentVertex, startVertex)
		currentCost += cost

		if currentCost < bestCost {
			bestCost = currentCost
			copy(bestPath, path)
		}

		return bestCost, bestPath
	}

	// Szuka najpierw najtańszej nie odwiedzonej ścieżki od obecnego wierzchołka
	bestPathValue := 2147483644
	for i := 0; i < rnn.graph.GetVerticesNum(); i++ {
		if cost, _ := rnn.graph.GetPath(currentVertex, i); !visited[i] && bestPathValue > cost {
			bestPathValue = cost
		}
	}
	// Jeżeli więcej niż 1 ścieżka jest najmniejsza to Rekurencyjnie sprawdzi wszystkie te ścieżki
	for i := 0; i < rnn.graph.GetVerticesNum(); i++ {
		if cost, _ := rnn.graph.GetPath(currentVertex, i); !visited[i] && bestPathValue == cost {
			visited[i] = true
			cost, _ := rnn.graph.GetPath(currentVertex, i)
			currentCost += cost
			path = append(path, i)

			bestCost, bestPath = rnn.RNNRec(startVertex, visited, i, currentCost, bestCost, bestPath, path)

			visited[i] = false
			currentCost -= cost
			path = path[:len(path)-1]
		}
	}

	return bestCost, bestPath
}
