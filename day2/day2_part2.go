package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Chunks(s string, chunkSize int) []string {
	if chunkSize <= 0 {
		return nil
	}

	var chunks []string

	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize

		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

func AllEqual(s []string) bool {
	base := s[0]
	for _, el := range s[1:] {
		if el != base {
			return false
		}
	}
	return true
}

func Day2Part2() {
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
			noInvalid := true
			intToStr := strconv.FormatUint(i, 10)
			strLen := len(intToStr)
			for j := 1; j < strLen && noInvalid; j++ {
				if strLen%j == 0 {
					chunks := Chunks(intToStr, j)
					if AllEqual(chunks) {
						invalidIdsSum += i
						noInvalid = false
					}
				}
			}
		}
	}
	fmt.Printf("Invalid IDs Sum: %d\n", invalidIdsSum)
}
