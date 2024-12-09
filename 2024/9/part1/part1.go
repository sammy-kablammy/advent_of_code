package main

// the problem is a little unclear how to handle IDs greater than 9. just forget
// the part where they represent the filesystem as a string ("00...111...2" or
// whatever). think of it as a list of integers, where each integer is either
// the ID of that block or something like -1 for an empty block.

// from there, this problem is basically just a cursed selection sort

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

const empty_block = -1

func print_filesystem(filesystem []int64) {
	for block := range filesystem {
		if filesystem[block] == empty_block {
			fmt.Print(".", " ")
		} else {
			fmt.Print(filesystem[block], " ")
		}
	}
	fmt.Println()
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	input = input[:len(input)-1] // remove trailing \n

	// construct filesystem
	var filesystem []int64
	id := int64(0)
	is_gap := false
	for i := range input {
		num, err := strconv.Atoi(string(input[i]))
		var val int64
		if err != nil {
			panic(err)
		}
		if is_gap {
			val = empty_block
		} else {
			val = id
			id += 1
		}
		is_gap = !is_gap
		for range num {
			filesystem = append(filesystem, val)
		}
	}

	// do filesystem consolidation
	left := 0
	right := len(filesystem) - 1
	for {
		for filesystem[left] != empty_block {
			left += 1
		}
		for filesystem[right] == empty_block {
			right -= 1
		}
		if left >= right {
			break
		}
		temp := filesystem[left]
		filesystem[left] = filesystem[right]
		filesystem[right] = temp
		// fmt.Println("left:", left, "   right:", right)
		// print_filesystem(filesystem)
	}

	checksum := int64(0)
	for i := range filesystem {
		if filesystem[i] == empty_block {
			break
		}
		checksum += int64(i) * filesystem[i]
	}
	fmt.Println("checksum:", checksum)
}

// oof spent like an hour trying to debug but i copied the input incorrectly
// 💀💀💀
