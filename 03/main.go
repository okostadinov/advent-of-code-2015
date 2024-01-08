package main

import (
	"fmt"
	"io/ioutil"
)

type houseCoords struct {
	x int
	y int
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	part1(input)
	part2(input)
}

func part1(input []byte) {
	visitedHouses := make([]houseCoords, 1)
	currentPosition := houseCoords{
		x: 0,
		y: 0,
	}

	for _, direction := range input {

		currentPosition = updateCoords(direction, currentPosition)

		if !(hasBeenVisited(visitedHouses, currentPosition)) {
			visitedHouses = append(visitedHouses, currentPosition)
		}
	}

	fmt.Println(len(visitedHouses))
}

func part2(input []byte) {
	santaCoords := houseCoords{
		x: 0,
		y: 0,
	}

	robotCoords := houseCoords{
		x: 0,
		y: 0,
	}

	visitedHouses := make([]houseCoords, 1)

	for i := 0; i < len(input); i++ {
		direction := input[i]

		if i%2 == 0 {
			santaCoords = updateCoords(direction, santaCoords)
			if !(hasBeenVisited(visitedHouses, santaCoords)) {
				visitedHouses = append(visitedHouses, santaCoords)
			}
		} else {
			robotCoords = updateCoords(direction, robotCoords)
			if !(hasBeenVisited(visitedHouses, robotCoords)) {
				visitedHouses = append(visitedHouses, robotCoords)
			}
		}
	}

	fmt.Println(len(visitedHouses))
}

func hasBeenVisited(houses []houseCoords, coords houseCoords) bool {
	var hasBeenVisited bool

	for _, house := range houses {
		if house == coords {
			hasBeenVisited = true
		}
	}

	return hasBeenVisited
}

func updateCoords(direction byte, coords houseCoords) houseCoords {
	if direction == '>' {
		coords.x += 1
	} else if direction == '<' {
		coords.x -= 1
	} else if direction == '^' {
		coords.y += 1
	} else {
		coords.y -= 1
	}

	return coords
}
