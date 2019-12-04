package main

import (
	"fmt"
	"strconv"
	"strings"
)

func onlyIncreases(value int) bool {
	current := 0
	digits := strconv.Itoa(value)
	for i := 0; i < 6; i++ {
		if int(digits[i]) < current {
			return false
		}
		current = int(digits[i])
	}
	return true
}

func hasDoubleDigit(value int) bool {
	digits := strconv.Itoa(value)
	earlier := int(digits[0])
	for i := 1; i < 6; i++ {
		if int(digits[i]) == earlier {
			return true
		}
		earlier = int(digits[i])
	}
	return false
}

func numberOfDoubles(value int) int {
	digits := strings.Split(strconv.FormatInt(int64(value), 10), "")
	earlier := digits[0]
	group := []string{digits[0]}
	var allGroups [][]string = make([][]string, 6)
	groupCount := 0
	for i := 1; i < 6; i++ {
		if digits[i] == earlier {
			group = append(group, digits[i])
		} else {
			allGroups[groupCount] = group
			groupCount++
			group = []string{digits[i]}
		}
		earlier = digits[i]
	}
	allGroups[groupCount] = group
	groupCount++
	doubles := 0

	for i := 0; i < groupCount; i++ {
		if len(allGroups[i]) == 2 {
			doubles++
		}
	}
	return doubles
}

func program() {
	const input string = "307237-769058"
	valueRange := strings.Split(input, "-")
	count := 0
	rangeStart, _ := strconv.Atoi(valueRange[0])
	rangeEnd, _ := strconv.Atoi(valueRange[1])

	for i := rangeStart; i <= rangeEnd; i++ {
		if !onlyIncreases(i) {
			continue
		}
		if !hasDoubleDigit(i) {
			continue
		}
		count++
	}
	fmt.Println("Part1:", count) // Part1: 889

	count = 0
	for i := rangeStart; i <= rangeEnd; i++ {
		if !onlyIncreases(i) {
			continue
		}
		if numberOfDoubles(i) == 0 {
			continue
		}
		count++
	}
	fmt.Println("Part2:", count) // Part2: 589
}

func main() {
	/*test1, _ := strconv.Atoi("111111")
	fmt.Println(test1)
	doubles1 := numberOfDoubles(test1)
	qualifies1 := onlyIncreases(test1) && doubles1 > 0
	fmt.Println("qualifies?", qualifies1, "expect false", doubles1)

	test2, _ := strconv.Atoi("223450")
	fmt.Println(test2)
	doubles2 := numberOfDoubles(test2)
	qualifies2 := onlyIncreases(test2) && doubles2 > 0
	fmt.Println("qualifies?", qualifies2, "expect false", doubles2)

	test3, _ := strconv.Atoi("123789")
	fmt.Println(test3)
	doubles3 := numberOfDoubles(test3)
	qualifies3 := onlyIncreases(test3) && doubles3 > 0
	fmt.Println("qualifies?", qualifies3, "expect false", doubles3)

	test4, _ := strconv.Atoi("112233")
	fmt.Println(test4)
	doubles4 := numberOfDoubles(test4)
	qualifies4 := onlyIncreases(test4) && doubles4 > 0
	fmt.Println("qualifies?", qualifies4, "expect true", doubles4)

	test5, _ := strconv.Atoi("123444")
	fmt.Println(test5)
	doubles5 := numberOfDoubles(test5)
	qualifies5 := onlyIncreases(test5) && doubles5 > 0
	fmt.Println("qualifies?", qualifies5, "expect false", doubles5)

	test6, _ := strconv.Atoi("111122")
	fmt.Println(test6)
	doubles6 := numberOfDoubles(test6)
	qualifies6 := onlyIncreases(test6) && doubles6 > 0
	fmt.Println("qualifies?", qualifies6, "expect true", doubles6)

	test7, _ := strconv.Atoi("345556")
	fmt.Println(test7)
	doubles7 := numberOfDoubles(test7)
	qualifies7 := onlyIncreases(test7) && doubles7 > 0
	fmt.Println("qualifies?", qualifies7, "expect false", doubles7)*/

	program()
}
