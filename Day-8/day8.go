package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func numberInLayer(layer string, number int) int {
	count := 0
	pixels := strings.Split(layer, "")
	for i := 0; i < len(pixels); i++ {
		pixel, _ := strconv.Atoi(pixels[i])

		if pixel == number {
			count++
		}
	}
	return count
}

func renderImage(final string, layer string) string {
	var image string
	pixelsFinal := strings.Split(final, "")
	pixelsLayer := strings.Split(layer, "")
	for i := 0; i < len(pixelsFinal); i++ {
		pixelF, _ := strconv.Atoi(pixelsFinal[i])
		pixelL, _ := strconv.Atoi(pixelsLayer[i])

		if pixelF == 2 {
			new := strconv.Itoa(pixelL)
			image = image + new
		} else {
			new := strconv.Itoa(pixelF)
			image = image + new
		}
	}
	return image
}

func main() {
	width := 25
	height := 6
	b, err := ioutil.ReadFile("input.txt")
	/*width := 2
	height := 2
	b, err := ioutil.ReadFile("test2.txt")*/
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	str = strings.TrimSuffix(str, "\n")

	layerLength := width * height
	layers := make([]string, len(str)/layerLength)

	for i := 0; i < len(layers); i++ {
		start := i * layerLength
		end := (i + 1) * layerLength
		layers[i] = str[start:end]
	}

	minZeros := layerLength + 1
	var minLayer int
	for index, layer := range layers {
		count := numberInLayer(layer, 0)
		if count < minZeros {
			minZeros = count
			minLayer = index
		}
	}
	fmt.Println("Layer with min zeros:", minLayer, "with:", minZeros)

	numOnes := numberInLayer(layers[minLayer], 1)
	numTwos := numberInLayer(layers[minLayer], 2)
	fmt.Println("Part 1:", numOnes*numTwos) // 2286

	final := layers[0]
	for i := 1; i < len(layers); i++ {
		final = renderImage(final, layers[i])
	}

	// Output layers
	f, _ := os.Create("output")
	defer f.Close()
	for i := 0; i < height; i++ {
		start := i * width
		end := (i + 1) * width
		row := final[start:end]
		broken := strings.Split(row, "")
		comma := strings.Join(broken, ",")
		fmt.Fprintln(f, comma)
	}
	// Part 2:
}
