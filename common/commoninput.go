package common

import (
	"errors"
	"os"
)

//InputProcessor...
func InputProcessor(numberOfArgs int) ([]string, error) {
	if numberOfArgs+1 < len(os.Args) {
		return nil, errors.New("Invalid number of arguments")
	}
	return os.Args, nil
}
