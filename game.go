package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

const (
	INCORRECT = "incorrect"
	InWord    = "in_word_but_wrong_position"
	CORRECT   = "in_word_and_correct_position"

	MaxGuessCount = 6

	GREEN  = "\x1b[32m"
	PURPLE = "\x1b[35m"
	RED    = "\x1b[31m"
	ENDC   = "\x1b[0m"
)

type Game struct {
	wordSet    map[string]struct{}
	targetWord string
	guessCount int
	wonStatus  bool
	guessesSet []string
}

func NewGame(words []string) *Game {
	wordSet := map[string]struct{}{}
	for _, word := range words {
		if _, ok := wordSet[word]; ok {
			continue
		}
		wordSet[word] = struct{}{}
	}
	var targetWord string
	rand.Seed(time.Now().UnixNano())
	targetWord = words[rand.Intn(len(words))]

	return &Game{
		wordSet:    wordSet,
		targetWord: targetWord,
		guessCount: 0,
		wonStatus:  false,
		guessesSet: []string{},
	}
}

func (g *Game) guessWord(guess string) []string {
	n := len(guess)
	solutionChars := map[string]bool{}
	for i := 0; i < n; i++ {
		solutionChars[g.targetWord[i:i+1]] = true
	}
	ret := make([]string, n)
	for i := 0; i < n; i++ {
		if guess[i] == g.targetWord[i] {
			ret[i] = CORRECT
		} else if _, ok := solutionChars[string(guess[i])]; ok {
			ret[i] = InWord
		} else {
			ret[i] = INCORRECT
		}
	}
	return ret
}

func (g *Game) guessWasGuessedBefore(guess string) bool {
	for _, i := range g.guessesSet {
		if i == guess {
			return true
		}
	}
	return false
}

func (g *Game) setColor(guess string, guessResult []string) string {
	n := len(guess)
	chars := []string{}
	for i := 0; i < n; i++ {
		chars = append(chars, guess[i:i+1])
	}
	var color string
	for i := 0; i < n; i++ {
		c := guess[i : i+1]
		if guessResult[i] == CORRECT {
			color = GREEN
		} else if guessResult[i] == InWord {
			color = PURPLE
		} else if guessResult[i] == INCORRECT {
			color = RED
		}

		chars[i] = color + c + ENDC
	}
	return "" + strings.Join(chars, "") + ENDC
}

func (g *Game) isCorrectGuess() {
	for _, i := range g.guessesSet {
		if i == CORRECT {
			g.wonStatus = true
			break
		}
	}
}
func (g *Game) play(s solver) {
	n := len(g.targetWord)
	for g.guessCount < MaxGuessCount && !g.wonStatus {
		guess := s.makeAGuess()
		if len(guess) != n {
			log.Printf("Guess %s is not of length %d", guess, n)
		}
		if g.guessWasGuessedBefore(guess) {
			log.Printf("Guess %s was already guessed", guess)
			continue
		}
		if _, ok := g.wordSet[guess]; !ok {
			log.Printf("Guess %s is not in the word set", guess)
			continue
		}
		guessObj := g.guessWord(guess)
		for _, i := range guessObj {
			g.guessesSet = append(g.guessesSet, i)
		}
		output := g.setColor(guess, guessObj)
		g.guessCount++
		fmt.Println(output)
		g.isCorrectGuess()
	}
	s.onEndOfGame()
}
