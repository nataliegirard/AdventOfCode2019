package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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

func Intcode(values []int64, inputValue int64) int64 {
	var i int64 = 0
	var lastOutput int64 = 0
	var relativeBase int64 = 0
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
			if mode1 == 2 {
				index += relativeBase
			}
			values[index] = inputValue
			i += 2
		case 4:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			fmt.Println("Output:", value1)
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

	//str := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	//str := "1102,34915192,34915192,7,4,7,99,0" // Outputs: 1219070632396864
	//str := "104,1125899906842624,99"

	values := convertInput(str)
	result := Intcode(values, 2)
	fmt.Println("Result:", result)
	// Part 1: 2752191671
	// Part 2: 87571
}
