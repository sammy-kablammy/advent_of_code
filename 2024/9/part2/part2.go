package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

const empty_block = -1

type file = struct {
	id int
	size int
}

func print_filesystem(filesystem []file) {
	for file_idx := range filesystem {
		for range filesystem[file_idx].size {
			if filesystem[file_idx].id == -1 {
				fmt.Print(".",)
			} else {
				fmt.Print(filesystem[file_idx].id)
			}
		}
		fmt.Print(" ")
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
	var filesystem []file
	id := 0
	is_gap := false
	for i := range input {
		size, err := strconv.Atoi(string(input[i]))
		var val int
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
		filesystem = append(filesystem, file{id: val, size: size})
	}

	// defragment filesystem
	for right_file_idx := len(filesystem)-1; right_file_idx >= 0; right_file_idx-- {
		right_file := filesystem[right_file_idx]
		if right_file.id == empty_block {
			continue
		}
		for left_file_idx := 0; left_file_idx < len(filesystem) && left_file_idx < right_file_idx; left_file_idx++ {
			left_file := filesystem[left_file_idx]
			space_remaining := left_file.size - right_file.size
			if left_file.id != empty_block {
				// can only swap into an empty chunk of blocks
				continue
			}
			if left_file.size < right_file.size {
				// this chunk of blocks is too small to hold the right file
				continue
			}
			// let's do a swap!
			filesystem[right_file_idx].id = empty_block
			filesystem[left_file_idx].id = right_file.id
			if left_file.size == right_file.size {
				// perfect swap; the right chunk fits exactly into the left chunk
			} else if left_file.size > right_file.size {
				// imperfect swap; the right chunk moves into the left, but
				// there is still space remaining in the left
				filesystem[left_file_idx].size = right_file.size
				// insert new file entry to represent the left block's leftover space
				new_file := file{id: empty_block, size: space_remaining}
				filesystem = slices.Insert(filesystem, left_file_idx + 1, new_file)
			}
			break
		}
	}

	block_idx := 0
	checksum := 0
	for file_idx := 0; file_idx < len(filesystem); file_idx++ {
		file := filesystem[file_idx]
		if file.id == empty_block {
			block_idx += file.size
			continue
		}
		// gauss is FUMING right now
		for range file.size {
			add := block_idx * file.id
			checksum += add
			block_idx++
		}
	}
	fmt.Println("checksum:", checksum)
}

// today i had another oopsie, lasting about 2 hours. this time, my copied
// solution was wrong. so i tmux appended the copied text to my clipboard
// instead of overwriting it. tmux and i are gonna have a stern talking to 💢

// apparently this is a kitty bug? https://superuser.com/questions/1558823/when-i-copy-from-something-on-the-cli-in-tmux-it-adds-to-the-buffer-instead-of-o
// no, it happens in gnome-terminal too. idk. this is annoying
