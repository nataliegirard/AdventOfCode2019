package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func convertArray(input string) []int {
	arr := strings.Split(input, ",")

	var values = []int{}

	for _, i := range arr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		values = append(values, j)
	}
	return values
}

// for Part 2
func Explore(values []int) int {
	result2 := 0
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			result2 = Intcode(values, noun, verb)
			if result2 == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return -1
}

func Intcode(input []int, noun int, verb int) int {
	values := make([]int, len(input))
	copy(values, input)
	values[1] = noun
	values[2] = verb

	i := 0
	for {
		switch values[i] {
		case 1:
			index1 := values[i+1]
			index2 := values[i+2]
			index3 := values[i+3]
			values[index3] = values[index1] + values[index2]
		case 2:
			index1 := values[i+1]
			index2 := values[i+2]
			index3 := values[i+3]
			values[index3] = values[index1] * values[index2]
		case 99:
			return values[0]
		default:
			return -1
		}
		i += 4
	}
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	str = strings.TrimSuffix(str, "\n")
	vals := convertArray(str)

	result := Intcode(vals, 12, 2)
	fmt.Println("Part 1 answer:", result) // 3765464

	ex1 := "1,9,10,3,2,3,11,0,99,30,40,50"
	vals1 := convertArray(ex1)
	res1 := Intcode(vals1, vals1[1], vals1[2])
	fmt.Println("Ex1: expect", 3500, "received", res1)

	ex2 := "1,0,0,0,99"
	vals2 := convertArray(ex2)
	res2 := Intcode(vals2, vals2[1], vals2[2])
	fmt.Println("Ex2: expect", 2, "received", res2)

	ex3 := "2,3,0,3,99"
	vals3 := convertArray(ex3)
	res3 := Intcode(vals3, vals3[1], vals3[2])
	fmt.Println("Ex3: expect", 2, "received", res3)

	ex4 := "2,4,4,5,99,0"
	vals4 := convertArray(ex4)
	res4 := Intcode(vals4, vals4[1], vals4[2])
	fmt.Println("Ex4: expect", 2, "received", res4)

	ex5 := "1,1,1,4,99,5,6,0,99"
	vals5 := convertArray(ex5)
	res5 := Intcode(vals5, vals5[1], vals5[2])
	fmt.Println("Ex5: expect", 30, "received", res5)

	// Part 2
	answer := Explore(vals)
	fmt.Println("Part 2 answer:", answer) // 7610
}
