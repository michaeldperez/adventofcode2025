package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func Map[T, O any](s []T, f func(T) O) []O {
	result := make([]O, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

func Day5Part1() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs := string(f)
	idx := strings.Index(fs, "\n\n")
	ranges := strings.Split(fs[:idx], "\n")
	ingredientIdsStr := strings.Split(strings.TrimSpace(fs[idx+1:]), "\n")
	ingredientIds := Map(ingredientIdsStr, func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	})

	set := mapset.NewSet[int]()

	for _, id := range ingredientIds {
		for _, rng := range ranges {
			r := strings.Split(rng, "-")
			low, _ := strconv.Atoi(r[0])
			high, _ := strconv.Atoi(r[1])
			if id >= low && id <= high {
				set.Add(id)
			}
		}
	}
	fmt.Printf("Number of fresh ingredents: %d\n", set.Cardinality())
}
