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
	perimeters := make(map[int]int)
	for regionNum, points := range points {
		for _, point := range points {
			r := point.row
			c := point.col
			perimeter := 4
			neighbors := []Point{{r - 1, c}, {r + 1, c}, {r, c - 1}, {r, c + 1}}
			for _, neighbor := range neighbors {
				if !isInBounds(grid, neighbor.row, neighbor.col) {
					continue
				}
				if grid[neighbor.row][neighbor.col] == grid[r][c] {
					perimeter -= 1
				}
			}
			// fmt.Println("Cell", r, c, "contributes", perimeter, "perimeter to its region")
			perimeters[regionNum] += perimeter
		}
	}

	finalResult := 0
	for regionNum := range points {
		area := areas[regionNum]
		perimeter := perimeters[regionNum]
		// fmt.Println("Region number", regionNum, "(" + string(debugLetter[regionNum]) + ")", "has area", area, "and perimeter", perimeter)
		finalResult += area * perimeter
	}

	fmt.Println("final result:", finalResult)
}
