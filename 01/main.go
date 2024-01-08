package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println(err)
	}

	var floor, position int
	var positionStored bool

	for i, letter := range input {
		if letter == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 && positionStored == false {
			positionStored = true
			position = i + 1
		}
	}

	fmt.Println(floor)
	fmt.Println(position)
}
