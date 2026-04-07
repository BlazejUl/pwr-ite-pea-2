package main

import (
	"fmt"
	"math"
	"time"

	"github.com/BlazejUl/pwr-ite-pea-1/atsp"
	"github.com/BlazejUl/pwr-ite-pea-1/graph"
	"github.com/BlazejUl/pwr-ite-pea-1/utils"
)

func main() {
	var G graph.Graph
	OutName := "out\\"
	fileName := ""
	inputV := 0
	b := 1
	n := 1
	r := 1
	rnd := 1
	rap := 1
	for {
		PrintMenu()
		var opt int
		if _, err := fmt.Scanln(&opt); err != nil {
			fmt.Println(err)
		}
		switch opt {
		case 1:
			fmt.Println("1 - załaduj macież z pliku")
			fmt.Println("2 - wygeneruj macież")
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
			default:
				fmt.Println("tylko numery od 1 - 2")
			}

		case 2:
			if G == nil {
				fmt.Println("graf nie został podany")
				break
			}
			PrintMenu2()
			if _, err := fmt.Scanln(&opt); err != nil {
				fmt.Println(err)
			}
			switch opt {
			case 1:
				bf := atsp.NewBruteforce(G)
				start := time.Now()
				cost, path := bf.Solve(0)
				lTimeInMikro := time.Since(start).Microseconds()
				name := fmt.Sprintf("%dBruteForce.txt", b)
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
			case 2:
				nn := atsp.NewNN(G)
				start := time.Now()
				cost, path := nn.Solve(0)
				lTimeInMikro := time.Since(start).Microseconds()
				name := fmt.Sprintf("%dNearestNeighbour.txt", n)
				n++
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
			case 3:
				rnn := atsp.NewRepetitiveNN(G)
				start := time.Now()
				cost, path := rnn.Solve(0)
				lTimeInMikro := time.Since(start).Microseconds()
				name := fmt.Sprintf("%dRepetitiveNN.txt", r)
				r++
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
			case 4:
				fmt.Println("Podaj mnożnik ilości permutacji x * liczbaMiast")
				if _, err := fmt.Scanln(&opt); err != nil {
					fmt.Println(err)
				}
				ra := atsp.NewRandom(G)
				start := time.Now()
				cost, path := ra.Solve(opt, 0)
				lTimeInMikro := time.Since(start).Microseconds()
				name := fmt.Sprintf("%dRandom.txt", rnd)
				rnd++
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
			default:
				fmt.Println("tylko numery od 1 - 4")
			}

		case 3:
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
					bt := atsp.NewBruteforce(G)
					nn := atsp.NewNN(G)
					rnn := atsp.NewRepetitiveNN(G)
					ra := atsp.NewRandom(G)
					bCost, _ := bt.Solve(0)
					nCost, _ := nn.Solve(0)
					rnCost, _ := rnn.Solve(0)
					rCost, _ := ra.Solve(10, 0)
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
		case 4:
			return
		default:
			fmt.Println("tylko numery od 1 - 4")
		}
	}

}

func PrintMenu() {
	fmt.Println("program do testowania rozwiązań problemu atsp")
	fmt.Println("1 - załaduj macież")
	fmt.Println("2 - rozwiąż macież")
	fmt.Println("3 - testuj trafność rozwiązań")
	fmt.Println("4 - Wyjdź")
}

func PrintMenu2() {
	fmt.Println("1 - rozwiąż za pomocą bruteforce")
	fmt.Println("2 - rozwiąż za pomocą nearest neighbour")
	fmt.Println("3 - rozwiąż za pomocą repetitive nearest neighbour")
	fmt.Println("4 - rozwiąż za pomocą random")
}
