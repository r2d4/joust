package main

import (
	"fmt"
	"os"
)

func main() {
	words, err := load("four.txt", 4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rounds := []round{
		// {"dude", 0},
		// {"writ", 1},
		// {"myth", 1},
		// {"cwms", 1},
		// {"byrl", 0},
		// {"wych", 0},
		// {"typp", 1},
		// {"goop", 0},
		// {"mass", 2},
		// {"damn", 1},
		{"moon", 1},
		{"scry", 1},
		{"typp", 0},
		{"bass", 0},
		{"word", 2},
		{"gore", 2},
		{"chew", 2},
		{"work", 1},
		{"drew", 2},
		{"ogre", 2},
		{"quiz", 0},
	}
	sols := play(words, rounds...)
	fmt.Println(bitmapToSlice(words, sols))
	stats(sols, words)
}
