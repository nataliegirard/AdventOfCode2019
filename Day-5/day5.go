package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func convertInput(input string) []int {
	arr := strings.Split(input, ",")

	values := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		values[i], _ = strconv.Atoi(arr[i])
	}
	return values
}

func parseInstruction(code int) (int, int, int, int) {
	opcode := code % 100
	mode1 := (code / 100) % 10
	mode2 := (code / 1000) % 10
	mode3 := (code / 10000) % 10

	return opcode, mode1, mode2, mode3
}

func getValue(values []int, mode int, index int) int {
	if mode == 1 {
		return index
	}
	return values[index]
}

func Intcode(values []int, inputValue int) int {
	i := 0
	lastOutput := 0
	for {
		opcode, mode1, mode2, _ := parseInstruction(values[i])
		fmt.Println(i, opcode)
		switch opcode {
		case 1:
			value1 := getValue(values, mode1, values[i+1])
			value2 := getValue(values, mode2, values[i+2])
			index3 := values[i+3]
			values[index3] = value1 + value2
			i += 4
		case 2:
			value1 := getValue(values, mode1, values[i+1])
			value2 := getValue(values, mode2, values[i+2])
			index3 := values[i+3]
			values[index3] = value1 * value2
			i += 4
		case 3:
			input := inputValue
			index := values[i+1]
			values[index] = input
			i += 2
		case 4:
			value1 := getValue(values, mode1, values[i+1])
			fmt.Println("Output:", value1)
			lastOutput = value1
			i += 2
		case 5:
			value1 := getValue(values, mode1, values[i+1])
			value2 := getValue(values, mode2, values[i+2])
			if value1 != 0 {
				i = value2
			} else {
				i += 3
			}
		case 6:
			value1 := getValue(values, mode1, values[i+1])
			value2 := getValue(values, mode2, values[i+2])
			if value1 == 0 {
				i = value2
			} else {
				i += 3
			}
		case 7:
			value1 := getValue(values, mode1, values[i+1])
			value2 := getValue(values, mode2, values[i+2])
			index3 := values[i+3]
			if value1 < value2 {
				values[index3] = 1
			} else {
				values[index3] = 0
			}
			i += 4
		case 8:
			value1 := getValue(values, mode1, values[i+1])
			value2 := getValue(values, mode2, values[i+2])
			index3 := values[i+3]
			if value1 == value2 {
				values[index3] = 1
			} else {
				values[index3] = 0
			}
			i += 4
		case 99:
			return lastOutput
		default:
			return -1
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
	result := Intcode(values, 1)
	fmt.Println("Part 1:", result) // 13547311

	values = convertInput(str)
	part2 := Intcode(values, 5)
	fmt.Println("Part 2:", part2) // 236453

	/*testInput := convertInput("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99")
	fmt.Println("TestInput", testInput)
	opcode, mode1, mode2, mode3 := parseInstruction(testInput[0]) // all ints
	fmt.Println("parsed instruction:", opcode, mode1, mode2, mode3)
	value1 := getValue(testInput, mode1, testInput[1])
	fmt.Println("First value", value1)
	value2 := getValue(testInput, mode2, testInput[2])
	fmt.Println("Second value", value2)
	testResult := Intcode(testInput, 8)
	fmt.Println("Answer", testResult)*/
}
