package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

func main() {
	var n, m, q0 int
	input.Scanf("%d\n%d\n%d\n", &n, &m, &q0)
	del := make([][]int, n)
	FI := make([][]rune, n)
	for i := 0; i < n; i++ {
		del[i] = make([]int, m)
		FI[i] = make([]rune, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			input.Scanf("%d ", &del[i][j])
		}
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < m; j++ {
			input.Scanf("%c ", &FI[i][j])
		}
	}

	for i := 0; i < m-1; i++ {
		input.Scanf("%c ", &FI[n-1][i])
	}
	input.Scanf("%c", &FI[n-1][m-1])

	fmt.Printf("digraph {\nrankdir = LR\n")
	fmt.Printf("dummy [label = \"\" , shape = none]")
	fmt.Printf("\n")

	for i := 0; i < n; i++ {
		fmt.Printf("%d [shape = circle]\n", i)
	}
	fmt.Printf("dummy -> %d\n", q0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d -> %d [label = \"%c(%c)\"]\n", i, del[i][j], 'a'+j, FI[i][j])
		}
	}
	fmt.Printf("}")
}
