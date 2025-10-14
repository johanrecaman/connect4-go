package ui

import (
	"time"

	"github.com/johanrecaman/connect4-go/game"
)

type gameState int

const (
	menuState gameState = iota
	playingState
	gameOverState
)

type Model struct {
	board       game.Board
	state       gameState
	level       game.AILevel
	currentTurn int
	winner      int
	lastAITime  time.Duration
	lastAIScore int
	message     string
	cursor      int
}

func NewModel() Model {
	return Model{
		board:       game.NewBoard(),
		state:       menuState,
		currentTurn: 1,
		cursor:      0,
	}
}
