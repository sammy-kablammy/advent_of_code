package main

import (
	"bufio"
	"fmt"
	"os"
)

type position = struct {
	row, col int
}

func printGrid(grid [][]byte) {
	for i := range grid {
		fmt.Println(string(grid[i]))
	}
}

// Starting at the given point, walk along the trail in an attempt to reach a 9.
// Return all positions of 9s that are reachable from the starting point.
// (Consider a starting position of 9 to be reachable from itself.)
func ascendTrail(grid [][]byte, starting_point position) []position {
	row := starting_point.row
	col := starting_point.col
	current_val := grid[starting_point.row][starting_point.col]

	if current_val < '0' || current_val > '9' {
		panic("Found non-numeric grid value")
	}

	if current_val == '9' {
		return []position{starting_point}
	}

	neighbors := make([]position, 4)
	neighbors[0] = position{row-1, col}
	neighbors[1] = position{row+1, col}
	neighbors[2] = position{row, col-1}
	neighbors[3] = position{row, col+1}

	var nines []position
	for i := range neighbors {
		neighbor := neighbors[i]
		if neighbor.row < 0 || neighbor.row >= len(grid) || neighbor.col < 0 || neighbor.col >= len(grid[0]) {
			continue // out of bounds!
		}
		next_val := grid[neighbor.row][neighbor.col]
		if next_val == current_val+1 {
			nines = append(nines, ascendTrail(grid, neighbor)...) // oooooh rare use of spread operator
		}
	}

	return nines
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var grid [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	totalScore := 0
	for row := range grid {
		for col := range grid {
			if grid[row][col] == '0' {
				nines := ascendTrail(grid, position{row, col})
				// Don't forget to dedupe. (If a 0 forks off into two trails
				// that later rejoin at the same 9, that 0's score is only 1.)
				nines_set := make(map[position]bool)
				for i := range nines {
					if !nines_set[nines[i]] {
						nines_set[nines[i]] = true
					}
				}
				totalScore += len(nines_set)
			}
		}
	}

	fmt.Println("total score:", totalScore)
}
