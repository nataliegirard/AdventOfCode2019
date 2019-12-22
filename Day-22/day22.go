package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

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
	newDeck := make([]int, len(deck))
	j := n
	for j < 0 {
		j += len(deck)
	}

	for i := 0; i < len(deck); i++ {
		newDeck[i] = deck[j]
		j++
		if j >= len(deck) {
			j -= len(deck)
		}
	}

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

func executeCommand(deck []int, command string) []int {
	newDeck := make([]int, len(deck))

	deal, _ := regexp.Compile("deal into new stack")
	cut, _ := regexp.Compile("cut (-?[0-9]+)")
	inc, _ := regexp.Compile("deal with increment ([0-9]+)")

	switch {
	case deal.MatchString(command):
		newDeck = dealToNewStack(deck)
		fmt.Println("deal new stack")
	case cut.MatchString(command):
		match := cut.FindStringSubmatch(command)
		num, _ := strconv.Atoi(match[1])
		newDeck = cutCards(deck, num)
		fmt.Println("cut cards", num)
	case inc.MatchString(command):
		match := inc.FindStringSubmatch(command)
		num, _ := strconv.Atoi(match[1])
		newDeck = dealWithIncrement(deck, num)
		fmt.Println("deal inc", num)
	}

	return newDeck
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

	for scanner.Scan() {
		line := scanner.Text()
		deck = executeCommand(deck, line)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := 0
	for i := 0; i < len(deck); i++ {
		if deck[i] == 2019 {
			result = i
		}
	}
	fmt.Println("Part 1:", result) // Part 1:

	//deck := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Function for deal to new stack
	/*newDeck := dealToNewStack(deck)
	fmt.Println(newDeck) // 9 8 7 6 5 4 3 2 1 0 */

	// Function for cut N cards
	/*newDeck := cutCards(deck, 3)
	fmt.Println(newDeck) // 3 4 5 6 7 8 9 0 1 2
	newDeck = cutCards(deck, -4)
	fmt.Println(newDeck) // 6 7 8 9 0 1 2 3 4 5 */

	// Function for deal with increment
	/*newDeck := dealWithIncrement(deck, 3)
	fmt.Println(newDeck) // 0 7 4 1 8 5 2 9 6 3 */

	// Example 1
	/*deck = dealWithIncrement(deck, 7)
	deck = dealToNewStack(deck)
	deck = dealToNewStack(deck)
	fmt.Println(deck) // 0 3 6 9 2 5 8 1 4 7 */

	// Example 2
	/*deck = cutCards(deck, 6)
	deck = dealWithIncrement(deck, 7)
	deck = dealToNewStack(deck)
	fmt.Println(deck) // 3 0 7 4 1 8 5 2 9 6 */

	// Example 3
	/*deck = dealWithIncrement(deck, 7)
	deck = dealWithIncrement(deck, 9)
	deck = cutCards(deck, -2)
	fmt.Println(deck) // 6 3 0 7 4 1 8 5 2 9 */

	// Example 4
	/*deck = dealToNewStack(deck)
	deck = cutCards(deck, -2)
	deck = dealWithIncrement(deck, 7)
	deck = cutCards(deck, 8)
	deck = cutCards(deck, -4)
	deck = dealWithIncrement(deck, 7)
	deck = cutCards(deck, 3)
	deck = dealWithIncrement(deck, 9)
	deck = dealWithIncrement(deck, 3)
	deck = cutCards(deck, -1)
	fmt.Println(deck) // 9 2 5 8 1 4 7 0 3 6 */
}
