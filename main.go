package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	maxWordLength = 5
)

var validWordsList []string

func init() {
	allWordsList := "words.txt"
	f, err := os.Open(allWordsList)
	if err != nil {
		panic(err.Error())
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if len(word) != maxWordLength {
			continue
		}
		validWordsList = append(validWordsList, strings.TrimSpace(scanner.Text()))
	}

	if err = scanner.Err(); err != nil {
		panic(err.Error())
	}
}

func main() {
	fmt.Println("Welcome to Wordle Solver!")
	play()
}

func play() {
	game := NewGame(validWordsList)
	game.play(solver{
		game,
	})
	var playAgain string
	fmt.Println("Do you want to play again? (y/n)")
	fmt.Scanln(&playAgain)
	if playAgain == "y" {
		play()
	} else {
		fmt.Println("Bye!")
	}
}
