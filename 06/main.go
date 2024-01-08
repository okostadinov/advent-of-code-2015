package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var commands = map[string]string{
	"on":     "turn on",
	"off":    "turn off",
	"toggle": "toggle",
}

func main() {
	input, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grid := [1000][1000]int{}

	instructions := strings.Split(string(input), "\n")

	for _, instruction := range instructions {
		handleInstructionA(&grid, instruction)
	}

	var lightsOnAmount int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				lightsOnAmount++
			}
		}
	}

	fmt.Println(lightsOnAmount)

	grid2 := [1000][1000]int{}

	for _, instruction := range instructions {
		handleInstructionB(&grid2, instruction)
	}

	var totalLightsOutput int

	for i := 0; i < len(grid2); i++ {
		for j := 0; j < len(grid2[i]); j++ {
			totalLightsOutput += grid2[i][j]
		}
	}

	fmt.Println(totalLightsOutput)
}

func handleInstructionA(grid *[1000][1000]int, instruction string) {
	command := getCommand(instruction)

	startX, endX, startY, endY := getCoords(instruction)

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			switch command {
			case "turn on":
				grid[x][y] = 1
			case "turn off":
				grid[x][y] = 0
			case "toggle":
				grid[x][y] = toggleLightA(grid[x][y])
			}
		}
	}
}

func toggleLightA(light int) int {
	if light == 0 {
		return 1
	}

	return 0
}

func handleInstructionB(grid *[1000][1000]int, instruction string) {
	command := getCommand(instruction)

	startX, endX, startY, endY := getCoords(instruction)

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			switch command {
			case "turn on":
				grid[x][y]++
			case "turn off":
				if grid[x][y] > 0 {
					grid[x][y]--
				}
			case "toggle":
				grid[x][y] += 2
			}
		}
	}
}

func getCommand(instruction string) (command string) {
	for _, comm := range commands {
		if strings.HasPrefix(instruction, comm) {
			command = comm
		}
	}

	return
}

func getCoords(instruction string) (startX, endX, startY, endY int) {
	instructionsArr := strings.Split(instruction, " ")
	startLightCoords := strings.Split(instructionsArr[len(instructionsArr)-3], ",")
	endLightCoords := strings.Split(instructionsArr[len(instructionsArr)-1], ",")

	startXStr, startYStr := startLightCoords[0], startLightCoords[1]
	endXStr, endYStr := endLightCoords[0], endLightCoords[1]

	startX, _ = strconv.Atoi(startXStr)
	endX, _ = strconv.Atoi(endXStr)
	startY, _ = strconv.Atoi(startYStr)
	endY, _ = strconv.Atoi(endYStr)

	return
}
