package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// do a single iteration, as defined in part1
func blink(stones []int) []int {
	var nextStones []int

	for _, stone := range stones {
		str := strconv.Itoa(stone)
		if stone == 0 {
			nextStones = append(nextStones, 1)
		} else if len(str)%2 == 0 {
			leftHalf, err := strconv.Atoi(str[:len(str)/2])
			if err != nil {
				panic(err)
			}
			rightHalf, err := strconv.Atoi(str[len(str)/2:])
			if err != nil {
				panic(err)
			}
			nextStones = append(nextStones, leftHalf, rightHalf)
		} else {
			nextStones = append(nextStones, stone * 2024)
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
	var stones []int
	for i := range tokens {
		stone, err := strconv.Atoi(tokens[i])
		if err != nil {
			panic(err)
		}
		stones = append(stones, stone)
	}

	const numIterations = 25
	for i := range numIterations {
		fmt.Println("After", i, "iterations:", stones)
		stones = blink(stones)
	}

	fmt.Println("Final:", stones)
	fmt.Println("Number of stones:", len(stones))
}
