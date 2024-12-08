package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position = struct {
	row, col int
}

// from the perspective of 'source', return the relative change in position
// needed to reach 'destination'.
func get_relative_position(source, destination Position) Position {
	return Position{destination.row - source.row, destination.col - source.col}
}

func is_in_bounds(grid [][]byte, position Position) bool {
	return position.row >= 0 && position.row < len(grid) &&
		position.col >= 0 && position.col < len(grid[position.row])
}

// what's different about part 2? antinodes can occur at antennas so long as
// there are at least two antennas of the same frequency? like, if there are two
// antennas of frequency A, both of those reside on antinodes; if there is only
// one antenna of frequency B, it is not an antinode. okay cool.

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var grid [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	antennas := make(map[byte][]Position)
	for row := range grid {
		for col := range grid[row] {
			cell := grid[row][col]
			if cell == '.' {
				continue
			}
			antennas[cell] = append(antennas[cell], Position{row, col})
		}
	}

	antinode_count := 0

	for _, positions := range antennas {
		for position1_idx := range positions {
			position1 := positions[position1_idx]

			// account for part 2's antenna same-cell antinodes
			if len(positions) > 1 {
				grid[position1.row][position1.col] = '#'
			}

			for position2_idx := range positions {
				if position1_idx == position2_idx {
					continue
				}
				position2 := positions[position2_idx]
				relative := get_relative_position(position1, position2)
				antinode := Position{position1.row - relative.row, position1.col - relative.col}
				for is_in_bounds(grid, antinode) {
					grid[antinode.row][antinode.col] = '#'
					antinode = Position{antinode.row - relative.row, antinode.col - relative.col}
				}
			}
		}
	}

	for row := range grid {
		for col := range grid {
			if grid[row][col] == '#' {
				antinode_count += 1
			}
		}
		fmt.Println(string(grid[row]))
	}
	fmt.Println("number of antinodes:", antinode_count)

}
