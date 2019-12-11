package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type Coordinate struct {
	X, Y int
}

type Tracking struct {
	Count, Color int64
}

func convertInput(input string) []int64 {
	arr := strings.Split(input, ",")

	values := make([]int64, 100000)
	for i := 0; i < len(arr); i++ {
		values[i], _ = strconv.ParseInt(arr[i], 10, 64)
	}
	return values
}

func parseInstruction(code int64) (int64, int64, int64, int64) {
	opcode := code % 100
	mode1 := (code / 100) % 10
	mode2 := (code / 1000) % 10
	mode3 := (code / 10000) % 10

	return opcode, mode1, mode2, mode3
}

func getValue(values []int64, mode int64, index int64, relativeBase int64) int64 {
	if mode == 1 {
		return index
	}
	if mode == 2 {
		return values[index+relativeBase]
	}
	return values[index]
}

func Intcode(values []int64, inbound chan int64, outbound chan int64) int64 {
	var i int64 = 0
	var lastOutput int64 = 0
	var relativeBase int64 = 0
	endloop := false
	for {
		opcode, mode1, mode2, mode3 := parseInstruction(values[i])
		//fmt.Println(i, opcode)
		switch opcode {
		case 1:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			value2 := getValue(values, mode2, values[i+2], relativeBase)
			index3 := values[i+3]
			if mode3 == 2 {
				index3 += relativeBase
			}
			values[index3] = value1 + value2
			i += 4
		case 2:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			value2 := getValue(values, mode2, values[i+2], relativeBase)
			index3 := values[i+3]
			if mode3 == 2 {
				index3 += relativeBase
			}
			values[index3] = value1 * value2
			i += 4
		case 3:
			index := values[i+1]
			inputValue := <-inbound
			if mode1 == 2 {
				index += relativeBase
			}
			values[index] = inputValue
			i += 2
		case 4:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			fmt.Println("Output:", value1)
			lastOutput = value1
			outbound <- value1
			i += 2
		case 5:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			value2 := getValue(values, mode2, values[i+2], relativeBase)
			if value1 != 0 {
				i = value2
			} else {
				i += 3
			}
		case 6:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			value2 := getValue(values, mode2, values[i+2], relativeBase)
			if value1 == 0 {
				i = value2
			} else {
				i += 3
			}
		case 7:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			value2 := getValue(values, mode2, values[i+2], relativeBase)
			index3 := values[i+3]
			if mode3 == 2 {
				index3 += relativeBase
			}

			if value1 < value2 {
				values[index3] = 1
			} else {
				values[index3] = 0
			}
			i += 4
		case 8:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			value2 := getValue(values, mode2, values[i+2], relativeBase)
			index3 := values[i+3]
			if mode3 == 2 {
				index3 += relativeBase
			}

			if value1 == value2 {
				values[index3] = 1
			} else {
				values[index3] = 0
			}
			i += 4
		case 9:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			relativeBase += value1
			i += 2
		case 99:
			close(outbound)
			i = <-inbound
			endloop = true
		}

		if endloop {
			time.Sleep(100 * time.Millisecond)
			break
		}
	}
	return lastOutput
}

func outputWriting(painting [][]int) {
	f, _ := os.Create("output")
	defer f.Close()
	for i := 0; i < len(painting); i++ {
		comma := fmt.Sprint(painting[i])
		fmt.Fprintln(f, comma)
	}
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	str = strings.Replace(str, "\n", "", -1)
	values := convertInput(str)

	toWorker := make(chan int64)
	fromWorker := make(chan int64)
	go Intcode(values, toWorker, fromWorker)

	locx := 0
	locy := 0
	dir := "u"
	panels := make(map[Coordinate]Tracking)

	var origin Coordinate
	origin.X = 0
	origin.Y = 0
	var initial Tracking
	initial.Color = 1
	initial.Count = 0
	panels[origin] = initial

	for {
		var coord Coordinate
		coord.X = locx
		coord.Y = locy

		var track Tracking

		if track, ok := panels[coord]; ok {
			track = panels[coord]

			var color int64 = track.Color
			toWorker <- color
		} else {
			track.Count = 0
			track.Color = 0
			toWorker <- 0
		}

		newColor, ok := <-fromWorker

		if !ok {
			break
		}

		newDirection, ok := <-fromWorker

		if !ok {
			break
		}

		track.Color = newColor
		track.Count++
		panels[coord] = track

		if newDirection == 0 {
			switch dir {
			case "u":
				locx--
				dir = "l"
			case "l":
				locy--
				dir = "d"
			case "d":
				locx++
				dir = "r"
			case "r":
				locy++
				dir = "u"
			}
		} else {
			switch dir {
			case "u":
				locx++
				dir = "r"
			case "l":
				locy++
				dir = "u"
			case "d":
				locx--
				dir = "l"
			case "r":
				locy--
				dir = "d"
			}
		}
	}

	fmt.Println("Number of panels touched:", len(panels))
	// Part 1: 2511

	maxX := 0
	minX := 0
	maxY := 0
	minY := 0
	for crd := range panels {
		if crd.X > maxX {
			maxX = crd.X
		}

		if crd.X < minX {
			minX = crd.X
		}

		if crd.Y > maxY {
			maxY = crd.Y
		}

		if crd.Y < minY {
			minY = crd.Y
		}
	}

	fmt.Println(minX, "-", maxX, "&&", minY, "-", maxY)
	hull := make([][]int, maxY-minY+1)

	for i := 0; i < maxY-minY+1; i++ {
		hull[i] = make([]int, maxX-minX+1)
	}

	for crd, val := range panels {
		fmt.Println(crd.X, crd.Y*-1)

		hull[crd.Y*-1][crd.X] = int(val.Color)
	}
	fmt.Println(hull)
	outputWriting(hull) // Part 2:
}
