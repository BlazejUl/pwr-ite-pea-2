package graph

import "fmt"

type AdMatrix struct {
	vertices int
	paths    int
	matrix   [][]int
}

func NewAdMatrix(vertices int) (*AdMatrix, error) {
	if vertices <= 0 {
		return nil, fmt.Errorf("vertices must be a positive number given: %d", vertices)
	}
	am := &AdMatrix{vertices: vertices}

	am.matrix = make([][]int, vertices)

	for i := 0; i < am.vertices; i++ {
		am.matrix[i] = make([]int, am.vertices)
	}

	return am, nil
}

func (am *AdMatrix) PutPath(startingV int, endingV int, value int) error {

	if startingV < 0 || startingV >= am.vertices {
		return fmt.Errorf("starting vertice out of bounds gotten: %d", startingV)
	}

	if endingV < 0 || endingV >= am.vertices {
		return fmt.Errorf("ending vertice out of bounds gotten: %d", endingV)
	}

	if value < -1 {
		return fmt.Errorf("value must be bigger than -1 gotten: %d", value)
	}

	if am.matrix[startingV][endingV] > -1 {
		am.paths++
	}
	am.matrix[startingV][endingV] = value

	return nil
}

func (am *AdMatrix) GetPath(startingV int, endingV int) (int, error) {
	if startingV < 0 || startingV >= am.vertices {
		return -2, fmt.Errorf("starting vertice out of bounds gotten: %d", startingV)
	}

	if endingV < 0 || endingV >= am.vertices {
		return -2, fmt.Errorf("ending vertice out of bounds gotten: %d", endingV)
	}

	return am.matrix[startingV][endingV], nil
}

func (am *AdMatrix) GetMatrix() [][]int {
	return am.matrix
}

func (am *AdMatrix) ToString() string {
	var str string
	for i := 0; i < am.vertices; i++ {
		for j := 0; j < am.vertices; j++ {
			str += fmt.Sprintf("%d ", am.matrix[i][j])
		}
		str += "\n"
	}
	return str
}

func (am *AdMatrix) GetCopy() Graph {
	amCopy := &AdMatrix{vertices: am.vertices, paths: am.paths}
	amCopy.matrix = make([][]int, amCopy.vertices)
	for i := 0; i < amCopy.vertices; i++ {
		amCopy.matrix[i] = make([]int, amCopy.vertices)
		copy(amCopy.matrix[i], am.matrix[i])
	}

	return amCopy
}

func (am *AdMatrix) GetVerticesNum() int {
	return am.vertices
}

func (am *AdMatrix) GetPathsNum() int {
	return am.paths
}
