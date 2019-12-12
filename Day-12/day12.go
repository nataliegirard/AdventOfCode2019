package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord3d struct {
	x, y, z int
}

type moon struct {
	position coord3d
	velocity coord3d
}

func calculateVelocity(moons []moon, index int, dimension string) int {
	currMoon := moons[index]
	vel := 0
	currMoonPos := 0

	switch dimension {
	case "x":
		currMoonPos = currMoon.position.x
		vel = currMoon.velocity.x
	case "y":
		currMoonPos = currMoon.position.y
		vel = currMoon.velocity.y
	case "z":
		currMoonPos = currMoon.position.z
		vel = currMoon.velocity.z
	}

	for i := 0; i < len(moons); i++ {
		if i == index {
			continue
		}

		pos := 0

		switch dimension {
		case "x":
			pos = moons[i].position.x
		case "y":
			pos = moons[i].position.y
		case "z":
			pos = moons[i].position.z
		}

		if currMoonPos < pos {
			vel++
		}
		if currMoonPos > pos {
			vel--
		}
	}

	return vel
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func calculateEnergy(moons []moon) int {
	totalEnergy := 0
	for i := 0; i < len(moons); i++ {
		pos := moons[i].position
		vel := moons[i].velocity

		pot := abs(pos.x) + abs(pos.y) + abs(pos.z)
		kin := abs(vel.x) + abs(vel.y) + abs(vel.z)
		total := pot * kin

		totalEnergy += total
	}
	return totalEnergy
}

func printMoons(moons []moon) {
	for i := 0; i < len(moons); i++ {
		pos := moons[i].position
		vel := moons[i].velocity
		fmt.Printf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>\n", pos.x, pos.y, pos.z, vel.x, vel.y, vel.z)
	}
}

func part1() int {
	filename := "input.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	moons := make([]moon, 4)
	moonCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "<", "", -1)
		line = strings.Replace(line, ">", "", -1)
		coords := strings.Split(line, ", ")

		moons[moonCount].position.x, _ = strconv.Atoi(strings.Replace(coords[0], "x=", "", -1))
		moons[moonCount].position.y, _ = strconv.Atoi(strings.Replace(coords[1], "y=", "", -1))
		moons[moonCount].position.z, _ = strconv.Atoi(strings.Replace(coords[2], "z=", "", -1))

		moonCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	steps := 0
	for steps = 0; steps < 1000; steps++ {
		for i := 0; i < len(moons); i++ {
			moons[i].velocity.x = calculateVelocity(moons, i, "x")
			moons[i].velocity.y = calculateVelocity(moons, i, "y")
			moons[i].velocity.z = calculateVelocity(moons, i, "z")
		}

		for i := 0; i < len(moons); i++ {
			moons[i].position.x += moons[i].velocity.x
			moons[i].position.y += moons[i].velocity.y
			moons[i].position.z += moons[i].velocity.z
		}
	}

	//printMoons(moons)

	totalEnergy := calculateEnergy(moons)
	return totalEnergy
}

func findPeriod(moons []moon, dimension string) int {
	period := 0
	for i := 0; i != -1; i++ {
		allstopped := true
		vel := 0
		for i := 0; i < len(moons); i++ {
			vel = calculateVelocity(moons, i, dimension)

			if vel != 0 {
				allstopped = false
			}

			switch dimension {
			case "x":
				moons[i].velocity.x = vel
			case "y":
				moons[i].velocity.y = vel
			case "z":
				moons[i].velocity.z = vel
			}
		}

		if allstopped {
			period = i + 1
			break
		}

		for i := 0; i < len(moons); i++ {
			switch dimension {
			case "x":
				moons[i].position.x += moons[i].velocity.x
			case "y":
				moons[i].position.y += moons[i].velocity.y
			case "z":
				moons[i].position.z += moons[i].velocity.z
			}
		}
	}
	return period
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b, c int) int {
	result := a * b / gcd(a, b)
	result = result * c / gcd(result, c)
	return result
}

func part2() int {
	filename := "input.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	moons := make([]moon, 4)
	moonCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "<", "", -1)
		line = strings.Replace(line, ">", "", -1)
		coords := strings.Split(line, ", ")

		moons[moonCount].position.x, _ = strconv.Atoi(strings.Replace(coords[0], "x=", "", -1))
		moons[moonCount].position.y, _ = strconv.Atoi(strings.Replace(coords[1], "y=", "", -1))
		moons[moonCount].position.z, _ = strconv.Atoi(strings.Replace(coords[2], "z=", "", -1))

		moonCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	periodX := findPeriod(moons, "x")
	periodY := findPeriod(moons, "y")
	periodZ := findPeriod(moons, "z")
	period := lcm(periodX, periodY, periodZ)
	//fmt.Println("Periods", periodX, periodY, periodZ, "Equal:", period*2)

	return period * 2
}

func main() {
	result := part1()
	fmt.Println("total Energy:", result) // Part 1: 7636

	result2 := part2()
	fmt.Println("Steps to initial:", result2) // Par 2: 281691380235984
}
