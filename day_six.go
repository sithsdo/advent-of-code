package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func signalLock() {
	file, err := os.Open("inputDate/day_six_input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currentString := ""

	for scanner.Scan() {
		current := scanner.Text()
		currentString += current
	}

	result := signalLockIndex(currentString, 14)

	// test_input_01 := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	// test_input_02 := "nppdvjthqldpwncqszvftbrmjlhg"
	// test_input_03 := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	// test_input_04 := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

	// test_input_01 := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	// test_input_02 := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	// test_input_03 := "nppdvjthqldpwncqszvftbrmjlhg"
	// test_input_04 := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"

	// result1 := signalLockIndex(test_input_01, 14)
	// result2 := signalLockIndex(test_input_02, 14)
	// result3 := signalLockIndex(test_input_03, 14)
	// result4 := signalLockIndex(test_input_04, 14)
	// fmt.Println("End Program!", result1, result2, result3, result4)

	fmt.Println("End Program!", result)

}

func signalLockIndex(input string, offset int) int {
	for i := offset; i <= len(input); i++ {
		currentSet := input[i-offset : i]
		if allCharDistinct(currentSet) {
			return i
		}
	}
	return 0
}

func allCharDistinct(str string) bool {
	for i, char1 := range str {
		cur := i + 1
		for _, char2 := range str[cur:] {
			if string(char1) == string(char2) {
				return false
			}
		}
	}
	return true
}
