package main

import (
	"fmt"
	"time"

	"github.com/BlazejUl/pwr-ite-pea-2/atsp"
	"github.com/BlazejUl/pwr-ite-pea-2/graph"
	"github.com/BlazejUl/pwr-ite-pea-2/utils"
)

func main() {
	var G graph.Graph
	OutName := "out\\"
	fileName := ""
	inputV := 0
	b := 1
	c := 1
	for {
		PrintMenu()
		var opt int
		if _, err := fmt.Scanln(&opt); err != nil {
			fmt.Println(err)
		}
		switch opt {
		case 1:
			fmt.Println("Podaj nazwę pliku .txt z macieżą (musi znajdować się w folderze in i nie może mieć spacji w nazwie oraz macież musi się kończyć znakiem nowej linii)")
			if _, err := fmt.Scanln(&fileName); err != nil {
				fmt.Println(err)
			} else {
				if G, err = utils.ReadGraphFromFile("in\\" + fileName); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(G.ToString())
				}
			}
			//
		case 2:
			fmt.Println("Podaj liczbę miast")
			if _, err := fmt.Scanln(&inputV); err != nil {
				fmt.Println(err)
			} else {
				if G, err = utils.GenerateAdMatrix(inputV); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(G.ToString())
				}
			}
			//
		case 3:
			if G == nil {
				fmt.Println("graf nie został podany")
				break
			}
			bfs := atsp.NewBranchAndBoundBestFirstSolver(G)
			start := time.Now()
			cost, path := bfs.Solve(0, false)
			lTimeInMikro := time.Since(start).Microseconds()
			name := fmt.Sprintf("%dBestFirst.txt", b)
			b++
			info := fmt.Sprintf("czas: %d µs\nkoszt: %d\nścieżka: %d\n", lTimeInMikro, cost, path)
			lw := fmt.Sprintf("%d\n", G.GetVerticesNum())
			graphStr := lw + G.ToString() + "\n"
			if err := utils.WriteFile(OutName+name, info+graphStr); err != nil {
				fmt.Printf("////////error %d", err)
			}

			if err := utils.WriteFile(OutName+"LastMatrix.txt", graphStr); err != nil {
				fmt.Printf("////////error %d", err)
			}
			fmt.Printf("czas: %d µs\nkoszt: %d\nścieżka: %d\n", lTimeInMikro, cost, path)
			//

		case 4:
			if G == nil {
				fmt.Println("graf nie został podany")
				break
			}
			bfs := atsp.NewBranchAndBoundBreadthFirstSolver(G)
			start := time.Now()
			cost, path := bfs.Solve(0, false)
			lTimeInMikro := time.Since(start).Microseconds()
			name := fmt.Sprintf("%dBreadthFirst.txt", c)
			c++
			info := fmt.Sprintf("czas: %d µs\nkoszt: %d\nścieżka: %d\n", lTimeInMikro, cost, path)
			lw := fmt.Sprintf("%d\n", G.GetVerticesNum())
			graphStr := lw + G.ToString() + "\n"
			if err := utils.WriteFile(OutName+name, info+graphStr); err != nil {
				fmt.Printf("////////error %d", err)
			}

			if err := utils.WriteFile(OutName+"LastMatrix.txt", graphStr); err != nil {
				fmt.Printf("////////error %d", err)
			}
			fmt.Printf("czas: %d µs\nkoszt: %d\nścieżka: %d\n", lTimeInMikro, cost, path)
		case 5:
			fmt.Println("Podaj ilość miast")
			n := 100
			if _, err := fmt.Scanln(&opt); err != nil {
				fmt.Println(err)
			} else {
				bestAll := 0
				breadthAll := 0
				for range n {
					G, _ = utils.GenerateAdMatrix(opt)
					bt := atsp.NewBranchAndBoundBestFirstSolver(G)
					nn := atsp.NewBranchAndBoundBreadthFirstSolver(G)
					bestCost, _ := bt.Solve(0, true)
					breadthCost, _ := nn.Solve(0, true)
					if bestCost > 0 {
						bestAll++
					}
					if breadthCost > 0 {
						breadthAll++
					}
				}
				bestR := (bestAll / n) * 100
				breadthR := (breadthAll / n) * 100
				fmt.Println("best | breadth")
				fmt.Printf("%d%% | %d%%\n", bestR, breadthR)
			}
		case 6:
			return
		default:
			fmt.Println("tylko numery od 1 - 6")
		}
	}

}

func PrintMenu() {
	fmt.Println("program do testowania rozwiązań problemu atsp")
	fmt.Println("1 - załaduj macież z pliku")
	fmt.Println("2 - wygeneruj macież")
	fmt.Println("3 - rozwiąż za pomocą best-first-search")
	fmt.Println("4 - rozwiąż za pomocą bredth-first-search")
	fmt.Println("5 - testuj % rozwiązań dla 3min")
	fmt.Println("6 - Wyjdź")
}
