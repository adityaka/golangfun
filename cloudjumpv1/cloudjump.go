package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
)

const(
	cumulusCloud = 0
      thunderCloud = 1
      )


func cloudJump(c []int32) int32 {
	count  := len(c)
	jumpCounter := 0
	currentPos := 0
	for idx := range c {
		if currentPos >= count -1 {
			break
		}
		if idx+2 <=  count - 1  && c[idx+1] == thunderCloud  && c[idx+2] == cumulusCloud {
			jumpCounter++
			currentPos+=2
		} else if idx + 2 <= count -1 && c[idx+1] == cumulusCloud && c[idx+2] == cumulusCloud {
			jumpCounter++
			currentPos+=2
		} else if idx+1
	}

}

func checkError(err error){
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
func main() {
	if len(os.Args) != 3 {
		panic("Invalid number of arguments")
	}
	countString := os.Args[1]
	itemsString := os.Args[2]
	count,err :=  strconv.ParseInt(countString, 10, 64)
	checkError(err)
	items := strings.Split(itemsString, " ")
	if len(items) != int(count) {
		panic(errors.New("Invalid argument count doesn't match length of input"))
	}
	input := make([]int32, count)
	for idx, val := range items {
		intVal,err  := strconv.Atoi(val)
		checkError(err)
		input[idx] = int32(intVal)
	}
	// call your code
	result := cloudJump(input)
	fmt.Println(result)
}
