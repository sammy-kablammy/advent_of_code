package main

import (
	"bufio"
	"fmt"
	"os"
)

type direction int
const (
	up direction = iota
	right
	down
	left
)

func print_grid(grid [][]byte) {
	for row := range grid {
		fmt.Println(string(grid[row]))
	}
	fmt.Println("")
}

func get_next_position(current_row, current_col int, current_direction direction) (next_row, next_col int) {
	next_row = current_row
	next_col = current_col
	switch current_direction {
	case up:
		next_row -= 1
	case down:
		next_row += 1
	case left:
		next_col -= 1
	case right:
		next_col += 1
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var grid [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	// find starting position
	current_row := -1
	current_col := -1
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '^' {
				current_row = r
				current_col = c
			}
		}
	}
	if current_row == -1 || current_col == -1 {
		_ = fmt.Errorf("starting position not found")
		os.Exit(1)
	}

	current_direction := up
	distance := 1 // (distance includes starting position, as per problem)

	next_row, next_col := get_next_position(current_row, current_col, current_direction)
	for next_row >= 0 && next_row < len(grid) && next_col >= 0 && next_col < len(grid[0]) {
		if grid[next_row][next_col] == '#' {
			// turn right
			current_direction = (current_direction + 1) % 4
		} else {
			// move forward
			current_row = next_row
			current_col = next_col
			if grid[current_row][current_col] == '.' {
				grid[current_row][current_col] = 'X'
				distance += 1
			}
		}
		// print_grid(grid)
		next_row, next_col = get_next_position(current_row, current_col, current_direction)
	}

	fmt.Println("distance:", distance)
}
