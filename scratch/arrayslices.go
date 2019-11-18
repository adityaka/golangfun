package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

//ValidateProgramArgumentCount - validates the program argument count
func ValidateProgramArgumentCount(numberOfArguments int) (*[]string, error) {
	if numberOfArguments+1 < len(os.Args) {
		return nil, errors.New("Invalid number of Arguments")
	}

	return &os.Args[1:len(os.Args)], nil
}

func main() {

	args, err := ValidateProgramArgumentCount(2)
	if err != nil {
		log.Fatal(err)
	}

	for arg := range *args {
		fmt.Println(arg)
	}

}
