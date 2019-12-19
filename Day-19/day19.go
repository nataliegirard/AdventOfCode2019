package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

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

func Intcode(values []int64, x int, y int) int64 {
	var i int64 = 0
	var lastOutput int64 = 0
	var relativeBase int64 = 0
	endloop := false
	inputCount := 0
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
			var inputValue int64
			if inputCount == 0 {
				inputValue = int64(x)
				inputCount++
			} else {
				inputValue = int64(y)
			}
			//fmt.Println("input:", inputValue)
			if mode1 == 2 {
				index += relativeBase
			}
			values[index] = inputValue
			i += 2
		case 4:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			//fmt.Println("Output:", value1)
			lastOutput = value1
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
			endloop = true
		}

		if endloop {
			time.Sleep(100 * time.Millisecond)
			break
		}
	}
	return lastOutput
}

func printArea(area [][]int64, toScreen bool) {
	if toScreen {
		for i := 0; i < len(area); i++ {
			fmt.Println(area[i])
		}
	} else {
		f, _ := os.Create("output.csv")
		defer f.Close()
		for i := 0; i < len(area); i++ {
			for j := 0; j < len(area[i]); j++ {
				fmt.Fprintf(f, "%d,", area[i][j])
			}
			fmt.Fprintf(f, "\n")
		}
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

	/*xStart := 35
	yStart := 52
	xSize := 50
	ySize := 10
	space := make([][]int64, ySize)
	for i := 0; i < ySize; i++ {
		space[i] = make([]int64, xSize)
	}

	count := 0
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			values = convertInput(str)
			output := Intcode(values, x+xStart, y+yStart)

			space[y][x] = output
			if output == 1 {
				count++
			}
		}
	}

	fmt.Println(count) // part 1: 189
	printArea(space, false)*/

	// Part 2
	// start 35,52
	// if x+99,y != 1 -> y++
	// else if x,y+99 != 1 -> x++
	// else, should fit 100x100 in beam
	x := 35
	y := 52
	size := 99
	var output int64
	for {
		//fmt.Println(x, y)
		values = convertInput(str)
		output = Intcode(values, x, y)
		if output != 1 {
			x++
			continue
		}

		values = convertInput(str)
		output = Intcode(values, x+size, y)
		if output != 1 {
			y++
			continue
		}

		values = convertInput(str)
		output = Intcode(values, x, y+size)
		if output != 1 {
			x++
			continue
		}

		break
	}

	fmt.Println("Solution:", x, y, x*10000+y) // Part 2: 7621042
}
