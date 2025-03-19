package floydWarshall

import (
	"fmt"
	"math"
)

func FloydWarshall(graph [][]int) ([][]int, [][]int) {
	n := len(graph)
	dist := make([][]int, n) // Matrix of the shortest distances
	next := make([][]int, n) // next[i][j] points to the next node on the shortest path from i to j

	// Initialize graphs
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		next[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = graph[i][j]
			if graph[i][j] != math.MaxInt32 && i != j {
				next[i][j] = j
			} else {
				next[i][j] = -1
			}
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 &&
					dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
					next[i][j] = next[i][k]
				}
			}
		}
	}

	return dist, next
}

// FindPath returns the shortest path between start and end vertexes
func FindPath(next [][]int, start, end int) []int {
	if next[start][end] == -1 {
		return nil
	}

	path := []int{start}
	for curr := start; curr != end; curr = next[curr][end] {
		path = append(path, next[curr][end])
	}
	return path
}

func PrintFloydWarshallGraph(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			if val == math.MaxInt32 {
				fmt.Print("∞\t")
			} else {
				fmt.Printf("%d\t", val)
			}
		}
		fmt.Println()
	}
}
