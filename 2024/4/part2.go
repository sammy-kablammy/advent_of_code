package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {

	var grid [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	// return the character at (row, col), or '.' if the position is out of bounds
	get := func(row, col int) byte {
		if row >= len(grid) || row < 0 {
			return '.'
		}
		if col >= len(grid[len(grid)-1]) || col < 0 {
			return '.'
		}
		return grid[row][col]
	}

	count := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] != 'A' {
				continue
			}

			top_left := get(r-1, c-1)
			bottom_left := get(r+1, c-1)
			top_right := get(r-1, c+1)
			bottom_right := get(r+1, c+1)

			has_slash := false
			if top_left == 'M' && bottom_right == 'S' ||
				top_left == 'S' && bottom_right == 'M' {
				has_slash = true
			}
			has_backslash := false
			if top_right == 'M' && bottom_left == 'S' ||
				top_right == 'S' && bottom_left == 'M' {
				has_backslash = true
			}

			if has_slash && has_backslash {
				count += 1
			}
		}
	}

	fmt.Println(count)
}
