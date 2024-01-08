package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Printf("readfile: %v\n", err)
		os.Exit(1)
	}

	inputStrings := strings.Split(string(input), "\n")
	niceStringsCount := 0

	for _, str := range inputStrings {
		if IsNiceStringVersionTwo(str) {
			niceStringsCount++
		}
	}

	fmt.Println(niceStringsCount)
}

func isNiceString(input string) bool {
	vowels := "aeiou"
	forbiddenStrings := []string{"ab", "cd", "pq", "xy"}

	vowelCount := 0
	hasLetterTwiceInARow := false

	var currentLetter string

	for _, ch := range input {
		letter := string(ch)

		currentSubstring := currentLetter + letter
		for _, substring := range forbiddenStrings {
			if currentSubstring == substring {
				return false
			}
		}

		if currentLetter == letter {
			hasLetterTwiceInARow = true
		}

		if strings.Contains(vowels, letter) {
			vowelCount++
		}

		currentLetter = letter
	}

	if vowelCount > 2 && hasLetterTwiceInARow {
		return true
	}

	return false
}

func IsNiceStringVersionTwo(input string) bool {
	letterPairs := make(map[string]int)

	var currentLetter, currentPair string

	currentTriplet := [3]string{"", "", ""}

	var hasConditionalTriplet, hasRepeatingPairs bool

	for _, ch := range input {
		letter := string(ch)

		if !hasRepeatingPairs {
			newPair := currentLetter + letter

			if len(newPair) == 2 {
				if newPair != currentPair {
					if _, ok := letterPairs[newPair]; ok {
						hasRepeatingPairs = true
					} else {
						letterPairs[newPair] = 1
					}
				} else {
					if _, ok := letterPairs[strings.ToUpper(newPair)]; ok {
						hasRepeatingPairs = true
					} else {
						letterPairs[strings.ToUpper(newPair)] = 1

					}
				}

			}

			currentLetter = letter
			currentPair = newPair
		}

		if !hasConditionalTriplet {
			currentTriplet[0], currentTriplet[1], currentTriplet[2] = currentTriplet[1], currentTriplet[2], letter
			if currentTriplet[0] == currentTriplet[2] {
				hasConditionalTriplet = true
			}
		}

		if hasConditionalTriplet && hasRepeatingPairs {
			return true
		}
	}

	return false
}
