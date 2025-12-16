package y2025

import (
	"adventofcode/solutions/shared"
	"slices"
	"strings"
)

func Day08(input []string) (solution shared.Solution[int, int]) {
	return day08solve(input, 1000)
}

func day08solve(input []string, pairs int) (solution shared.Solution[int, int]) {
	type point struct {
		id      int
		x, y, z int
	}
	var points []point
	for idx, s := range input {
		ints := shared.IntSlice(strings.Split(s, ","))
		points = append(points, point{
			id: idx,
			x:  ints[0], y: ints[1], z: ints[2],
		})
	}
	var edges []shared.Edge
	for _, a := range points {
		for _, b := range points[a.id+1:] {
			edges = append(edges, shared.Edge{
				U:      a.id,
				V:      b.id,
				Weight: (b.x-a.x)*(b.x-a.x) + (b.y-a.y)*(b.y-a.y) + (b.z-a.z)*(b.z-a.z),
			})
		}
	}
	slices.SortFunc(edges, func(a, b shared.Edge) int {
		return a.Weight - b.Weight
	})
	dsu := shared.NewDSU(len(points))
	for _, edge := range edges[:pairs] {
		dsu.Union(edge.U, edge.V)
	}
	var clusterSizes []int
	for c := range dsu.Components() {
		clusterSizes = append(clusterSizes, len(c))
	}
	slices.Sort(clusterSizes)
	slices.Reverse(clusterSizes)
	solution.Part1 = clusterSizes[0] * clusterSizes[1] * clusterSizes[2]

	mst := shared.Kruskal(len(points), edges)
	last := mst[len(mst)-1]
	solution.Part2 = points[last.U].x * points[last.V].x
	return
}
