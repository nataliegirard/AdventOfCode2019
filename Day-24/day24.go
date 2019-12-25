package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func countBugsAdj(area []string, index int) int {
	count := 0
	//north
	if index-5 >= 0 && area[index-5] == "#" {
		count++
	}

	//south
	if index+5 < 25 && area[index+5] == "#" {
		count++
	}

	//west
	if index%5 > 0 && area[index-1] == "#" {
		count++
	}

	//east
	if index%5 < 4 && area[index+1] == "#" {
		count++
	}

	return count
}

func bugLives(area []string, index int) bool {
	count := countBugsAdj(area, index)
	if count == 1 {
		return true
	}
	return false
}

func getsInfested(area []string, index int) bool {
	count := countBugsAdj(area, index)
	if count == 1 || count == 2 {
		return true
	}
	return false
}

func iterateArea(area []string) []string {
	newArea := make([]string, len(area))

	for i := 0; i < len(area); i++ {
		if area[i] == "#" {
			if bugLives(area, i) {
				newArea[i] = "#"
			} else {
				newArea[i] = "."
			}
		} else if area[i] == "." {
			if getsInfested(area, i) {
				newArea[i] = "#"
			} else {
				newArea[i] = "."
			}
		}
	}

	return newArea
}

func countBugs3d(layers map[int]([]string), layerIndex int, pos int) int {
	count := 0
	below := layers[layerIndex-1]
	above := layers[layerIndex+1]

	if len(below) == 0 {
		below = make([]string, 25)
		for i := 0; i < 25; i++ {
			below[i] = "."
		}
		below[12] = "?"
	}
	if len(above) == 0 {
		above = make([]string, 25)
		for i := 0; i < 25; i++ {
			above[i] = "."
		}
		above[12] = "?"
	}

	// north
	if pos < 5 && above[7] == "#" {
		// top edge
		count++
	} else if pos == 17 {
		// below ? -> count bottom edge of layer below
		l := below
		if l[20] == "#" {
			count++
		}
		if l[21] == "#" {
			count++
		}
		if l[22] == "#" {
			count++
		}
		if l[23] == "#" {
			count++
		}
		if l[24] == "#" {
			count++
		}
	} else if pos-5 >= 0 && layers[layerIndex][pos-5] == "#" {
		// normal
		count++
	}

	// south
	if pos >= 20 && above[17] == "#" {
		// bottom edge
		count++
	} else if pos == 7 {
		// above ? -> count top edge of layer below
		l := below
		if l[0] == "#" {
			count++
		}
		if l[1] == "#" {
			count++
		}
		if l[2] == "#" {
			count++
		}
		if l[3] == "#" {
			count++
		}
		if l[4] == "#" {
			count++
		}
	} else if pos+5 < 25 && layers[layerIndex][pos+5] == "#" {
		// normal
		count++
	}

	// west
	if pos%5 == 0 && above[11] == "#" {
		// left edge
		count++
	} else if pos == 13 {
		// right of ? -> count right edge of layer below
		l := below
		if l[4] == "#" {
			count++
		}
		if l[9] == "#" {
			count++
		}
		if l[14] == "#" {
			count++
		}
		if l[19] == "#" {
			count++
		}
		if l[24] == "#" {
			count++
		}
	} else if pos%5 > 0 && layers[layerIndex][pos-1] == "#" {
		// normal
		count++
	}

	// east
	if pos%5 == 4 && above[13] == "#" {
		// right edge
		count++
	} else if pos == 11 {
		// left of ? -> count left edge of layer below
		l := below
		if l[0] == "#" {
			count++
		}
		if l[5] == "#" {
			count++
		}
		if l[10] == "#" {
			count++
		}
		if l[15] == "#" {
			count++
		}
		if l[20] == "#" {
			count++
		}
	} else if pos%5 < 4 && layers[layerIndex][pos+1] == "#" {
		// normal
		count++
	}

	return count
}

func bugLives3d(layers map[int]([]string), layerIndex int, pos int) bool {
	count := countBugs3d(layers, layerIndex, pos)
	if count == 1 {
		return true
	}
	return false
}

func getsInfested3d(layers map[int]([]string), layerIndex int, pos int) bool {
	count := countBugs3d(layers, layerIndex, pos)
	if count == 1 || count == 2 {
		return true
	}
	return false
}

func iterateLayer(layers map[int]([]string), layerIndex int) []string {
	newArea := make([]string, 25)
	area := layers[layerIndex]

	if len(area) == 0 {
		area = make([]string, 25)
		for i := 0; i < 25; i++ {
			area[i] = "."
		}
		area[12] = "?"
		layers[layerIndex] = area
	}

	for i := 0; i < len(area); i++ {
		if area[i] == "#" {
			if bugLives3d(layers, layerIndex, i) {
				newArea[i] = "#"
			} else {
				newArea[i] = "."
			}
		} else if area[i] == "." {
			if getsInfested3d(layers, layerIndex, i) {
				newArea[i] = "#"
			} else {
				newArea[i] = "."
			}
		} else {
			newArea[i] = area[i]
		}
	}
	return newArea
}

func printMap(area []string) {
	for i := 0; i < len(area); i++ {
		fmt.Printf("%s", area[i])
		if i%5 == 4 {
			fmt.Printf("\n")
		}
	}
}

func main() {
	/*
		filename := "input.txt"

		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		var area []string
		for scanner.Scan() {
			line := scanner.Text()
			row := strings.Split(line, "")
			area = append(area, row...)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		solutions := make(map[string]int)
		var result []string
		for i := 0; ; i++ {
			//fmt.Println("round:", i)
			//printMap(area)
			area = iterateArea(area)
			str := strings.Join(area, "")
			if solutions[str] == 1 {
				result = area
				break
			} else {
				solutions[str] = 1
			}
		}

		fmt.Println("Found solution:")
		printMap(result)
		bioDiversity := 0
		for i := 0; i < len(result); i++ {
			if result[i] == "#" {
				bioDiversity += int(math.Pow(2, float64(i)))
			}
		}
		fmt.Println("Part 1 :", bioDiversity) // Part 1: 32506911
	*/

	// Part 2
	filename := "input2.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var area []string
	var layers = make(map[int]([]string))
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		area = append(area, row...)
	}
	layers[0] = area

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 200; i++ {
		fmt.Println("round:", i+1)
		newLayers := make(map[int]([]string))
		min := 0
		max := 0
		for j := range layers {
			if j < min {
				min = j
			}
			if j > max {
				max = j
			}
			newLayers[j] = iterateLayer(layers, j)
		}

		t := iterateLayer(layers, min-1)
		hasBug := false
		for i := 0; i < 25; i++ {
			if t[i] == "#" {
				hasBug = true
			}
		}
		if hasBug {
			newLayers[min-1] = t
		}

		t = iterateLayer(layers, max+1)
		hasBug = false
		for i := 0; i < 25; i++ {
			if t[i] == "#" {
				hasBug = true
			}
		}
		if hasBug {
			newLayers[max+1] = t
		}

		layers = newLayers
	}
	for i := -5; i < 6; i++ {
		fmt.Println("Layer:", i)
		printMap(layers[i])
	}

	countBugs := 0
	for _, l := range layers {
		for i := 0; i < 25; i++ {
			if l[i] == "#" {
				countBugs++
			}
		}
	}
	fmt.Println("Part 2:", countBugs) // Part 2: 2025
}
