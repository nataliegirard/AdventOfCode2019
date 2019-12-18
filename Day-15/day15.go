package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type coordinate struct {
	x, y int
}

type cell struct {
	marking string
	steps   int
}

func convertInput(input string) []int64 {
	arr := strings.Split(input, ",")

	values := make([]int64, 100000)
	for i := 0; i < len(arr); i++ {
		values[i], _ = strconv.ParseInt(arr[i], 10, 64)
	}
	return values
}

func parseInstruction(code int64) (int64, int64, int64, int64) {
	opcode := code % 100
	mode1 := (code / 100) % 10
	mode2 := (code / 1000) % 10
	mode3 := (code / 10000) % 10

	return opcode, mode1, mode2, mode3
}

func getValue(values []int64, mode int64, index int64, relativeBase int64) int64 {
	if mode == 1 {
		return index
	}
	if mode == 2 {
		return values[index+relativeBase]
	}
	return values[index]
}

func Intcode(values []int64, message chan int64) int64 {
	var i int64 = 0
	var lastOutput int64 = 0
	var relativeBase int64 = 0
	endloop := false
	for {
		opcode, mode1, mode2, mode3 := parseInstruction(values[i])
		//fmt.Println(i, opcode)
		switch opcode {
		case 1:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			value2 := getValue(values, mode2, values[i+2], relativeBase)
			index3 := values[i+3]
			if mode3 == 2 {
				index3 += relativeBase
			}
			values[index3] = value1 + value2
			i += 4
		case 2:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			value2 := getValue(values, mode2, values[i+2], relativeBase)
			index3 := values[i+3]
			if mode3 == 2 {
				index3 += relativeBase
			}
			values[index3] = value1 * value2
			i += 4
		case 3:
			index := values[i+1]
			inputValue := <-message
			if mode1 == 2 {
				index += relativeBase
			}
			values[index] = inputValue
			i += 2
		case 4:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			//fmt.Println("Output:", value1)
			lastOutput = value1
			message <- value1
			i += 2
		case 5:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			value2 := getValue(values, mode2, values[i+2], relativeBase)
			if value1 != 0 {
				i = value2
			} else {
				i += 3
			}
		case 6:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			value2 := getValue(values, mode2, values[i+2], relativeBase)
			if value1 == 0 {
				i = value2
			} else {
				i += 3
			}
		case 7:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			value2 := getValue(values, mode2, values[i+2], relativeBase)
			index3 := values[i+3]
			if mode3 == 2 {
				index3 += relativeBase
			}

			if value1 < value2 {
				values[index3] = 1
			} else {
				values[index3] = 0
			}
			i += 4
		case 8:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			value2 := getValue(values, mode2, values[i+2], relativeBase)
			index3 := values[i+3]
			if mode3 == 2 {
				index3 += relativeBase
			}

			if value1 == value2 {
				values[index3] = 1
			} else {
				values[index3] = 0
			}
			i += 4
		case 9:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			relativeBase += value1
			i += 2
		case 99:
			close(message)
			<-message
			endloop = true
		}

		if endloop {
			time.Sleep(100 * time.Millisecond)
			break
		}
	}
	return lastOutput
}

func find(a []int, x int) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func moveInDirection(coord coordinate, direction int) coordinate {
	newCoord := coord

	switch direction {
	case 1:
		newCoord.y++
	case 2:
		newCoord.y--
	case 3:
		newCoord.x--
	case 4:
		newCoord.x++
	}

	return newCoord
}

func getNewDirection(area map[coordinate]cell, curr coordinate, direction int) (int, coordinate) {
	moves := []int{1, 4, 2, 3}
	dir := find(moves, direction)
	dir = (dir + 1) % 4 // try to move to the right
	coord := moveInDirection(curr, moves[dir])

	i := 0
	for area[coord].marking == "#" && i < 4 {
		//fmt.Println(area[coord], coord, moves[dir], "dir", dir)
		dir = (dir - 1) % 4
		if dir < 0 {
			dir = dir + 4
		}
		//fmt.Println("dir:", dir)
		coord = moveInDirection(curr, moves[dir])
		i++
	}
	//fmt.Println("New direction", moves[dir], coord, "'", area[coord], "'")
	return moves[dir], coord
}

func printArea(area map[coordinate]cell) [][]string {
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for key := range area {
		if key.x < minX {
			minX = key.x
		}
		if key.x > maxX {
			maxX = key.x
		}
		if key.y < minY {
			minY = key.y
		}
		if key.y > maxY {
			maxY = key.y
		}
	}

	room := make([][]string, maxY-minY+1)
	for i := 0; i < maxY-minY+1; i++ {
		room[i] = make([]string, maxX-minX+1)
	}

	for key, value := range area {
		if value.marking == "#" {
			room[key.y-minY][key.x-minX] = value.marking
		} else {
			room[key.y-minY][key.x-minX] = strconv.Itoa(value.steps)
		}
	}

	f, _ := os.Create("output.csv")
	defer f.Close()
	for i := 0; i < maxY-minY+1; i++ {
		comma := strings.Join(room[i], ",")
		fmt.Fprintln(f, comma)
	}

	return room
}

func floodMaze(maze map[coordinate]cell, start coordinate) {
	currSteps := maze[start].steps
	c := start
	t := maze[c]

	// north
	c.y--
	t = maze[c]
	if t.marking != "#" {
		if t.steps == -1 || t.steps > currSteps+1 {
			t.steps = currSteps + 1
			maze[c] = t
			floodMaze(maze, c)
		}
	}
	c.y++

	// south
	c.y++
	t = maze[c]
	if t.marking != "#" {
		if t.steps == -1 || t.steps > currSteps+1 {
			t.steps = currSteps + 1
			maze[c] = t
			floodMaze(maze, c)
		}
	}
	c.y--

	// west
	c.x--
	t = maze[c]
	if t.marking != "#" {
		if t.steps == -1 || t.steps > currSteps+1 {
			t.steps = currSteps + 1
			maze[c] = t
			floodMaze(maze, c)
		}
	}
	c.x++

	// east
	c.x++
	t = maze[c]
	if t.marking != "#" {
		if t.steps == -1 || t.steps > currSteps+1 {
			t.steps = currSteps + 1
			maze[c] = t
			floodMaze(maze, c)
		}
	}
	c.x--
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	str = strings.Replace(str, "\n", "", -1)
	values := convertInput(str)

	message := make(chan int64)
	go Intcode(values, message)

	area := make(map[coordinate]cell)
	output := 1
	var curr coordinate
	var start coordinate
	var origin coordinate
	direction := 0
	i := 0
	for i < 200000 {
		//fmt.Println("(", i, ")", output, "in direction", direction, curr)
		//store map information based on received output
		var c cell
		c.steps = -1
		switch direction {
		case 1:
			if output == 0 {
				c.marking = "#"
				area[curr] = c
				curr.y--
				direction = 2
			} else {
				c.marking = "."
				area[curr] = c
			}
		case 2:
			if output == 0 {
				c.marking = "#"
				area[curr] = c
				curr.y++
				direction = 1
			} else {
				c.marking = "."
				area[curr] = c
			}
		case 3:
			if output == 0 {
				c.marking = "#"
				area[curr] = c
				curr.x++
				direction = 4
			} else {
				c.marking = "."
				area[curr] = c
			}
		case 4:
			if output == 0 {
				c.marking = "#"
				area[curr] = c
				curr.x--
				direction = 3
			} else {
				c.marking = "."
				area[curr] = c
			}
		default:
			c.marking = "."
			area[curr] = c
		}

		if output == 1 && i == 0 {
			c.marking = "o"
			area[curr] = c
			origin = curr
		}
		if output == 2 {
			c = cell{"*", 0}
			start = curr
			area[curr] = c
		}

		newDir, coord := getNewDirection(area, curr, direction)
		curr = coord
		direction = newDir
		message <- int64(newDir)

		receive, ok := <-message
		if !ok {
			break
		}
		output = int(receive)
		i++
	}

	floodMaze(area, start)
	printArea(area)

	maxSteps := 0
	for _, value := range area {
		if value.steps > maxSteps {
			maxSteps = value.steps
		}
	}
	fmt.Println("Part 1", area[origin].steps, "\nPart 2", maxSteps) // part 1: 300, part 2: 312
}
