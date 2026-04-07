package utils

import (
	"math/rand"

	"github.com/BlazejUl/pwr-ite-pea-2/graph"
)

func GenerateAdMatrix(vertices int) (*graph.AdMatrix, error) {
	matrix, err := graph.NewAdMatrix(vertices)

	if err != nil {
		return nil, err
	}

	for i := 0; i < vertices; i++ {
		for j := 0; j < vertices; j++ {
			if i == j {
				matrix.PutPath(i, j, -1)
			} else {
				rndm := rand.Intn(100)
				if err = matrix.PutPath(i, j, rndm); err != nil {
					return nil, err
				}
			}
		}
	}
	return matrix, nil
}
