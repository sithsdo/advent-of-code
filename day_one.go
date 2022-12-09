package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Elf struct {
	currentCalories int
	currentElf      int
}

// Instructions: https://adventofcode.com/2022/day/1

func caloricCount() {
	fmt.Println("Begin reading File!")

	file, err := os.Open("day_one_input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0

	first := Elf{0, 1}
	second := Elf{0, 1}
	third := Elf{0, 1}
	tempCurrent := Elf{0, 1}

	for scanner.Scan() {
		if scanner.Text() != "" {
			// update temp current calorie load
			currentVal, _ := strconv.Atoi(scanner.Text())
			tempCurrent.currentCalories += currentVal
		}

		if scanner.Text() == "" {
			// check if tempCurrent is more then current top three
			if tempCurrent.currentCalories > first.currentCalories {
				first.currentCalories = tempCurrent.currentCalories
				first.currentElf = tempCurrent.currentElf
			} else if tempCurrent.currentCalories > second.currentCalories {
				second.currentCalories = tempCurrent.currentCalories
				second.currentElf = tempCurrent.currentElf
			} else if tempCurrent.currentCalories > third.currentCalories {
				third.currentCalories = tempCurrent.currentCalories
				third.currentElf = tempCurrent.currentElf
			}

			// reset temp current to default
			tempCurrent.currentElf += 1
			tempCurrent.currentCalories = 0
		}
		lineCount += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total Line Count: ", lineCount)
	fmt.Println("First Elf: ", first.currentElf, "First Count: ", first.currentCalories)
	fmt.Println("Second Elf: ", second.currentElf, "Second Count: ", second.currentCalories)
	fmt.Println("Third Elf: ", third.currentElf, "Third Count: ", third.currentCalories)

	// Calculate total of top 3 Elf calorie loads
	total := first.currentCalories + second.currentCalories + third.currentCalories
	fmt.Println("Total: ", total)
	fmt.Println("Done!")
}
