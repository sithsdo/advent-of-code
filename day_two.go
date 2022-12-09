package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
)

// A, X === Rock === 1
// B, Y === Paper === 2
// C, Z === Scissors === 3

func rockPaperScissors() {
	// Score Tracking
	lost := 0
	draw := 3
	won := 6
	fmt.Println("Begin reading File!", lost, draw, won)

	file, err := os.Open("inputDate/day_two_input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// playerOne := 0
	playerTwo := 0

	for scanner.Scan() {
		current := scanner.Text()
		p1 := current[0:1]
		p2 := current[len(current)-1:]

		// playerTwo += gameController(p1, p2)
		playerTwo += gameControllerV2(p1, p2)
		// fmt.Println("First: ", current[0:1], "Second: ", current[len(current)-1:])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total Score: ", playerTwo)
}

func gameController(p1 string, p2 string) (a int) {
	// A, X === Rock === 1
	// B, Y === Paper === 2
	// C, Z === Scissors === 3
	if p1 == "A" && p2 == "X" {
		return 3 + 1
	} else if p1 == "B" && p2 == "Y" {
		return 3 + 2
	} else if p1 == "C" && p2 == "Z" {
		// Draw
		// Rock and Rock || Paper and Paper || Scissors and Scissors
		return 3 + 3
	} else if p1 == "A" && p2 == "Y" {
		// loss
		// Rock < Paper
		return 2 + 6
	} else if p1 == "A" && p2 == "Z" {
		// win
		// Rock > Scissors
		return 3 + 0
	} else if p1 == "B" && p2 == "X" {
		// loss
		// Paper > Rock
		return 1 + 0
	} else if p1 == "B" && p2 == "Z" {
		// win
		// Paper < Scissors
		return 3 + 6
	} else if p1 == "C" && p2 == "X" {
		// loss
		// Scissors < Rock
		return 1 + 6
	} else if p1 == "C" && p2 == "Y" {
		// loss
		// Scissors > Paper
		return 2 + 0
	}
	return 0
}

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
func gameControllerV2(p1 string, p2 string) (a int) {
	// A, X === Rock === 1
	// B, Y === Paper === 2
	// C, Z === Scissors === 3

	// X == Lose
	// Y == Draw
	// Z == Win
	if p2 == "X" {
		// lose
		switch p1 {
		case "A":
			return 3 + 0
		case "B":
			return 1 + 0
		case "C":
			return 2 + 0
		default:
			return 0
		}
	} else if p2 == "Y" {
		// draw
		switch p1 {
		case "A":
			return 1 + 3
		case "B":
			return 2 + 3
		case "C":
			return 3 + 3
		default:
			return 0
		}
	} else if p2 == "Z" {
		// win
		switch p1 {
		case "A":
			return 2 + 6
		case "B":
			return 3 + 6
		case "C":
			return 1 + 6
		default:
			return 0
		}
	}
	return 0
}
