package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getPattern(basePattern []int, index int, size int) []int {
	pattern := make([]int, size)

	count := 0 // number of times the basePattern pos was used
	pos := 0   // position in the basePattern
	i := 0
	first := true
	for {
		if i == size {
			break
		}
		repeats := index + 1
		if first {
			first = false
			count++
			continue
		}

		if count == repeats {
			pos = (pos + 1) % len(basePattern)
			count = 0
		} else {
			pattern[i] = basePattern[pos]
			count++
			i++
		}
	}

	return pattern
}

func getNewValue(pattern []int, signal []int) int {
	val := 0
	for i := 0; i < len(signal); i++ {
		val += (pattern[i] * signal[i])
	}

	val = val % 10
	if val < 0 {
		val *= -1
	}
	return val
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	str = strings.TrimSuffix(str, "\n")

	inputSignal := strings.Split(str, "")
	signal := make([]int, len(inputSignal))
	for i := 0; i < len(inputSignal); i++ {
		signal[i], _ = strconv.Atoi(inputSignal[i])
	}

	/*
		basePattern := []int{0, 1, 0, -1}
		workingSignal := make([]int, len(signal))
		for a := 0; a < 100; a++ {
			for i := 0; i < len(signal); i++ {
				pattern := getPattern(basePattern, i, len(signal))
				workingSignal[i] = getNewValue(pattern, signal)
			}
			signal = workingSignal
		}
		fmt.Println("last phase", signal[:8]) // Part 1: 40580215
	*/

	offset := signal[0]*1000000 + signal[1]*100000 + signal[2]*10000 + signal[3]*1000 + signal[4]*100 + signal[5]*10 + signal[6]

	workingSignal := make([]int, len(signal)*10000)
	for i := 0; i < len(workingSignal); i += len(signal) {
		copy(workingSignal[i:], signal)
	}
	signal = workingSignal

	for i := 0; i < 100; i++ {
		workingSignal[len(signal)-1] = signal[len(signal)-1]
		for j := len(signal) - 2; j >= offset; j-- {
			workingSignal[j] = (signal[j] + workingSignal[j+1]) % 10
			if workingSignal[j] < 0 {
				workingSignal[j] *= -1
			}
		}

		signal = workingSignal
	}

	fmt.Println("Part 2:", signal[offset:offset+8]) // Part 2: 22621597
}
