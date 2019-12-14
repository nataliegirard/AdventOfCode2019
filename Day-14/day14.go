package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type chemical struct {
	amount int
	name   string
}

type equation struct {
	inputs []chemical
	output chemical
}

func parseLine(line string) equation {
	var equ equation
	parts := strings.Split(line, " => ")
	ins := strings.Split(parts[0], ", ")
	outs := parts[1]

	inputs := make([]chemical, len(ins))
	for i := range ins {
		t := strings.Split(ins[i], " ")
		inputs[i].amount, _ = strconv.Atoi(t[0])
		inputs[i].name = t[1]
	}

	equ.inputs = inputs
	temp := strings.Split(outs, " ")
	equ.output.amount, _ = strconv.Atoi(temp[0])
	equ.output.name = temp[1]

	return equ
}

func getStock(recipes map[string]equation, product chemical, stock map[string]int, required map[string]int) (map[string]int, map[string]int) {
	//fmt.Println("Making", product)

	for stock[product.name] < product.amount {
		//fmt.Println("Missing", product.amount-stock[product.name], product.name)
		eq := recipes[product.name]
		//fmt.Println("Using recipe:", eq)

		// execute recipe
		for i := range eq.inputs {
			if eq.inputs[i].name == "ORE" {
				required["ORE"] += eq.inputs[i].amount
				continue
			}

			if stock[eq.inputs[i].name] < eq.inputs[i].amount {
				// if we don't have input in stock, get it
				stock, required = getStock(recipes, eq.inputs[i], stock, required)
			}

			// remove inputs from stock to product output
			stock[eq.inputs[i].name] -= eq.inputs[i].amount
			required[eq.inputs[i].name] += eq.inputs[i].amount
		}

		// increase stock of output
		stock[eq.output.name] += eq.output.amount
	}

	//fmt.Println("Created all products", stock, required)

	return stock, required
}

func exactOre(recipes map[string]equation, eq equation) float64 {
	if len(eq.inputs) == 1 && eq.inputs[0].name == "ORE" {
		return float64(eq.inputs[0].amount) / float64(eq.output.amount)
	}

	var total float64 = 0
	for i := range eq.inputs {
		needs := exactOre(recipes, recipes[eq.inputs[i].name])
		total += needs * float64(eq.inputs[i].amount)
	}

	return total / float64(eq.output.amount)
}

func main() {
	filename := "input.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	recipes := make(map[string]equation)
	for scanner.Scan() {
		line := scanner.Text()
		eq := parseLine(line)

		if _, ok := recipes[eq.output.name]; ok {
			fmt.Println("ERROR")
		}

		recipes[eq.output.name] = eq
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var fuel chemical
	fuel.name = "FUEL"
	fuel.amount = 1

	stock := make(map[string]int)
	required := make(map[string]int)
	stock, required = getStock(recipes, fuel, stock, required)
	//fmt.Println(required)
	//fmt.Println(stock)

	fmt.Println("1 FUEL requires", required["ORE"], "ORE") // Part 1: 654909

	exact := exactOre(recipes, recipes["FUEL"])
	fmt.Println(math.Floor(float64(1000000000000) / exact)) // Part 2: 2876992
}
