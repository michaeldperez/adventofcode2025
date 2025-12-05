package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func mapSlice[I, O any](s []I, f func(I) O) []O {
	result := make([]O, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

func getLargestLocalJolt(s []int) int {
	toStringSlice := mapSlice(s, func(i int) string {
		str := strconv.Itoa(i)
		return str
	})
	toString := strings.Join(toStringSlice, "")
	toInt, _ := strconv.Atoi(toString)
	return toInt
}

func Day3Part2() {
	largestJolt := 0
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	banks := strings.Split(strings.TrimSpace(string(f)), "\n")

	toInt := func(s string) int {
		v, _ := strconv.Atoi(s)
		return v
	}

	for _, bank := range banks {
		bankToIntSlice := mapSlice(strings.Split(bank, ""), toInt)
		batteriesOnline := make([]int, 0, 12)

		want := 11
		idx := 0
		for want > 0 && len(bankToIntSlice) >= want {
			bound := len(bankToIntSlice) - want
			front := bankToIntSlice[:bound]
			max := slices.Max(front)
			idx = slices.Index(bankToIntSlice, max)
			bankToIntSlice = bankToIntSlice[idx+1:]
			batteriesOnline = append(batteriesOnline, max)
			want--
		}
		batteriesOnline = append(batteriesOnline, slices.Max(bankToIntSlice))
		largestLocalJolt := getLargestLocalJolt(batteriesOnline)
		largestJolt += largestLocalJolt

	}
	fmt.Printf("Largest total output: %d\n", largestJolt)
}
