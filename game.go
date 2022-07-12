package main

import (
	"log"
	"math/rand"
	"time"
)

const (
	INCORRECT = "incorrect"
	InWord    = "in_word_but_wrong_position"
	CORRECT   = "in_word_and_correct_position"

	MaxGuessCount = 6
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
		s.game.guessWord(guess)
		guessObj := s.makeAGuess()
		g.guessesSet = append(g.guessesSet, guessObj)
		g.guessCount++
		solution := g.guessWord(guess)
		log.Printf("Guess %s: %v", guess, solution)
		if g.wonStatus {
			log.Printf("Guess %s won", guess)
			break
		}
	}
}
