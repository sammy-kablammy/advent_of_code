// iiiiiiiiiit's golang today! hooray!!!! we love go :)

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
			if grid[r][c] != 'X' {
				continue
			}

			// rightward
			if get(r, c+1) == 'M' &&
				get(r, c+2) == 'A' &&
				get(r, c+3) == 'S' {
					count += 1
			}
			// leftward
			if get(r, c-1) == 'M' &&
				get(r, c-2) == 'A' &&
				get(r, c-3) == 'S' {
					count += 1
			}

			// downward
			if get(r+1, c) == 'M' &&
				get(r+2, c) == 'A' &&
				get(r+3, c) == 'S' {
					count += 1
			}
			// upward
			if get(r-1, c) == 'M' &&
				get(r-2, c) == 'A' &&
				get(r-3, c) == 'S' {
					count += 1
			}

			// up, right
			if get(r-1, c+1) == 'M' &&
				get(r-2, c+2) == 'A' &&
				get(r-3, c+3) == 'S' {
					count += 1
			}
			// up, left
			if get(r-1, c-1) == 'M' &&
				get(r-2, c-2) == 'A' &&
				get(r-3, c-3) == 'S' {
					count += 1
			}
			// down, right
			if get(r+1, c+1) == 'M' &&
				get(r+2, c+2) == 'A' &&
				get(r+3, c+3) == 'S' {
					count += 1
			}
			// down, left
			if get(r+1, c-1) == 'M' &&
				get(r+2, c-2) == 'A' &&
				get(r+3, c-3) == 'S' {
					count += 1
			}
		}
	}

	fmt.Println(count)
}
