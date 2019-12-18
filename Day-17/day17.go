package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type coordinate struct {
	x, y int
}

type cell struct {
	maze  string
	steps int
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

func Intcode(values []int64, inputs []int64) [][]string {
	var i int64 = 0
	var relativeBase int64 = 0
	var scaffold [][]string
	var line []string
	inputIndex := 0
	row := 0
	index := 0
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
			inputValue := inputs[inputIndex]
			inputIndex++
			fmt.Println("**Asking for input**")
			if mode1 == 2 {
				index += relativeBase
			}
			values[index] = inputValue
			i += 2
		case 4:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			fmt.Println("Output:", value1)
			if value1 == 10 {
				row++
				index = 0
				scaffold = append(scaffold, line)
				line = make([]string, 0)
				fmt.Println("---")
			} else {
				val := int(value1)
				line = append(line, string(val))
				index++
			}
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
			fmt.Println("End of program")
			endloop = true
		}

		if endloop {
			time.Sleep(100 * time.Millisecond)
			break
		}
	}
	fmt.Println("return")
	return scaffold
}

func atIntersection(scaffold [][]string, x int, y int) bool {
	// Check up
	if y-1 >= 0 && scaffold[y-1][x] != "#" {
		return false
	}

	// Check down
	if y+1 < len(scaffold) && scaffold[y+1][x] != "#" {
		return false
	}

	// Check left
	if x-1 >= 0 && scaffold[y][x-1] != "#" {
		return false
	}

	// Check right
	if x+1 < len(scaffold[y]) && scaffold[y][x+1] != "#" {
		return false
	}

	return true
}

func printArea(scaffold [][]string) {
	for i := 0; i < len(scaffold); i++ {
		fmt.Println(scaffold[i])
	}
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	str = strings.Replace(str, "\n", "", -1)
	values := convertInput(str)

	emptyInput := make([]int64, 1)
	emptyInput[0] = int64(0)

	scaffold := Intcode(values, emptyInput)
	scaffold = scaffold[:len(scaffold)-1]
	printArea(scaffold)

	total := 0
	for y := 0; y < len(scaffold); y++ {
		for x := 0; x < len(scaffold[y]); x++ {
			if atIntersection(scaffold, x, y) {
				total += x * y
			}
		}
	}
	fmt.Println("Result:", total) // Part 1: 4044

	// Solution calculated by hand
	mainRoutine := []int64{65, 44, 66, 44, 65, 44, 67, 44, 65, 44, 66, 44, 67, 44, 66, 44, 67, 44, 66, 10}
	functionA := []int64{82, 44, 56, 44, 76, 44, 49, 48, 44, 76, 44, 49, 50, 44, 82, 44, 52, 10}
	functionB := []int64{82, 44, 56, 44, 76, 44, 49, 50, 44, 82, 44, 52, 44, 82, 44, 52, 10}
	functionC := []int64{82, 44, 56, 44, 76, 44, 49, 48, 44, 82, 44, 56, 10}
	videoFeed := []int64{110, 10}

	inputValues := mainRoutine
	inputValues = append(inputValues, functionA...)
	inputValues = append(inputValues, functionB...)
	inputValues = append(inputValues, functionC...)
	inputValues = append(inputValues, videoFeed...)

	values = convertInput(str)
	values[0] = int64(2)
	scaffold = Intcode(values, inputValues)
	printArea(scaffold) // Part 2:
}
