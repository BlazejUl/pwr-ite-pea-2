package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BlazejUl/pwr-ite-pea-1/graph"
)

func ReadGraphFromFile(filePath string) (graph.Graph, error) {
	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	rdr := bufio.NewReader(f)
	var line string
	line, err = rdr.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.TrimSpace(line)

	vertices, err := strconv.Atoi(line)
	if err != nil {
		return nil, err
	}

	if vertices < 1 {
		return nil, fmt.Errorf("vertices must be more than 1 gotten: %d", vertices)
	}

	G, err := graph.NewAdMatrix(vertices)

	if err != nil {
		return nil, err
	}

	for i := range vertices {
		line, err = rdr.ReadString('\n')

		if err != nil {
			return nil, err
		}

		paths := strings.Fields(strings.TrimSpace(line))

		for j, path := range paths {
			value, err := strconv.Atoi(path)

			if err != nil {
				return nil, err
			}

			err = G.PutPath(i, j, value)

			if err != nil {
				return nil, err
			}

		}
	}
	return G, nil
}

func WriteFile(filePath string, info string) error {
	f, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer f.Close()

	rdr := bufio.NewWriter(f)

	_, err = rdr.WriteString(info)

	if err != nil {
		return err
	}

	err = rdr.Flush()

	if err != nil {
		return err
	}

	return nil
}
