package main

import (
	"bufio"
	"fmt"
	"os"
)

// some planning:
// we can detect when we're caught in a loop when:
// - the current position is already marked (i.e. we've been here before)
// - AND, the current direction is the same as it was the last time we were here
//
// i think the simplest adjustment to make this happen is by, instead of marking
// visited squares with an X, we use either U, D, R, or L, based on the
// direction we were traveling when visiting that square in the past.

type direction int
const (
	up direction = iota
	right
	down
	left
)
var direction_chars = [...]byte{'U', 'R', 'D', 'L'}

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

// i'm like 90% sure that the builtin copy() is a shallow copy, necessitating this 2D copy:
func slice_2d_copy(source [][]byte) [][]byte {
	// var destination [][]byte
	destination := make([][]byte, len(source))
	for r := range source {
		destination[r] = make([]byte, len(source[r]))
		copy(destination[r], source[r])
	}
	return destination
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var original_grid [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		original_grid = append(original_grid, []byte(line))
	}

	// find starting position
	starting_row := -1
	starting_col := -1
	for r := range original_grid {
		for c := range original_grid[r] {
			if original_grid[r][c] == '^' {
				starting_row = r
				starting_col = c
			}
		}
	}
	if starting_row == -1 || starting_col == -1 {
		_ = fmt.Errorf("starting position not found")
		os.Exit(1)
	}

	num_possible_loops := 0
	for row := range original_grid {
		for col := range original_grid[row] {
			if original_grid[row][col] == '#' || original_grid[row][col] == '^' {
				// (recall that the starting point can't be an obstruction)
				continue
			}
			grid := slice_2d_copy(original_grid)
			// try placing an obstruction
			grid[row][col] = 'O'
			// see if the newly placed obstruction induces a loop
			current_row := starting_row
			current_col := starting_col
			current_direction := up
			next_row, next_col := get_next_position(current_row, current_col, current_direction)
			for next_row >= 0 && next_row < len(grid) && next_col >= 0 && next_col < len(grid[0]) {
				next_cell := grid[next_row][next_col]
				if direction_chars[current_direction] == next_cell {
					num_possible_loops += 1
					break
				}
				if next_cell == '#' || next_cell == 'O' {
					// turn right
					current_direction = (current_direction + 1) % 4
				} else {
					// move forward
					current_row = next_row
					current_col = next_col
					if grid[current_row][current_col] == '.' {
						grid[current_row][current_col] = direction_chars[current_direction]
					}
				}
				// print_grid(grid)
				next_row, next_col = get_next_position(current_row, current_col, current_direction)
			}
			// reset and try another obstruction placement
			grid[row][col] = '.'
		}
	}

	fmt.Println("num possible loops:", num_possible_loops)
}
