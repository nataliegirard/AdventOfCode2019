package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type packet struct {
	x, y int64
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

func Intcode(values []int64, inbound chan int64, outbound chan int64, signal chan bool) int64 {
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
			acknowledged := false
			for !acknowledged {
				select {
				case signal <- true:
					acknowledged = true
				}
			}
			inputValue := <-inbound
			if mode1 == 2 {
				index += relativeBase
			}
			values[index] = inputValue
			i += 2
		case 4:
			value1 := getValue(values, mode1, values[i+1], relativeBase)
			//fmt.Println("Output:", value1)
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
			close(inbound)
			close(outbound)
			endloop = true
		}

		if endloop {
			time.Sleep(100 * time.Millisecond)
			break
		}
	}
	return lastOutput
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	str = strings.Replace(str, "\n", "", -1)

	numBots := 50
	toWorkers := make([]chan int64, numBots)
	signals := make([]chan bool, numBots)
	fromWorkers := make([]chan int64, numBots)
	packetsForWorker := make([][]packet, numBots)
	for i := 0; i < numBots; i++ {
		values := convertInput(str)
		toWorkers[i] = make(chan int64, 1)
		signals[i] = make(chan bool, 1)
		fromWorkers[i] = make(chan int64, 3)
		go Intcode(values, toWorkers[i], fromWorkers[i], signals[i])
	}
	time.Sleep(time.Second)
	for i := 0; i < numBots; i++ {
		<-signals[i]
		toWorkers[i] <- int64(i)
	}

	var final int64
	var nat packet
	nat.y = int64(-1)
	var lastSent int64 = 0
	temp := false
	count := 0
	for {
		end := false
		idle := 0
		count++
		//fmt.Println("Loop", count)
		for i := 0; i < numBots; i++ {
			select {
			case dest := <-fromWorkers[i]:
				//fmt.Println("received from", i, "for", dest)
				x := <-fromWorkers[i]
				y := <-fromWorkers[i]
				if dest == 255 {
					final = y
					nat.x = x
					nat.y = y
				} else {
					var p packet
					p.x = x
					p.y = y
					packetsForWorker[dest] = append(packetsForWorker[dest], p)
				}
			case sig := <-signals[i]:
				temp = sig
				if len(packetsForWorker[i]) == 0 {
					idle++
					//fmt.Println("sent -1 to", i)
					toWorkers[i] <- int64(-1)
				} else {
					p := packetsForWorker[i][0]
					packetsForWorker[i] = packetsForWorker[i][1:]
					//fmt.Println("sent a packet to", i, "length of queue:", len(packetsForWorker[i]))
					toWorkers[i] <- p.x
					sig = <-signals[i]
					toWorkers[i] <- p.y
				}
			default:
				//fmt.Println("nothing for", i)
			}
		}
		//fmt.Println("Idle?", idle, idle == numBots)
		if idle == numBots {
			if nat.y == -1 {
				continue
			}
			packetsForWorker[0] = append(packetsForWorker[0], nat)

			if nat.y == lastSent {
				fmt.Println("Seen this packet before", lastSent)
				end = true
			} else {
				lastSent = nat.y
				fmt.Println("nat", nat.y)
			}
		}
		if end {
			break
		}
	}
	fmt.Println(temp, "final:", final) // Part 1: 23213, part 2: 17874
}
