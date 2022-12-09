package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

// func PriorityEstimator() {
// 	// Lowercase item types a through z have priorities 1 through 26.
// 	// Uppercase item types A through Z have priorities 27 through 52.
// 	lower := "abcdefghijklmnopqrstuvwxyz"
// 	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 	file, err := os.Open("inputDate/day_three_input.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	prioritySum := 0

// 	for scanner.Scan() {
// 		current := scanner.Text()

// 		firstHalf := current[0 : len(current)/2]
// 		secondHalf := current[len(current)-len(current)/2:]
// 		currentSet := ""

// 		for _, letter := range firstHalf {
// 			if strings.Contains(secondHalf, string(letter)) && strings.Contains(currentSet, string(letter)) != true {
// 				currentSet += string(letter)
// 				fmt.Println(currentSet)
// 				if unicode.IsUpper(letter) {
// 					prioritySum += (strings.Index(upper, string(letter)) + 26 + 1)
// 					// fmt.Println("Upper ", string(letter), strings.Index(upper, string(letter)))
// 					// fmt.Println(firstHalf, " - ", secondHalf)
// 				} else {
// 					prioritySum += (strings.Index(lower, string(letter)) + 1)
// 					// fmt.Println("Lower ", string(letter), strings.Index(lower, string(letter)))
// 					// fmt.Println(firstHalf, " - ", secondHalf)
// 				}
// 			}
// 		}

// 		// fmt.Println("Length - ", len(firstHalf), " - ", "Length - ", len(secondHalf))
// 		// fmt.Println(firstHalf, " - ", secondHalf)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Priority Sum: ", prioritySum)
// 	fmt.Println("End Program!")
// }

func PriorityEstimator() {
	// Lowercase item types a through z have priorities 1 through 26.
	// Uppercase item types A through Z have priorities 27 through 52.
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	file, err := os.Open("inputDate/day_three_input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	prioritySum := 0

	set := []string{}
	currentSet := ""

	for scanner.Scan() {
		current := scanner.Text()

		// firstHalf := current[0 : len(current)/2]
		// secondHalf := current[len(current)-len(current)/2:]
		// currentSet := ""

		set = append(set, current)
		if len(set) == 3 {
			// fmt.Println(len(set))
			for _, letter := range set[0] {
				if (strings.Contains(set[1], string(letter)) && strings.Contains(set[2], string(letter))) && strings.Contains(currentSet, string(letter)) != true {
					currentSet += string(letter)
					if unicode.IsUpper(letter) {
						prioritySum += (strings.Index(upper, string(letter)) + 26 + 1)
						fmt.Println(currentSet, (strings.Index(upper, string(letter)) + 26 + 1))
						fmt.Println(set)
						// fmt.Println("Upper ", string(letter), strings.Index(upper, string(letter)))
						// fmt.Println(firstHalf, " - ", secondHalf)
					} else {
						prioritySum += (strings.Index(lower, string(letter)) + 1)
						fmt.Println(currentSet, (strings.Index(lower, string(letter)) + 26 + 1))
						fmt.Println(set)
						// fmt.Println("Lower ", string(letter), strings.Index(lower, string(letter)))
						// fmt.Println(firstHalf, " - ", secondHalf)
					}
					break
				}
			}
			set = []string{}
			currentSet = ""

			// fmt.Println("CLEAR, ", set)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Priority Sum: ", prioritySum)
	fmt.Println("End Program!")
}
