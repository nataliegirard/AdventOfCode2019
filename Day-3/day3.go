package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func maxInt(a int, b int) int {
	if a > b {
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
	return maxUp, maxDown * -1, maxRight, maxLeft * -1
}

func getDimensions(line1 []string, line2 []string) (int, int, int, int) {
	maxUp1, maxDown1, maxRight1, maxLeft1 := findMax(line1)
	maxUp2, maxDown2, maxRight2, maxLeft2 := findMax(line2)

	maxUp := maxInt(maxUp1, maxUp2)
	maxDown := maxInt(maxDown1, maxDown2)
	maxRight := maxInt(maxRight1, maxRight2)
	maxLeft := maxInt(maxLeft1, maxLeft2)

	return maxRight + maxLeft + 1, maxUp + maxDown + 1, maxLeft, maxUp
}

func readFile(filename string) ([]string, []string) {
	lines := make([][]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()
		line := strings.Split(l, ",")
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines[0], lines[1]
}

func printGrid(grid [][]string) {
	f, _ := os.Create("output")
	defer f.Close()
	for i := 0; i < len(grid); i++ {
		fmt.Fprintln(f, grid[i])
	}
}

func addLine(grid [][]string, posx int, posy int, line []string, marker string) {
	prev := "."
	for _, val := range line {
		switch val[0] {
		case 'U':
			t, _ := strconv.Atoi(val[1:])
			for i := 0; i < t; i++ {
				posy--
				if grid[posy][posx] != "." && grid[posy][posx] != marker {
					grid[posy][posx] = "X"
				} else {
					grid[posy][posx] = marker
				}
			}
		case 'D':
			t, _ := strconv.Atoi(val[1:])
			for i := 0; i < t; i++ {
				posy++
				if grid[posy][posx] != "." && grid[posy][posx] != marker {
					grid[posy][posx] = "X"
				} else {
					grid[posy][posx] = marker
				}
			}
		case 'R':
			t, _ := strconv.Atoi(val[1:])
			for i := 0; i < t; i++ {
				posx++
				if grid[posy][posx] != "." && grid[posy][posx] != marker {
					grid[posy][posx] = "X"
				} else {
					grid[posy][posx] = marker
				}
			}
		case 'L':
			t, _ := strconv.Atoi(val[1:])
			for i := 0; i < t; i++ {
				posx--
				if grid[posy][posx] != "." && grid[posy][posx] != marker {
					grid[posy][posx] = "X"
				} else {
					grid[posy][posx] = marker
				}
			}
		}
		prev = grid[posy][posx]
		grid[posy][posx] = "+"
	}
	grid[posy][posx] = prev
}

func distanceTwoPoints(ax int, ay int, bx int, by int) int {
	diffx := ax - bx
	diffy := ay - by
	if diffx < 0 {
		diffx *= -1
	}
	if diffy < 0 {
		diffy *= -1
	}

	return diffx + diffy
}

func findIntersectionDistance(grid [][]string, ox int, oy int) int {
	minDistance := 10000

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != "X" {
				continue
			}
			d := distanceTwoPoints(ox, oy, j, i)
			if d < minDistance {
				minDistance = d
			}
		}
	}
	return minDistance
}

func runProgram(lineA []string, lineB []string) int {
	horz, vert, startx, starty := getDimensions(lineA, lineB)

	grid := make([][]string, vert)
	for i := 0; i < vert; i++ {
		grid[i] = make([]string, horz)
		for j := 0; j < horz; j++ {
			grid[i][j] = "."
		}
	}

	grid[starty][startx] = "o"
	addLine(grid, startx, starty, lineA, "a")
	addLine(grid, startx, starty, lineB, "b")
	//printGrid(grid)
	result := findIntersectionDistance(grid, startx, starty)
	return result
}

func main() {
	/*line1A, line1B := readFile("ex1.txt")
	result1 := runProgram(line1A, line1B)
	fmt.Println("ex1:", result1, "expected 6")

	line2A, line2B := readFile("ex2.txt")
	result2 := runProgram(line2A, line2B)
	fmt.Println("ex2:", result2, "expected 159")

	line3A, line3B := readFile("ex3.txt")
	result3 := runProgram(line3A, line3B)
	fmt.Println("ex3:", result3, "expected 135")*/

	lineA, lineB := readFile("input.txt")
	result := runProgram(lineA, lineB)
	fmt.Println("Part 1:", result) // 1285
}
