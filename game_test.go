package main

import (
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want *Game
	}{
		{
			name: "new game",
			args: args{
				words: []string{"a", "b", "c"},
			},
			want: &Game{
				wordSet:    map[string]struct{}{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
				targetWord: "c",
				guessCount: 0,
				wonStatus:  false,
				guessesSet: []string{},
			},
		},
		{
			name: "new game with duplicate words",
			args: args{
				words: []string{"a", "b", "c", "a"},
			},
			want: &Game{
				wordSet:    map[string]struct{}{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
				targetWord: "a",
				guessCount: 0,
				wonStatus:  false,
				guessesSet: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGame(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_guessWasGuessedBefore(t *testing.T) {
	type fields struct {
		wordSet    map[string]struct{}
		targetWord string
		guessCount int
		wonStatus  bool
		guessesSet []string
	}
	type args struct {
		guess string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "guess was guessed before",
			fields: fields{
				wordSet:    map[string]struct{}{"a": struct{}{}},
				targetWord: "a",
				guessCount: 0,
				wonStatus:  false,
				guessesSet: []string{"a"},
			},
			args: args{
				guess: "a",
			},
			want: true,
		},
		{
			name: "guess was not guessed before",
			fields: fields{
				wordSet:    map[string]struct{}{"a": struct{}{}},
				targetWord: "a",
				guessCount: 0,
				wonStatus:  false,
				guessesSet: []string{},
			},
			args: args{
				guess: "b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				wordSet:    tt.fields.wordSet,
				targetWord: tt.fields.targetWord,
				guessCount: tt.fields.guessCount,
				wonStatus:  tt.fields.wonStatus,
				guessesSet: tt.fields.guessesSet,
			}
			if got := g.guessWasGuessedBefore(tt.args.guess); got != tt.want {
				t.Errorf("guessWasGuessedBefore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_guessWord(t *testing.T) {
	type fields struct {
		wordSet    map[string]struct{}
		targetWord string
		guessCount int
		wonStatus  bool
		guessesSet []string
	}
	type args struct {
		guess string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "guess word",
			fields: fields{
				wordSet:    map[string]struct{}{"a": struct{}{}},
				targetWord: "a",
				guessCount: 0,
				wonStatus:  false,
				guessesSet: []string{},
			},
			args: args{
				guess: "a",
			},
			want: []string{"in_word_and_correct_position"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				wordSet:    tt.fields.wordSet,
				targetWord: tt.fields.targetWord,
				guessCount: tt.fields.guessCount,
				wonStatus:  tt.fields.wonStatus,
				guessesSet: tt.fields.guessesSet,
			}
			if got := g.guessWord(tt.args.guess); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("guessWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
