package main

import (
	"fmt"
	"math"
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
	rap := 1
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
			cost, path := bfs.Solve(0)
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
			return
		case 5:
			fmt.Println("Podaj ilość miast")
			if _, err := fmt.Scanln(&opt); err != nil {
				fmt.Println(err)
			} else {
				raport := "   NN   |  ReNN  | Random \n"
				bAllCost := 0
				nAllCost := 0
				rnAllCost := 0
				rAllCost := 0
				for range 100 {
					G, _ = utils.GenerateAdMatrix(opt)
					bt := atsp.NewBranchAndBoundBestFirstSolver(G)
					nn := atsp.NewBranchAndBoundBreadthFirstSolver(G)
					rnn := atsp.NewBranchAndBoundBestFirstSolver(G)
					ra := atsp.NewBranchAndBoundBestFirstSolver(G)
					bCost, _ := bt.Solve(0)
					nCost, _ := nn.Solve(0)
					rnCost, _ := rnn.Solve(0)
					rCost, _ := ra.Solve(0)
					bAllCost = bAllCost + bCost
					nAllCost = nAllCost + nCost
					rnAllCost = rnAllCost + rnCost
					rAllCost = rAllCost + rCost
					nB := (math.Abs(float64((bCost - nCost))) / float64(bCost)) * 100
					rnB := (math.Abs(float64((bCost - rnCost))) / float64(bCost)) * 100
					rB := (math.Abs(float64((bCost - rCost))) / float64(bCost)) * 100
					iterS := fmt.Sprintf("%.2f%% | %.2f%% | %.2f%%\n", nB, rnB, rB)
					raport = raport + iterS
				}
				nBk := (math.Abs(float64((bAllCost - nAllCost))) / float64(bAllCost)) * 100
				rnBk := (math.Abs(float64((bAllCost - rnAllCost))) / float64(bAllCost)) * 100
				rBk := (math.Abs(float64((bAllCost - rAllCost))) / float64(bAllCost)) * 100
				lIterS := fmt.Sprintf("%.2f%% | %.2f%% | %.2f%%\n", nBk, rnBk, rBk)
				raport = raport + "--------------------\n"
				raport = raport + lIterS
				name := fmt.Sprintf("%dTestAll.txt", rap)
				if err := utils.WriteFile(OutName+name, raport); err != nil {
					fmt.Printf("////////error %d", err)
				}
				rap++
				fmt.Println("   NN   |  ReNN  | Random ")
				fmt.Printf("%.2f%% | %.2f%% | %.2f%%\n", nBk, rnBk, rBk)
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
