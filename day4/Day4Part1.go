package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Day4Part1() {
	accessibleRolls := 0
	f, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fs := strings.TrimSpace(string(f))

	toRows := strings.Split(fs, "\n")

	grid := Map(toRows, func(r string) []string {
		return strings.Split(r, "")
	})

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != "@" {
				continue
			}
			if CheckNeighbors(row, col, grid) < 4 {
				accessibleRolls += 1
			}
		}

	}

	fmt.Printf("Accessible roles: %d\n", accessibleRolls)
}
