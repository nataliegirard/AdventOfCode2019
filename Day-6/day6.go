package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseLine(line string) (string, string) {
	parts := strings.Split(line, ")")
	return parts[0], parts[1]
}

func countOrbits(m map[string]string, key string, count int) int {
	newKey := m[key]

	if m[newKey] != "" {
		count = countOrbits(m, newKey, count+1)
	} else {
		count++
	}

	return count
}

func getOrbitChain(m map[string]string, key string, path []string) []string {
	path = append(path, m[key])
	newKey := m[key]
	newPath := []string{}
	if m[newKey] != "" {
		newPath = getOrbitChain(m, newKey, path)
	} else {
		newPath = append(path, newKey)
	}
	return newPath
}

func mergePaths(path1 []string, path2 []string) int {
	var uniquePath map[string]int
	uniquePath = make(map[string]int)

	for _, node := range path1 {
		uniquePath[node]++
	}
	for _, node := range path2 {
		uniquePath[node]++
	}

	minPath := []string{}
	for key := range uniquePath {
		if uniquePath[key] == 1 {
			minPath = append(minPath, key)
		}
	}

	return len(minPath)
}

func main() {
	var m map[string]string
	m = make(map[string]string)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		cmo, obj := parseLine(line)
		m[obj] = cmo
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1: Get orbit count
	orbits := 0
	for key := range m {
		orbits += countOrbits(m, key, 0)
	}
	fmt.Println("orbits:", orbits) // Part 1: 254447

	// Part 2: Minimum path from YOU to SAN
	emptyPath1 := []string{}
	chain1 := getOrbitChain(m, "YOU", emptyPath1)
	emptyPath2 := []string{}
	chain2 := getOrbitChain(m, "SAN", emptyPath2)

	jumps := mergePaths(chain1, chain2)
	fmt.Println("Jumps needed", jumps) // Part 2: 445
}
