package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// [V]     [B]                     [C]
// [C]     [N] [G]         [W]     [P]
// [W]     [C] [Q] [S]     [C]     [M]
// [L]     [W] [B] [Z]     [F] [S] [V]
// [R]     [G] [H] [F] [P] [V] [M] [T]
// [M] [L] [R] [D] [L] [N] [P] [D] [W]
// [F] [Q] [S] [C] [G] [G] [Z] [P] [N]
// [Q] [D] [P] [L] [V] [D] [D] [C] [Z]
//  1   2   3   4   5   6   7   8   9

func blockSort() {
	file, err := os.Open("inputDate/day_five_input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gameBoard := [9]string{"QFMRLWCV", "DQL", "PSRGWCNB", "LCDHBQG", "VGLFZS", "DGNP", "DZPVFCW", "CPDMS", "ZNWTVMPC"}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		current := scanner.Text()
		currentArr := strings.Split(current, " ")

		numberToMove, _ := strconv.Atoi(currentArr[1])
		initialIndex, _ := strconv.Atoi(currentArr[3])
		newIndex, _ := strconv.Atoi(currentArr[5])

		//grab last item to move
		tmpCurrentStr := gameBoard[initialIndex-1]
		tmpItemToMove := ""

		if len(tmpCurrentStr) > numberToMove {
			tmpItemToMove = tmpCurrentStr[len(tmpCurrentStr)-numberToMove:]
			gameBoard[initialIndex-1] = string(tmpCurrentStr[:len(tmpCurrentStr)-numberToMove])
		} else {
			tmpItemToMove = tmpCurrentStr
			gameBoard[initialIndex-1] = ""
		}

		gameBoard[newIndex-1] += tmpItemToMove
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("End Program!", gameBoard)
}

// VWLCWGSDQ
// TCGLQSLPW

// // for i := 0; i < numberToMove; i++ {
// 			//grab last item to move
// 			tmpCurrentStr := gameBoard[initialIndex-1]

// 			if len(tmpCurrentStr) > 0 {
// 				tmpItemToMove := string(tmpCurrentStr[len(tmpCurrentStr)-1:])

// 				// remove last item that has been moved
// 				gameBoard[initialIndex-1] = string(tmpCurrentStr[:len(tmpCurrentStr)-1])

// 				// add new item to new location
// 				gameBoard[newIndex-1] += tmpItemToMove
// 				// fmt.Println("Count - ", initialIndex, tmpItemToMove, tmpCurrentStr[:len(tmpCurrentStr)-1])
// 			}

// 			// // // add new item to new location
// 			// gameBoard[newIndex] += tmpItemToMove

// 			// fmt.Println("Count - ", tmpItemToMove, tmpCurrentStr[:len(tmpCurrentStr)-1])
// 			// fmt.Println("Count - ", tmpItemToMove, "Start Index - ", tmpCurrentStr, "End Index - ", gameBoard[newIndex])
// 		// }
