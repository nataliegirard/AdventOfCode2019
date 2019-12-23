package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type instruction struct {
	operation string
	num       int
}

func dealToNewStack(deck []int) []int {
	newDeck := make([]int, len(deck))
	i := len(deck)
	for j := 0; j < len(deck); j++ {
		i--
		newDeck[i] = deck[j]
	}

	return newDeck
}

func cutCards(deck []int, n int) []int {
	var newDeck []int
	j := n
	for j < 0 {
		j += len(deck)
	}

	newDeck = deck[j:]
	newDeck = append(newDeck, deck[:j]...)

	return newDeck
}

func dealWithIncrement(deck []int, n int) []int {
	newDeck := make([]int, len(deck))
	j := 0
	for i := 0; i < len(deck); i++ {
		newDeck[j] = deck[i]
		j = (j + n) % len(deck)
	}
	return newDeck
}

func executeCommand(command instruction, deck []int) []int {
	switch command.operation {
	case "rev":
		deck = dealToNewStack(deck)
	case "cut":
		deck = cutCards(deck, command.num)
	case "inc":
		deck = dealWithIncrement(deck, command.num)
	}
	return deck
}

func parseCommand(command string) instruction {
	var parsed instruction

	deal, _ := regexp.Compile("deal into new stack")
	cut, _ := regexp.Compile("cut (-?[0-9]+)")
	inc, _ := regexp.Compile("deal with increment ([0-9]+)")

	switch {
	case deal.MatchString(command):
		parsed.operation = "rev"
	case cut.MatchString(command):
		match := cut.FindStringSubmatch(command)
		num, _ := strconv.Atoi(match[1])
		parsed.operation = "cut"
		parsed.num = num
	case inc.MatchString(command):
		match := inc.FindStringSubmatch(command)
		num, _ := strconv.Atoi(match[1])
		parsed.operation = "inc"
		parsed.num = num
	}

	return parsed
}

func simplifyCommands(commands []instruction, deckSize int) ([]instruction, bool) {
	changed := false

	i := 1

	for {
		// check command-1 and command
		op1 := commands[i-1].operation
		op2 := commands[i].operation
		arg1 := commands[i-1].num
		arg2 := commands[i].num

		var secondHalf []instruction
		if i+1 < len(commands) {
			t := commands[i+1:]
			for j := 0; j < len(t); j++ {
				secondHalf = append(secondHalf, t[j])
			}
		}

		if op1 == op2 {
			if op1 == "inc" {
				newComm := commands[i]
				newComm.num = (arg1 * arg2) % deckSize
				commands = append(commands[:i-1], newComm)
				commands = append(commands, secondHalf...)
				changed = true
			}

			if op1 == "cut" {
				newComm := commands[i]
				newComm.num = (arg1 + arg2) % deckSize
				commands = append(commands[:i-1], newComm)
				commands = append(commands, secondHalf...)
				changed = true
			}

			if op1 == "rev" {
				commands = append(commands[:i-1], secondHalf...)
				changed = true
			}
		} else {
			if op1 == "rev" && op2 == "cut" {
				first := commands[i-1]
				second := commands[i]
				second.num = arg2 * -1
				commands = append(commands[:i-1], second, first)
				commands = append(commands, secondHalf...)
				changed = true
				i++
			} else if op1 == "rev" && op2 == "inc" {
				first := commands[i-1]
				second := commands[i]
				var newComm instruction
				newComm.operation = "cut"
				newComm.num = 1 - arg2

				commands = append(commands[:i-1], second, newComm, first)
				commands = append(commands, secondHalf...)
				changed = true
				i += 2
			} else if op1 == "cut" && op2 == "inc" {
				first := commands[i-1]
				second := commands[i]
				first.num = (arg1 * arg2) % deckSize
				commands = append(commands[:i-1], second, first)
				commands = append(commands, secondHalf...)
				changed = true
				i++
			} else {
				i++
			}
		}

		if i == len(commands) {
			break
		}
	}

	return commands, changed
}

func main() {
	deckSize := 10007
	deck := make([]int, deckSize)
	for i := 0; i < deckSize; i++ {
		deck[i] = i
	}

	filename := "input.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var commands []instruction
	for scanner.Scan() {
		line := scanner.Text()
		t := parseCommand(line)
		commands = append(commands, t)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	changed := true
	for changed {
		commands, changed = simplifyCommands(commands, deckSize)
	}

	for _, c := range commands {
		deck = executeCommand(c, deck)
	}

	fmt.Println("Reduced\n", commands)

	result := 0
	for i := 0; i < len(deck); i++ {
		if deck[i] == 2019 {
			result = i
		}
	}
	fmt.Println("Part 1:", result) // Part 1: 8379

}
