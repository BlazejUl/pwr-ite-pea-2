package graph

type Graph interface {
	PutPath(startingV int, endingV int, value int) error
	GetPath(startingV int, endingV int) (int, error)
	GetMatrix() [][]int
	ToString() string
	GetCopy() Graph
	GetVerticesNum() int
	GetPathsNum() int
}
