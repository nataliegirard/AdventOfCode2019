package main

import (
	"fmt"
	"io/ioutil"
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

func Intcode(values []int64, message chan int64, waiting chan bool) int64 {
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
			waiting <- true
			inputValue := <-message
			if mode1 == 2 {
				index += relativeBase
			}
			values[index] = inputValue
			i += 2
		case 4:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			fmt.Println("Output:", value1)
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
			close(waiting)
			<-message
			<-waiting
			endloop = true
		}

		if endloop {
			time.Sleep(100 * time.Millisecond)
			break
		}
	}
	return lastOutput
}

func playGame(message chan int64, waiting chan bool) int64 {
	var score int64
	var ball int64
	var paddle int64
	for {
		select {
		case _, ok := <-waiting:
			if !ok {
				return score
			}

			if paddle > ball {
				message <- int64(-1)
			} else if paddle < ball {
				message <- int64(1)
			} else {
				message <- int64(0)
			}
		case x := <-message:
			y := <-message
			value := <-message

			if x == -1 && y == 0 {
				score = value
			} else if value == 3 {
				paddle = x
			} else if value == 4 {
				ball = x
			}
		default:
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

	message := make(chan int64)
	waiting := make(chan bool)
	values[0] = int64(2) // For part 2
	go Intcode(values, message, waiting)

	/* Part 1:
	var outputs []int64
	count := 0
	i := 0
	for {
		output, ok := <-message

		if !ok {
			break
		}

		outputs = append(outputs, output)
		i++
		if i%3 == 0 {
			if output == 2 {
				count++
			}
		}
	}
	fmt.Println(outputs)
	fmt.Println(i, count) // Part 1: 236
	*/

	score := playGame(message, waiting)
	fmt.Println("Score:", score) // Part 2: 11040
}
