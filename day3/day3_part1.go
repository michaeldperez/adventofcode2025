package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day3Part1() {
	largestJolt := 0
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	banks := strings.Split(strings.TrimSpace(string(f)), "\n")

	for _, bank := range banks {
		largestLocalJolt := 0
		stop := len(bank) - 1
		for i, c := range bank[:stop] {
			firstDigit, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			for _, d := range bank[i+1:] {
				secondDigit, err := strconv.Atoi(string(d))
				if err != nil {
					log.Fatal(err)
				}
				total := (firstDigit * 10) + secondDigit

				if total > largestLocalJolt {
					largestLocalJolt = total
				}
			}
		}
		largestJolt += largestLocalJolt
	}
	fmt.Printf("Largest total output: %d\n", largestJolt)
}
