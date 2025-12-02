package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func filter[T any](s []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(s))
	for _, v := range s {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	dial := 50
	zeros := 0

	f, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	rawInput := string(f)
	rotations := strings.Split(rawInput, "\n")
	rotations = filter(rotations, func(rotation string) bool { return len(rotation) > 0 })
	for _, rotation := range rotations {
		r := string(rotation[0])
		digits := rotation[1:]
		total, err := strconv.Atoi(digits)

		if err != nil {
			log.Fatal(err)
		}
		for _ = range total {
			i := 1
			if r == "L" {
				i *= -1
			}
			dial = ((dial+i)%100 + 100) % 100
			if dial == 0 {
				zeros++
			}
		}
	}
	fmt.Printf("Password: %d\n", zeros)
}
