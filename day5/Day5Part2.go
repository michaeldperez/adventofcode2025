package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day5Part2() {
	totalFreshIngredientIds := 0
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs := string(f)
	idx := strings.Index(fs, "\n\n")
	ranges := strings.Split(fs[:idx], "\n")

	rangeSlice := Map(ranges, func(s string) []int {
		slice := strings.Split(s, "-")
		low, _ := strconv.Atoi(slice[0])
		high, _ := strconv.Atoi(slice[1])
		return []int{low, high}
	})

	slices.SortFunc(rangeSlice, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	var prev []int = nil
	for _, rng := range rangeSlice {
		if prev == nil {
			prev = rng
		} else if prev[1] < rng[0] {
			totalFreshIngredientIds += prev[1] - prev[0] + 1
			prev = rng
		} else {
			prev = []int{prev[0], max(prev[1], rng[1])}
		}
	}
	totalFreshIngredientIds += prev[1] - prev[0] + 1
	fmt.Printf("Total number of fresh ingredient IDs: %d\n", totalFreshIngredientIds)
}
