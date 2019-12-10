package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Asteroid struct {
	X, Y, Count int
}

func main() {
	filename := "test1.txt"

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

		fmt.Println("Checking asteroid:", asteroids[i])
		for j := 0; j < len(asteroids); j++ {
			if i == j {
				continue
			}

		}
	}

	fmt.Println("asteroids", asteroids)
}
