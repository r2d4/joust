package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words, err := load("four.txt", 4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rounds := []round{
	// {"test", 0},
	// {"ates", 0},
	// {"oast", 0},
	// {"auto", 1},
	// {"vrow", 0},
	// {"jarl", 1},
	// {"bind", 0},
	// {"caph", 1},

	// {"cede", 1},
	// {"goji", 1},
	// {"naps", 0},
	// {"anon", 0},
	// {"ands", 0},
	// {"orzo", 0},
	// {"daft", 0},
	// {"alan", 1},
	}
	// rounds := []round{
	// 	// {"dude", 0},
	// 	// {"writ", 1},
	// 	// {"myth", 1},
	// 	// {"cwms", 1},
	// 	// {"byrl", 0},
	// 	// {"wych", 0},
	// 	// {"typp", 1},
	// 	// {"goop", 0},
	// 	// {"mass", 2},
	// 	// {"damn", 1},
	// 	{"moon", 1},
	// 	{"scry", 1},
	// 	{"typp", 0},
	// 	{"bass", 0},
	// 	{"word", 2},
	// 	{"gore", 2},
	// 	{"chew", 2},
	// 	{"work", 1},
	// 	{"drew", 2},
	// 	{"ogre", 2},
	// 	{"quiz", 0},
	// }

	// cliPlay(words)
	var guess string
	blacklist := []string{}
	sols := playRounds(words, rounds...)
	m := reverseIndex(words)
	guess = bestword(sols, words, m, blacklist)
	totalRounds := 0
	for {
		var word string
		var score int
		fmt.Printf("===ROUND %d===\n", totalRounds)
		fmt.Println("enter word")
		fmt.Scanln(&word)
		for word != "x" && len(word) != 4 {
			fmt.Println("pick another")
			fmt.Scanln(&word)
		}
		for word == "x" {
			fmt.Println("removing", guess)
			blacklist = append(blacklist, guess)
			fmt.Println("calculating next guess...")
			guess = bestword(sols, words, m, blacklist)
			fmt.Println("enter word")
			fmt.Scanln(&word)
		}
		fmt.Println("enter score")
		fmt.Scanln(&score)
		totalRounds = totalRounds + 1
		word = strings.ToUpper(word)
		addResult(word, score, sols, m, words)
		fmt.Println("calculating best word...")
		guess = bestword(sols, words, m, blacklist)
	}
}
