package shared

import (
	"iter"
	"maps"
	"slices"
)

type Grid map[complex128]byte

func (g *Grid) FindAny(b byte) complex128 {
	for k, v := range *g {
		if v == b {
			return k
		}
	}
	return 0
}

func (g *Grid) FindAll(b byte) []complex128 {
	var ret []complex128
	for k, v := range *g {
		if v == b {
			ret = append(ret, k)
		}
	}
	return ret
}

func NewGrid(arr []string) Grid {
	ret := make(map[complex128]byte, len(arr)*len(arr[0]))
	for ri, row := range arr {
		for ci, byt := range []byte(row) {
			ret[complex(float64(ci), float64(ri))] = byt
		}
	}
	return ret
}

type Edge struct {
	U, V   int
	Weight int
}

// Find the minimum spanning tree of n nodes
func Kruskal(n int, edges []Edge) []Edge {
	slices.SortFunc(edges, func(a, b Edge) int {
		return a.Weight - b.Weight
	})
	dsu := NewDSU(n)
	var mst []Edge
	for _, edge := range edges {
		if dsu.Union(edge.U, edge.V) {
			mst = append(mst, edge)
			if len(mst) == n-1 {
				break
			}
		}
	}
	return mst
}

type DSU struct {
	parent []int
	rank   []int
}

// Creates a new disjoint-set union for merging connected components in Kruskal's algorithm
func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range n {
		parent[i] = i
	}
	return &DSU{parent, rank}
}

func (d *DSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

// Joins the nodes U and V and merges the component roots
func (d *DSU) Union(U, V int) bool {
	rootU := d.find(U)
	rootV := d.find(V)
	if rootU == rootV {
		return false
	}
	if d.rank[rootU] < d.rank[rootV] {
		d.parent[rootU] = V
	} else if d.rank[rootU] > d.rank[rootV] {
		d.parent[rootV] = U
	} else {
		d.parent[rootV] = U
		d.rank[rootU]++
	}
	return true
}

// Returns the groups of connected nodes
func (d *DSU) Components() iter.Seq[[]int] {
	roots := make(map[int][]int)
	for i := range d.parent {
		r := d.find(i)
		roots[r] = append(roots[r], i)
	}
	return maps.Values(roots)
}
