package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const numberOfArguments int = 2
const up string = "U"
const down string = "D"
const seaLevel = 0

func countingValleys(n int, s string) int {
	currentLevel := seaLevel
	valleyCount := 0
	for _, val := range s {
		currentAction := strings.ToUpper(string(rune(val)))
		if currentAction == up {
			currentLevel++
			if currentLevel == seaLevel {
				valleyCount++
			}

		} else if currentAction == down {
			currentLevel--
		}

	}
	return valleyCount
}

func main() {

	if len(os.Args) != numberOfArguments+1 {
		log.Fatal("Invalid number of Arguments")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	series := os.Args[2]

	valleyCount := countingValleys(n, series)
	fmt.Println("Number of Valleys ", valleyCount)

}
