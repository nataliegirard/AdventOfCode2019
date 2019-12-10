package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type Asteroid struct {
	X, Y, Count int
}

func findClosest(list []Asteroid, origin Asteroid) int {
	var closest float64 = 100000
	var closestIndex int
	for i := 0; i < len(list); i++ {
		diffx := origin.X - list[i].X
		diffy := origin.Y - list[i].Y
		diffX := float64(diffx * diffx)
		diffY := float64(diffy * diffy)
		distance := math.Sqrt(diffX + diffY)

		if distance < closest {
			closest = distance
			closestIndex = i
		}
	}
	return closestIndex
}

func main() {
	filename := "input.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	asteroids := make([]Asteroid, 0)
	y := 0
	for scanner.Scan() {
		l := scanner.Text()
		line := strings.Split(l, "")

		for x := 0; x < len(line); x++ {
			if line[x] == "#" {
				newAsteroid := Asteroid{x, y, 0}
				asteroids = append(asteroids, newAsteroid)
			}
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(asteroids); i++ {
		// for each asteroid, go through list of asteroids and count how many can be seen
		seen := make(map[float64]int)

		for j := 0; j < len(asteroids); j++ {
			if i == j {
				continue
			}
			diffx := asteroids[i].X - asteroids[j].X
			diffy := asteroids[i].Y - asteroids[j].Y
			angle := math.Atan2(float64(diffy), float64(diffx))
			seen[angle] = 1
		}
		asteroids[i].Count = len(seen)
	}
	fmt.Println("asteroids", asteroids)

	maxSeen := 0
	var bestAsteroid Asteroid
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i].Count > maxSeen {
			maxSeen = asteroids[i].Count
			bestAsteroid = asteroids[i]
		}
	}
	fmt.Println("Best:", bestAsteroid) // Part 1: (29, 28) count: 256
	fmt.Println("")

	// Part 2
	// get all angles from the best asteroid
	angles := make(map[float64]([]Asteroid))
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] == bestAsteroid {
			fmt.Println("Found best")
			continue
		}
		diffx := bestAsteroid.X - asteroids[i].X
		diffy := bestAsteroid.Y - asteroids[i].Y
		a := math.Atan2(float64(diffy), float64(diffx))
		a = a * 180 / math.Pi
		if a < 0 {
			a = a + 360
		}
		angles[a] = append(angles[a], asteroids[i])
	}
	fmt.Println("Angles", angles)

	// go through each angle from top clockwise, removing and counting the asteroids it hits
	c := 0
	angleList := make([]float64, len(angles))
	for key := range angles {
		angleList[c] = key
		c++
	}
	fmt.Println("")
	fmt.Println(angleList)
	fmt.Println(len(angleList))

	start := 0
	sort.Float64s(angleList)
	for i := 0; i < len(angleList); i++ {
		if angleList[i] == float64(90) {
			start = i
		}
	}
	fmt.Println("start", start)

	count := 0
	for i := start; count < 200; i++ {
		if i == len(angleList) {
			i = 0
		}
		possibles := angles[angleList[i]]
		//fmt.Println(i, angleList[i])
		length := len(possibles)
		if length == 0 {
			continue
		}

		closest := findClosest(possibles, bestAsteroid)
		temp := possibles[closest]
		angles[angleList[i]] = append(possibles[:closest], possibles[closest+1:]...)
		count++

		if count == 1 {
			fmt.Println("1:", temp)
			fmt.Println(angleList[i])
		}
		if count == 2 {
			fmt.Println("2:", temp)
			fmt.Println(angleList[i])
		}
		if count == 3 {
			fmt.Println("3:", temp)
			fmt.Println(angleList[i])
		}
		if count == 10 {
			fmt.Println("10:", temp)
		}
		if count == 20 {
			fmt.Println("20:", temp)
		}
		if count == 50 {
			fmt.Println("50:", temp)
		}
		if count == 100 {
			fmt.Println("100:", temp)
		}
		if count == 199 {
			fmt.Println("199:", temp)
		}

		if count == 200 {
			fmt.Println("Found 200:", temp)
		}

		if count == 201 {
			fmt.Println("Found 201:", temp)
		}
	}
	// part 2: {17 7 240} X coordinate by 100 and then add its Y
	// 170
}
