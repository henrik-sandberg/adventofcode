package y2025

import (
	"adventofcode/solutions/shared"
	"math"
	"slices"
	"strings"
)

type day08 struct {
	points map[int]struct {
		x, y, z int
	}
	edges []struct {
		pointA, pointB int
		weight         float64
	}
}

func Day08(input []string) (solution shared.Solution[int, int]) {
	return (&day08{}).solve(input, 1000)
}

func (d *day08) init(input []string) {
	d.points = make(map[int]struct {
		x, y, z int
	})
	for idx, s := range input {
		ints := shared.IntSlice(strings.Split(s, ","))
		d.points[idx] = struct{ x, y, z int }{x: ints[0], y: ints[1], z: ints[2]}
	}
	for idxA := range input {
		a := d.points[idxA]
		for idxB := idxA + 1; idxB < len(input); idxB++ {
			b := d.points[idxB]
			d.edges = append(d.edges, struct {
				pointA, pointB int
				weight         float64
			}{
				idxA, idxB, math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2) + math.Pow(float64(b.z-a.z), 2),
			})
		}
	}
	slices.SortFunc(d.edges, func(a, b struct {
		pointA, pointB int
		weight         float64
	}) int {
		return int(a.weight - b.weight)
	})
}

func (d *day08) solve(input []string, pairs int) (solution shared.Solution[int, int]) {
	d.init(input)
	clusters := d.allClusterSizes(pairs)
	solution.Part1 = clusters[0] * clusters[1] * clusters[2]

	low := 0
	high := len(d.edges) - 1
	mid := 0
	for low <= high {
		mid = low + (high-low)/2
		if len(d.allClusterSizes(mid)) > 1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	edge := d.edges[mid-1]
	solution.Part2 = d.points[edge.pointA].x * d.points[edge.pointB].x
	return

}

func (d *day08) allClusterSizes(maxEdges int) []int {
	edges := make(map[int][]int)
	for _, edge := range d.edges[:maxEdges] {
		edges[edge.pointA] = append(edges[edge.pointA], edge.pointB)
		edges[edge.pointB] = append(edges[edge.pointB], edge.pointA)
	}
	var clusters []int
	visited := make(map[int]bool)
	for _, edge := range d.edges {
		queue := []int{edge.pointA, edge.pointB}
		var node int
		var size int
		for len(queue) > 0 {
			node, queue = queue[0], queue[1:]
			if visited[node] {
				continue
			}
			size++
			visited[node] = true
			queue = append(queue, edges[node]...)
		}
		if size > 0 {
			clusters = append(clusters, size)
		}
	}
	slices.Sort(clusters)
	slices.Reverse(clusters)
	return clusters
}
