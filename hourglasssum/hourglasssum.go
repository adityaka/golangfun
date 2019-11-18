package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func hourGlassSum(arr [][]int32) int32 {
	const maxLen int = 6
	const maxHourGlassLen int = 3
	hgSum := int32(0)
	maxHgSum := int32(-2147483647)
	for i := 0; i < maxLen; i++ {

		for j := 0; j < maxLen; j++ {
			maxColIndex := j + 2
			maxRowIndex := i + 2
			if maxColIndex >= maxLen || maxRowIndex >= maxLen {
				break
			}

			hgSum += arr[i][j] + arr[i][j+1] + arr[i][j+2]
			hgSum += arr[i+1][j+1]
			hgSum += arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if hgSum > maxHgSum {
				maxHgSum = hgSum
			}
			hgSum = 0
		}
	}
	return maxHgSum
}

func main() {
	if len(os.Args) != 7 {
		log.Fatal("Invalid Arguments")
	}
	var arr [][]int32

	for idx, val := range os.Args {
		if idx == 0 {
			continue
		}

		elements := strings.Split(val, " ")
		if len(elements) != 6 {
			log.Fatal("row size should be 6")
		}
		var row []int32
		for i := 0; i < len(elements); i++ {
			intElement, err := strconv.Atoi(elements[i])
			if err != nil {
				panic(err)
			}
			row = append(row, int32(intElement))

		}
		arr = append(arr, row)
	}
	sumOfHourGlasses := hourGlassSum(arr)
	output, exists := os.LookupEnv("OUTPUT")
	if exists {
		intVal, err := strconv.Atoi(output)
		if err != nil {
			panic(err)
		}
		if int32(intVal) != sumOfHourGlasses {
			log.Println("TEST CASE FAILED Expected output ", output)
		}
	} else {
		log.Println("No test cases output defined")
	}
	fmt.Println("Max Sum Of hourglasses is = ", sumOfHourGlasses)
}
