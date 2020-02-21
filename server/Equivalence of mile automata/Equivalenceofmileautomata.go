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
		gra1, gra2     graph
		n1, n2         int
		m1, m2         int
		q01, q02       int
		delta1, delta2 [][]int
	)
	gra1.count = 0
	gra2.count = 0
	fmt.Scan(&n1, &m1, &q01)
	gra1, delta1 = initial(gra1, n1, m1, delta1)

	fmt.Scan(&n2, &m2, &q02)
	gra2, delta2 = initial(gra2, n2, m2, delta2)

	gra1 = aufenkampHohn(q01, n1, m1, gra1, delta1)
	gra2 = aufenkampHohn(q02, n2, m2, gra2, delta2)

	var str1, str2 []string
	str1 = make([]string, gra1.count*m1)
	str2 = make([]string, gra2.count*m2)

	count := 0

	for i := range gra1.v {
		if gra1.v[i].isused == true {
			for j := range gra1.v[i].d {
				if gra1.v[i].d[j].pi.isused == true {
					str1[count] += fmt.Sprintf("%v %v %v %v", gra1.v[i].index, gra1.v[i].pi.index, 'a'+j, gra1.fi[i][j])
					count++
				}
			}
		}
	}
	count = 0
	// fmt.Printf("%d -> %d [label = \"%c(%s)\"]\n", gra.v[i].index, gra.v[i].d[j].pi.index, 'a'+j, gra.fi[i][j])
	for i := range gra2.v {
		if gra2.v[i].isused == true {
			for j := range gra2.v[i].d {
				if gra2.v[i].d[j].pi.isused == true {
					str2[count] += fmt.Sprintf("%v %v %v %v", gra2.v[i].index, gra2.v[i].pi.index, 'a'+j, gra2.fi[i][j])
					count++
				}
			}
		}
	}

	// fmt.Print(str1)
	// fmt.Println()
	// fmt.Print(str2)

	count = 0
	if len(str1) == len(str2) {
		for i := 0; i < len(str1); i++ {
			for j := 0; j < len(str1); j++ {
				if str1[i] == str2[j] {
					count++
				}
			}
		}
	} else {
		fmt.Print("NOT EQUAL\n")
		return
	}
	if count == len(str1) {
		fmt.Print("EQUAL\n")
		return
	} else {
		fmt.Print("NOT EQUAL\n")
		return
	}
}
