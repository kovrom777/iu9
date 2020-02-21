package main

import (
	"fmt"
)

type graph struct {
	v     []str
	fi    [][]string
	count int
}

type str struct {
	dep, index int
	pi, parent *str
	isused     bool
	d          []*str
}

func find(S *str) *str {
	if S.parent == S {
		return S
	}
	S.parent = find(S.parent)
	return S.parent

}

func union(a *str, b *str) {
	art, brt := find(a), find(b)
	if art.dep < brt.dep {
		art.parent = brt
	} else {
		brt.parent = art
		if art.dep == brt.dep && art != brt {
			art.dep++
		}
	}
}

func initial(gra graph, n, m int, delta [][]int) (graph, [][]int) {

	delta = make([][]int, n)
	for i := 0; i < n; i++ {
		delta[i] = make([]int, m)
	}

	gra.v = make([]str, n)
	for i := range gra.v {
		gra.v[i].d = make([]*str, m)
		for j := range gra.v[i].d {
			fmt.Scan(&delta[i][j])
			gra.v[i].d[j] = &gra.v[delta[i][j]]
		}
	}

	gra.fi = make([][]string, n)
	for i := 0; i < n; i++ {
		gra.fi[i] = make([]string, m)
		for j := range gra.fi[i] {
			fmt.Scan(&gra.fi[i][j])
		}
	}
	return gra, delta
}

func (G *graph) split1(n int, m int) int {
	q := n
	for i := 0; i < n; i++ {
		G.v[i].parent = &G.v[i]
		G.v[i].dep = 0
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if find(&G.v[i]) != find(&G.v[j]) {
				eq := true
				for k := 0; k < m; k++ {
					if G.fi[i][k] != G.fi[j][k] {
						eq = false
						break
					}
				}
				if eq {
					union(&G.v[i], &G.v[j])
					q--
				}
			}
		}
	}
	for i := range G.v {
		G.v[i].pi = find(&G.v[i])
	}
	return q
}

func (G *graph) split(n int, m int) int {
	q := n
	for i := 0; i < n; i++ {
		G.v[i].parent = &G.v[i]
		G.v[i].dep = 0
	}
	for i := range G.v {
		for j := i + 1; j < n; j++ {
			if G.v[i].pi == G.v[j].pi && find(&G.v[i]) != find(&G.v[j]) {
				eq := true
				for k := 0; k < m; k++ {
					w1 := G.v[i].d[k]
					w2 := G.v[j].d[k]
					if w1.pi != w2.pi {
						eq = false
						break
					}
				}
				if eq {
					union(&G.v[i], &G.v[j])
					q--
				}
			}
		}
	}
	for i := range G.v {
		G.v[i].pi = find(&G.v[i])
	}
	return q
}

func aufenkampHohn(a, n, m int, g graph, delta [][]int) graph {
	h := g.split1(n, m)
	for h1 := g.split(n, m); h != h1; h1 = g.split(n, m) {
		h = h1
	}
	g.DFS(a, delta, g.v[a].pi)
	return g
}

func (G *graph) DFS(b int, delta [][]int, s *str) {
	s.index = G.count
	G.count++
	s.isused = true
	for i := range delta[b] {
		next := delta[b][i]
		if !s.d[i].pi.isused {
			G.DFS(next, delta, s.d[i].pi)
		}
	}
}

func main() {
	var (
		gra   graph
		n     int
		m     int
		q0    int
		delta [][]int
	)
	gra.count = 0
	fmt.Scan(&n, &m, &q0)

	gra, delta = initial(gra, n, m, delta)

	gra = aufenkampHohn(q0, n, m, gra, delta)
	fmt.Print("digraph {\n")
	fmt.Print("rankdir=LR\n")
	fmt.Print("dummy [label = \"\", shape = none]\n")
	for i := 0; i < gra.count; i++ {
		fmt.Printf("%d [shape=circle]\n", i)
	}
	fmt.Printf("dummy -> %d\n", 0)

	for i := range gra.v {
		if gra.v[i].isused == true {
			for j := range gra.v[i].d {
				if gra.v[i].d[j].pi.isused == true {
					fmt.Printf("%d -> %d [label = \"%c(%s)\"]\n", gra.v[i].index, gra.v[i].d[j].pi.index, 'a'+j, gra.fi[i][j])
				}
			}
		}
	}
	fmt.Print("}\n")
}
