package main

import "fmt"

func printView(view [][]int) {
	fmt.Println("Start of View\n\n")

	for i := 0; i < len(view); i++ {
		fmt.Print("\t")
		for j := 0; j < len(view); j++ {
			fmt.Printf("%v ", view[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("\n\nEnd of View")
}

func regionsBySlashes(grid []string) int {

	n := len(grid)
	var view [][]int = make([][]int, n*n*2)

	V := 4
	E := 4

	for i := 0; i < len(view); i++ {
		view[i] = make([]int, len(view))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {

		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {

			switch string(grid[i][j]) {
			case "/":
				for t, k := i*n*2, (j+1)*n*2; t < (i+1)*(n*2); {
					k--
					view[t][k] = 1
					t++
				}
			case "\\":
				for t, k := (i+1)*n*2, (j+1)*n*2; k > (j * n * 2); {
					t--
					k--
					view[t][k] = 1

				}
			}

		}
	}
	printView(view)
	return 2 + E - V
}

//|Face yüz | = 2 + |Edge Kenar| - |Vertex Köşe|

func main() {

	//var text []string = []string{"\\/", "/\\"}
	var text []string = []string{"/\\", "\\/"}
	regionsBySlashes(text)
}
