package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func printArea(area [][]string) {
	f, _ := os.Create("output.csv")
	defer f.Close()
	for i := 0; i < len(area); i++ {
		comma := strings.Join(area[i], ",")
		fmt.Fprintln(f, comma)
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

	var maze [][]string
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		maze = append(maze, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	printArea(maze)
}
