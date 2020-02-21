package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

func DFS(used []int, del [][]int, start int, index *int, m int) {
	used[start] = (*index)
	(*index) += 1
	for i := 0; i < m; i++ {
		if used[del[start][i]] == -1 {
			DFS(used, del, del[start][i], index, m)
		}
	}
}

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

	var del2 = make([][]int, n)
	var FI2 = make([][]rune, n)
	var used = make([]int, n)

	for i := 0; i < n; i++ {
		del2[i] = make([]int, m)
		FI2[i] = make([]rune, m)
		used[i] = -1
	}
	var in int
	DFS(used, del, q0, &in, m)

	for i := 0; i < n; i++ {
		if used[i] != -1 {
			for j := 0; j < m; j++ {
				FI2[used[i]][j] = FI[i][j]
				del2[used[i]][j] = used[del[i][j]]
			}
		}
	}

	fmt.Printf("\n%d\n%d\n%d\n", in, m, 0)

	for i := 0; i < in; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", del2[i][j])
		}
		fmt.Printf("\n")
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%c ", FI2[i][j])
		}
		fmt.Println()
	}

}
