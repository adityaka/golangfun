package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func repeatedString(s string, n int64) int64 {
	const searchChar string = "a"
	var seachCharCount int64 = 0
	for _, val := range s {
		if string(rune(val)) == searchChar {
			seachCharCount++
		}
	}
	var numberOfStrings int64 = n / int64(len(s))
	charCount := numberOfStrings * seachCharCount
	var remainderOfCharacters = n % int64(len(s))
	if remainderOfCharacters > 0 {
		for _, val := range s[:remainderOfCharacters] {
			if string(rune(val)) == searchChar {
				charCount++
			}
		}
	}
	return charCount
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Invalid arguments")
	}

	s := os.Args[1]
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	strCount := repeatedString(s, int64(n))
	fmt.Println(strCount)

}
