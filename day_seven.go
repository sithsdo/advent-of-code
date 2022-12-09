package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var minRequiredSize = 100000

func daySevenMain() {
	file, err := os.Open("inputDate/day_seven_input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	processFile(scanner)
}

func processFile(scanner *bufio.Scanner) {
	digitCheck := regexp.MustCompile(`^[0-9]+$`)      // regex for checking for type int
	directoryDepthTrackingList := []string{"/"}       // list of distinct folder names including historical depth
	directorySizeTrackingList := make(map[string]int) // map of folder keys with size

	for scanner.Scan() {
		current := string(scanner.Text())           // current line being read from input file
		currentSplit := strings.Split(current, " ") // split string by spaces

		if digitCheck.MatchString(currentSplit[0]) { // current line read is file and size, add it to current folder total size

			// iterate over current list of directories
			for i := len(directoryDepthTrackingList) - 1; i >= 0; i-- {
				dir := directoryDepthTrackingList[i]
				directorySizeTrackingList[dir] += convertToInt(currentSplit[0])
			}

		} else if strings.Contains(current, "$ cd") {
			direction := moveOutOrIn(current)

			switch direction {
			case "DOWN":
				dirDrillDown := strings.Split(current, " ")
				directoryDepthTrackingList = append(directoryDepthTrackingList, generateKey(directoryDepthTrackingList, dirDrillDown[2]))
			case "UP":
				if len(directoryDepthTrackingList) > 0 {
					directoryDepthTrackingList = directoryDepthTrackingList[:len(directoryDepthTrackingList)-1]
				}
			default:
				directoryDepthTrackingList = []string{"/"}
			}

		}

	}

	calcTotals(directorySizeTrackingList)
	fmt.Println("Total Calculated Size for Root Folder - ", directorySizeTrackingList["/"])
	fmt.Println("End Program!")
}

// Calculate sum of folder sizes
func calcTotals(dirList map[string]int) {
	totalSum := 0
	folderSizeToRemove := dirList["/"]
	toDeleteLimit := 70000000 - dirList["/"]
	if toDeleteLimit < 30000000 {
		toDeleteLimit = 30000000 - toDeleteLimit
	}
	for _, size := range dirList {
		if size <= minRequiredSize {
			totalSum += size
		}

		if size >= toDeleteLimit && size < folderSizeToRemove {
			folderSizeToRemove = size
		}
	}
	fmt.Println("Some of Folder sizes under limit - Phase 1! ", totalSum)
	fmt.Println("Folder to Delete - Phase 2! ", folderSizeToRemove)
}

// folder navigation
func moveOutOrIn(str string) string {
	if str == "$ cd /" {
		return "ROOT"
	} else if str == "$ cd .." {
		return "UP"
	} else {
		return "DOWN"
	}
}

// Generate a distinct key for current directory based off of parent folders
func generateKey(list []string, identifier string) string {
	tmpKey := ""
	for _, current := range list {
		tmpKey += string(current)
	}
	return tmpKey + identifier
}
