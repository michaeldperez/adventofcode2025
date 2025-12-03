package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2Part1() {
	var invalidIdsSum uint64 = 0
	f, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	idRanges := strings.Split(string(f), ",")

	for _, idRange := range idRanges {
		openRange := strings.Split(idRange, "-")
		beginRangeStr := strings.TrimSpace(string(openRange[0]))
		endRangeStr := strings.TrimSpace(string(openRange[1]))

		start, err := strconv.ParseUint(beginRangeStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		end, err := strconv.ParseUint(endRangeStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		for i := start; i <= end; i++ {
			intToStr := strconv.FormatUint(i, 10)
			strLen := len(intToStr)
			halfwayPoint := strLen / 2
			front := intToStr[:halfwayPoint]
			back := intToStr[halfwayPoint:]
			if front == back {
				invalidIdsSum += i
			}
		}
	}
	fmt.Printf("Invalid IDs Sum: %d\n", invalidIdsSum)
}
