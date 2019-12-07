package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
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

func Intcode(values []int, inputValue int, secondInput int) int {
	i := 0
	lastOutput := 0
	inputCount := 0
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
			index := values[i+1]
			if inputCount == 0 {
				fmt.Println("Input 0:", inputValue)
				values[index] = inputValue
			} else {
				fmt.Println("Input 1:", secondInput)
				values[index] = secondInput
			}
			inputCount++
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

func getThrust(values []int, inputSequence []int, secondInput int) int {
	numAmps := len(inputSequence)

	for i := 0; i < numAmps; i++ {
		var program = make([]int, len(values))
		copy(program, values)
		secondInput = Intcode(program, inputSequence[i], secondInput)
		fmt.Println(i, "Second Input", secondInput)
		fmt.Println("")
	}
	return secondInput
}

//https://www.golangprograms.com/golang-program-to-generate-slice-permutations-of-number-entered-by-user.html
func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

func amplifier(values []int, id int, phaseSignal int, inbound chan int, outbound chan int) {
	i := 0
	lastOutput := 0
	inputCount := 0
	endloop := false
	for {
		opcode, mode1, mode2, _ := parseInstruction(values[i])

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
			index := values[i+1]
			if inputCount == 0 {
				fmt.Println(id, "Input 0:", phaseSignal)
				values[index] = phaseSignal
			} else {
				secondInput := <-inbound
				fmt.Println(id, "Input 1:", secondInput)
				values[index] = secondInput
			}
			inputCount++
			i += 2
		case 4:
			value1 := getValue(values, mode1, values[i+1])
			fmt.Println(id, "Output:", value1)
			lastOutput = value1
			outbound <- value1
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
			//outbound <- lastOutput
			fmt.Println(id, "closing", lastOutput)
			close(outbound)
			endloop = true
			temp := <-inbound
			inputCount = temp
		}

		if endloop {
			time.Sleep(100 * time.Millisecond)
			break
		}
	}
}

func feedback(program []int, phases []int) int {
	toA := make(chan int)
	AtoB := make(chan int)
	BtoC := make(chan int)
	CtoD := make(chan int)
	DtoE := make(chan int)
	fromE := make(chan int)

	var programA = make([]int, len(program))
	copy(programA, program)
	go amplifier(programA, 0, phases[0], toA, AtoB)

	var programB = make([]int, len(program))
	copy(programB, program)
	go amplifier(programB, 1, phases[1], AtoB, BtoC)

	var programC = make([]int, len(program))
	copy(programC, program)
	go amplifier(programC, 2, phases[2], BtoC, CtoD)

	var programD = make([]int, len(program))
	copy(programD, program)
	go amplifier(programD, 3, phases[3], CtoD, DtoE)

	var programE = make([]int, len(program))
	copy(programE, program)
	go amplifier(programE, 4, phases[4], DtoE, fromE)

	toA <- 0
	maxThrust := 0
	for {
		v, ok := <-fromE
		fmt.Println("Main", ok, v)
		if ok {
			toA <- v
			maxThrust = v
		} else {
			break
		}
	}

	fmt.Println("Max thrust", maxThrust)
	return maxThrust
}

func findMaxThrust(values []int, phases []int, part int) int {
	var program = make([]int, len(values))
	copy(program, values)

	maxThrust := 0
	if part == 2 {
		perms := permutation(phases)
		for i := 0; i < len(perms); i++ {
			result := feedback(program, perms[i])
			if result > maxThrust {
				maxThrust = result
			}
		}
		return maxThrust
	}

	perms := permutation(phases)
	for i := 0; i < len(perms); i++ {
		result := getThrust(program, perms[i], maxThrust)
		if result > maxThrust {
			maxThrust = result
		}
	}
	return maxThrust
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	str = strings.Replace(str, "\n", "", -1)
	values := convertInput(str)

	/*var phases = []int{0, 1, 2, 3, 4}
	result := findMaxThrust(values, phases,1)
	fmt.Println("Part 1:", result) // 17440*/

	var phases = []int{5, 6, 7, 8, 9}
	part2 := findMaxThrust(values, phases, 2)
	fmt.Println("Part 2:", part2) // 27561242

	/*testInput := convertInput("3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0")
	inputSequence := []int{1, 0, 4, 3, 2}
	testResult := getThrust(testInput, inputSequence,0)
	fmt.Println("Answer", testResult)*/

	/*testInput := convertInput("3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0")
	var phases = []int{0, 1, 2, 3, 4}
	testResult := findMaxThrust(testInput, phases,1)
	fmt.Println("Answer:", testResult)*/

	/*testInput := convertInput("3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10")
	var phases = []int{5, 6, 7, 8, 9}
	testResult := findMaxThrust(testInput, phases, 2)
	fmt.Println("Answer:", testResult)*/
}
