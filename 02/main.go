package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println(err)
	}

	var totalWrappingPaper, totalRibbon int

	var presentDimensions = strings.Split(string(input), "\n")

	for _, pd := range presentDimensions {
		var dimensions = strings.Split(pd, "x")

		var l, _ = strconv.Atoi(dimensions[0])
		var w, _ = strconv.Atoi(dimensions[1])
		var h, _ = strconv.Atoi(dimensions[2])

		surfaceArea := getSurfaceArea(l, w, h)
		slackPaper := getSmallestSide(l, w, h)
		necessaryPaper := surfaceArea + slackPaper
		totalWrappingPaper += necessaryPaper

		wrappingRibbon := getWrappingRibbon(l, w, h)
		bowRibbon := l * w * h
		totalRibbon += wrappingRibbon + bowRibbon
	}

	fmt.Println(totalWrappingPaper)
	fmt.Println(totalRibbon)
}

func getSurfaceArea(l, w, h int) int {
	return 2*l*w + 2*w*h + 2*h*l
}

func getSmallestSide(l, w, h int) int {
	x := l * w
	y := w * h
	z := h * l

	if x <= y && x <= z {
		return x
	} else if y <= x && y <= z {
		return y
	} else {
		return z
	}
}

func getWrappingRibbon(l, w, h int) int {
	if l >= w && l >= h {
		return 2*w + 2*h
	} else if w >= h && w >= l {
		return 2*h + 2*l
	} else {
		return 2*l + 2*w
	}
}
