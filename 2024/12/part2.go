package main

import (
	"bufio"
	"fmt"
	"os"
)

// some observations:
// - regions are not guaranteed to start at the letter 'A'
// - multiple disjoint regions can share a letter

type Point = struct {
	row, col int
}

func isInBounds(grid [][]byte, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[row])
}

func main() {
	// parse input file
	scanner := bufio.NewScanner(os.Stdin)
	var grid [][]byte
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
		// fmt.Println(scanner.Text())
	}

	debugLetter := make(map[int]byte) // map region numbers to their labels

	// find all the different regions in the input
	region := make(map[Point]int)
	// maps use 0 for nonexistant keys, so region numbers need to be 1-based
	regionNumber := 1
	for row := range grid {
		for col := range grid[row] {

			if isPointAlreadyExplored := region[Point{row, col}] >= 1; isPointAlreadyExplored {
				continue
			}

			// run BFS to find all points in this region
			frontier := []Point{{row, col}}
			for i := 0; i < len(frontier); i++ {
				r := frontier[i].row
				c := frontier[i].col
				if isPointAlreadyExplored := region[Point{r, c}] >= 1; isPointAlreadyExplored {
					continue
				}
				// fmt.Println("Point", r, c, "is in region", regionNumber)
				region[Point{r, c}] = regionNumber
				neighbors := []Point{{r - 1, c}, {r + 1, c}, {r, c - 1}, {r, c + 1}}
				for _, neighbor := range neighbors {
					if !isInBounds(grid, neighbor.row, neighbor.col) || region[neighbor] >= 1  {
						continue
					}
					if grid[neighbor.row][neighbor.col] == grid[r][c] {
						frontier = append(frontier, neighbor)
					}
				}
			}
			debugLetter[regionNumber] = grid[row][col]
			regionNumber++

		}
	}

	// same map as above but reversed; this is a surprise tool that will help us later
	points := make(map[int][]Point)
	for point, regionNum := range region {
		points[regionNum] = append(points[regionNum], point)
	}

	// compute stats of each region
	areas := make(map[int]int)
	for regionNum, points := range points {
		areas[regionNum] = len(points)
	}
	// the simplest region must have 4 sides, and every turn along its boundary
	// adds one more. so, instead of counting sides, we'll count turns. the
	// exact strategy i'm using for counting corners comes from here:
	// https://www.reddit.com/r/adventofcode/comments/1hcdnk0/comment/m1nkmol/
	sides := make(map[int]int)
	// return the region at the given position, or -1 if out of bounds
	get := func (row, col int) int {
		if isInBounds(grid, row, col) {
			return region[Point{row, col}]
		}
		return -1
	}
	for regionNum, points := range points {
		for _, point := range points {
			r := point.row
			c := point.col
			me := region[Point{r,c}]
			up := get(r-1, c)
			down := get(r+1, c)
			right := get(r, c+1)
			left := get(r, c-1)
			// check for corner in upper left
			if (up != me && left != me) ||
				(up == me && left == me && get(r-1, c-1) != me) {
				sides[regionNum]++
			}
			// check for corner in upper right
			if (up != me && right != me) ||
				(up == me && right == me && get(r-1, c+1) != me) {
				sides[regionNum]++
			}
			// check for corner in lower left
			if (down != me && left != me) ||
				(down == me && left == me && get(r+1, c-1) != me) {
				sides[regionNum]++
			}
			// check for corner in lower right
			if (down != me && right != me) ||
				(down == me && right == me && get(r+1, c+1) != me) {
				sides[regionNum]++
			}
		}
	}

	finalResult := 0
	for regionNum := range points {
		area := areas[regionNum]
		sides := sides[regionNum]
		fmt.Println("Region number", regionNum, "(" + string(debugLetter[regionNum]) + ")", "has area", area, "and this many sides:", sides)
		finalResult += area * sides
	}

	fmt.Println("final result:", finalResult)
}
