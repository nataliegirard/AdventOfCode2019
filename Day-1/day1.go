package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum1 := 0
	sum2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		mass, _ := strconv.Atoi(line)
		fuel := (mass / 3) - 2
		sum1 += fuel

		fuelPart := fuel
		subSum := 0
		for fuelPart > 0 {
			subSum += fuelPart
			fuelPart = (fuelPart / 3) - 2
		}
		sum2 += subSum
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1", sum1)
	fmt.Println("Part 2", sum2)
}
