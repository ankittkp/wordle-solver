package main

import "testing"

func Test_solver_makeAGuess(t *testing.T) {
	type fields struct {
		game *Game
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "makeAGuess",
			fields: fields{
				game: &Game{
					wordSet:    map[string]struct{}{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
					targetWord: "a",
					guessCount: 0,
					wonStatus:  false,
					guessesSet: []string{},
				},
			},
			want: "",
		},
		{
			name: "makeAGuess",
			fields: fields{
				game: &Game{
					wordSet:    map[string]struct{}{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
					targetWord: "c",
					guessCount: 0,
					wonStatus:  false,
					guessesSet: []string{"a", "b"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &solver{
				game: tt.fields.game,
			}
			if got := s.makeAGuess(); got != tt.want {
				t.Errorf("makeAGuess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solver_onEndOfGame(t *testing.T) {
	type fields struct {
		game *Game
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "onEndOfGame",
			fields: fields{
				game: &Game{
					wordSet:    map[string]struct{}{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
					targetWord: "c",
					guessCount: 0,
					wonStatus:  false,
					guessesSet: []string{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &solver{
				game: tt.fields.game,
			}
			s.onEndOfGame()
		})
	}
}
