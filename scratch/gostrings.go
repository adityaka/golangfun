package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal(os.Args[0] + " < some string > ")
	}

	for idx, value := range os.Args[1] {
		fmt.Println("Char at", idx, " = ", string(rune(value)))
		fmt.Printf("Value at index %d = %c(%v)\n", idx, value, value)
	}

}
