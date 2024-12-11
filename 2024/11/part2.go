// today is absolutely evil and i love it. i 100% took the bait

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// do a single iteration, as defined in part1
func blinkMap(stones map[int]int) map[int]int {
	nextStones := make(map[int]int)

	for stone, count := range stones {
		str := strconv.Itoa(stone)
		if stone == 0 {
			nextStones[1] += count
		} else if len(str)%2 == 0 {
			leftHalf, err := strconv.Atoi(str[:len(str)/2])
			if err != nil {
				panic(err)
			}
			rightHalf, err := strconv.Atoi(str[len(str)/2:])
			if err != nil {
				panic(err)
			}
			nextStones[leftHalf] += count
			nextStones[rightHalf] += count
		} else {
			nextStones[stone * 2024] += count
		}
	}

	return nextStones
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	tokens := strings.Split(string(input[:len(input)-1]), " ")
	stones := make(map[int]int)
	for i := range tokens {
		stone, err := strconv.Atoi(tokens[i])
		if err != nil {
			panic(err)
		}
		stones[stone] += 1
	}

	const numIterations = 75
	for i := range numIterations {
		fmt.Println("After", i, "iterations:", stones)
		stones = blinkMap(stones)
	}

	result := 0
	for _, count := range stones {
		result += count
	}

	fmt.Println("Final:", stones)
	fmt.Println("Number of stones:", result)

}
