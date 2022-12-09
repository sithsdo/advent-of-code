package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func containedPairs() {
	file, err := os.Open("inputDate/day_four_input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sumOfContainedPairs := 0

	for scanner.Scan() {
		current := scanner.Text()

		currentArr := strings.Split(current, ",")
		firstSet := strings.Split(currentArr[0], "-")
		secondSet := strings.Split(currentArr[1], "-")

		baseRange1 := makeRange(convertToInt(firstSet[0]), convertToInt(firstSet[1]))
		baseRange2 := makeRange(convertToInt(secondSet[0]), convertToInt(secondSet[1]))

		for _, cur := range baseRange2 {
			if contains(baseRange1, cur) {
				sumOfContainedPairs++
				break
			}

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("End Program!", sumOfContainedPairs)
}

func convertToInt(str string) int {
	p1, _ := strconv.Atoi(str)
	return p1
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
