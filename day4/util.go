package main

import (
	"fmt"
	"strings"
)

type Pair struct {
	X, Y int
}

func Map[I, O any](s []I, f func(I) O) []O {
	result := make([]O, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

func InBounds[I any](x, y int, grid [][]I) bool {
	numRows := len(grid)
	numCols := len(grid[0])

	return 0 <= x && x < numRows && 0 <= y && y < numCols
}

func CheckNeighbors(x, y int, grid [][]string) int {
	sum := 0
	for _, pair := range []Pair{{X: -1, Y: -1}, {X: -1, Y: 0}, {X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: -1}, {X: 0, Y: -1}} {
		newX := x + pair.X
		newY := y + pair.Y
		if InBounds(newX, newY, grid) && grid[newX][newY] == "@" {
			sum += 1
		}
	}
	return sum
}

func PrettyPrintGrid(grid [][]string) {
	fmt.Println("[")
	for _, row := range grid {
		fmt.Println("\t" + strings.Join(row, ""))
	}
	fmt.Println("]")
}
