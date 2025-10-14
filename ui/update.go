package ui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/johanrecaman/connect4-go/game"
)

type aiMoveMsg struct {
	result game.MoveResult
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case menuState:
			return m.updateMenu(msg)
		case playingState:
			return m.updatePlaying(msg)
		case gameOverState:
			return m.updateGameOver(msg)
		}
	case aiMoveMsg:
		return m.handleAIMove(msg), nil
	}
	return m, nil
}

func (m Model) updateMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < 2 {
			m.cursor++
		}
	case "enter":
		m.level = game.AILevel(m.cursor)
		m.state = playingState
		m.message = "Humano (🔴) começa! Use A, S, D, F, G, H, J para escolher a coluna."
	}
	return m, nil
}

func (m Model) updatePlaying(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	}

	if m.currentTurn == 2 {
		return m, nil
	}

	col := -1
	switch msg.String() {
	case "a", "A":
		col = 0
	case "s", "S":
		col = 1
	case "d", "D":
		col = 2
	case "f", "F":
		col = 3
	case "g", "G":
		col = 4
	case "h", "H":
		col = 5
	case "j", "J":
		col = 6
	}

	if col != -1 {
		if m.board.MakeMove(col, 1) {
			if m.board.CheckWin(1) {
				m.winner = 1
				m.state = gameOverState
				m.message = "🎉 Humano (🔴) venceu!"
				return m, nil
			}
			if m.board.IsFull() {
				m.winner = 0
				m.state = gameOverState
				m.message = "⚖️  Empate!"
				return m, nil
			}
			m.currentTurn = 2
			m.message = "IA (🟡) está pensando..."
			return m, m.makeAIMove()
		} else {
			m.message = "❌ Coluna cheia ou inválida! Tente outra."
		}
	}

	return m, nil
}

func (m Model) updateGameOver(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "r":
		return NewModel(), nil
	}
	return m, nil
}

func (m Model) makeAIMove() tea.Cmd {
	return func() tea.Msg {
		result := game.GetBestMove(&m.board, m.level)
		return aiMoveMsg{result: result}
	}
}

func (m Model) handleAIMove(msg aiMoveMsg) Model {
	result := msg.result
	m.lastAITime = result.Duration
	m.lastAIScore = result.Score

	if result.Column != -1 {
		m.board.MakeMove(result.Column, 2)
		if m.board.CheckWin(2) {
			m.winner = 2
			m.state = gameOverState
			m.message = "🤖 IA (🟡) venceu!"
			return m
		}
		if m.board.IsFull() {
			m.winner = 0
			m.state = gameOverState
			m.message = "⚖️  Empate!"
			return m
		}
		m.currentTurn = 1
		m.message = "Sua vez! Use A, S, D, F, G, H, J para escolher a coluna."
	}

	return m
}
