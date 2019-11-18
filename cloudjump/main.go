package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const thunderCloud int32 = 1
const cumulusCloud int32 = 0

func isThunderCloud(cloudValue int32) bool {
	if cloudValue == thunderCloud {
		return true
	}
	return false
}

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	return jumpingOnCloudsInternalv1(c)
}

func jumpingOnCloudsInternalv1(c []int32) int32 {
	fmt.Println("input length : ", len(c))
	var jumpCounter int32 = 0
	var currentPos int32 = 0
	for idx := range c {
		if int32(idx) < currentPos {
			continue
		}

		if int(currentPos) >= len(c) {
			break
		}
		nextHop := idx + 1
		nextTwoHops := idx + 2
		if nextHop <= len(c)-1 {
			if c[nextHop] == cumulusCloud {
				jumpCounter++
				if nextTwoHops < len(c)-1 {
					if c[nextTwoHops] == cumulusCloud {
						currentPos += 2
					} else if c[nextTwoHops] == thunderCloud {
						currentPos++
					}
				}
			} else if c[nextHop] == thunderCloud {
				jumpCounter++
				currentPos += 2
			}
		}
	}
	return jumpCounter
}

func jumpingOnCloudsInternal(c []int32) int32 {
	fmt.Println("input length : ", len(c))
	var jumpCounter int32 = 0
	var currentPos int32 = 0
	for idx := range c {
		if int32(idx) < currentPos {
			continue
		}
		nextHop := idx + 1
		nextTwoHops := idx + 2
		if nextHop <= len(c)-1 {
			if c[nextHop] == cumulusCloud {
				jumpCounter++
				if nextTwoHops < len(c)-1 {
					if c[nextTwoHops] == cumulusCloud {
						currentPos += 2
					} else if c[nextTwoHops] == thunderCloud {
						currentPos++
					}
				}
			} else if c[nextHop] == thunderCloud {
				jumpCounter++
				currentPos += 2
			}
		}
	}
	return jumpCounter

}

func main() {

	count, err := strconv.ParseInt(os.Args[1], 10, 64)
	checkError(err)
	valuesString := os.Args[2]
	checkError(err)
	values := strings.Split(valuesString, " ")
	if len(values) != int(count) {
		err = errors.New("Invalid input")
		checkError(err)
	}

	intVals := make([]int32, count)
	for idx, val := range values {
		intVal, err := strconv.ParseInt(val, 10, 32)
		checkError(err)
		intVals[idx] = int32(intVal)
	}

	result := jumpingOnClouds(intVals)
	fmt.Println(result)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
