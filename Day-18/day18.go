package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type coordinate struct {
	x, y, steps int
	name        string
}

func findObjects(vault [][]string) []coordinate {
	var objects []coordinate
	var c coordinate
	for y := 0; y < len(vault); y++ {
		for x := 0; x < len(vault[y]); x++ {
			c.x = x
			c.y = y
			c.name = vault[y][x]
			c.steps = -1
			if vault[y][x] == "#" {
				continue
			} else if vault[y][x] != "." {
				objects = append(objects, c)
			}
		}
	}
	return objects
}

func floodMaze(maze [][]string, start coordinate) {
	currSteps, _ := strconv.Atoi(maze[start.y][start.x])
	c := start
	t := maze[start.y][start.x]
	nextStep := strconv.Itoa(currSteps + 1)

	// north
	c.y--
	t = maze[c.y][c.x]
	if t != "#" {
		a, _ := strconv.Atoi(t)
		if a == -1 || a > currSteps+1 {
			t = nextStep
			maze[c.y][c.x] = t
			floodMaze(maze, c)
		}
	}
	c.y++

	// south
	c.y++
	t = maze[c.y][c.x]
	if t != "#" {
		a, _ := strconv.Atoi(t)
		if a == -1 || a > currSteps+1 {
			t = nextStep
			maze[c.y][c.x] = t
			floodMaze(maze, c)
		}
	}
	c.y--

	// west
	c.x--
	t = maze[c.y][c.x]
	if t != "#" {
		a, _ := strconv.Atoi(t)
		if a == -1 || a > currSteps+1 {
			t = nextStep
			maze[c.y][c.x] = t
			floodMaze(maze, c)
		}
	}
	c.x++

	// east
	c.x++
	t = maze[c.y][c.x]
	if t != "#" {
		a, _ := strconv.Atoi(t)
		if a == -1 || a > currSteps+1 {
			t = nextStep
			maze[c.y][c.x] = t
			floodMaze(maze, c)
		}
	}
	c.x--
}

func printArea(area [][]string, toscreen bool) {
	f, _ := os.Create("output.csv")
	defer f.Close()
	for i := 0; i < len(area); i++ {
		comma := strings.Join(area[i], ",")
		fmt.Fprintln(f, comma)
		if toscreen {
			fmt.Println(strings.Join(area[i], ""))
		}
	}
}

func main() {
	filename := "input.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var vault [][]string
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		vault = append(vault, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	objects := findObjects(vault)

	var start coordinate
	for i := 0; i < len(objects); i++ {
		if objects[i].name == "@" {
			start = objects[i]
		}
	}
	printArea(vault, false)

	steps := vault
	for y := 0; y < len(steps); y++ {
		for x := 0; x < len(steps[y]); x++ {
			if steps[y][x] != "#" {
				steps[y][x] = strconv.Itoa(-1)
			}
		}
	}
	steps[start.y][start.x] = "0"

	floodMaze(steps, start)
	for i := 0; i < len(objects); i++ {
		obj := objects[i]
		objects[i].steps, _ = strconv.Atoi(steps[obj.y][obj.x])
	}

	var keys []coordinate
	for i := 0; i < len(objects); i++ {
		character := rune(objects[i].name[0])
		if unicode.IsLower(character) {
			keys = append(keys, objects[i])
			steps[objects[i].y][objects[i].x] = objects[i].name
		}
	}
	fmt.Println("Keys:", keys)
	printArea(steps, false)
}
