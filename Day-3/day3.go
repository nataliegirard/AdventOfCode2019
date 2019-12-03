package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMax(line []string) (int, int, int, int) {
	posHorz := 0
	posVert := 0
	maxUp := 0
	maxDown := 0
	maxLeft := 0
	maxRight := 0

	for _, val := range line {
		switch val[0] {
		case 'U':
			t, _ := strconv.Atoi(val[1:])
			posVert += t
			if posVert > maxUp {
				maxUp = posVert
			}
		case 'D':
			t, _ := strconv.Atoi(val[1:])
			posVert -= t
			if posVert < maxDown {
				maxDown = posVert
			}
		case 'R':
			t, _ := strconv.Atoi(val[1:])
			posHorz += t
			if posHorz > maxRight {
				maxRight = posHorz
			}
		case 'L':
			t, _ := strconv.Atoi(val[1:])
			posHorz -= t
			if posHorz < maxLeft {
				maxLeft = posHorz
			}
		}
	}
	return maxUp, maxDown, maxRight, maxLeft
}

func getDimensions(line1 []string, line2 []string) (int, int, int, int) {
	fmt.Println(line1)
	fmt.Println(line2)

	maxUp1, maxDown1, maxRight1, maxLeft1 := findMax(line1)
	maxUp2, maxDown2, maxRight2, maxLeft2 := findMax(line2)

	maxUp := maxInt(maxUp1, maxUp2)
	maxDown := minInt(maxDown1, maxDown2)
	maxRight := maxInt(maxRight1, maxRight2)
	maxLeft := minInt(maxLeft1, maxLeft2)

	fmt.Println("position values", maxUp, maxDown, maxLeft, maxRight)

	return maxRight - maxLeft + 1, maxUp - maxDown + 1, maxLeft, maxUp
}

func main() {
	line1 := "R8,U5,L5,D3"
	line2 := "U7,R6,D4,L4"
	lineA := strings.Split(line1, ",")
	lineB := strings.Split(line2, ",")
	horz, vert, startx, starty := getDimensions(lineA, lineB)
	fmt.Println("Horizontal", horz, "Vertical", vert, "origin", startx, starty)
}
