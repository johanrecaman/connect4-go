package ui

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	switch m.state {
	case menuState:
		return m.viewMenu()
	case playingState:
		return m.viewPlaying()
	case gameOverState:
		return m.viewGameOver()
	}
	return ""
}

func (m Model) viewMenu() string {
	s := strings.Builder{}
	s.WriteString("\n")
	s.WriteString("  ╔════════════════════════════════════════╗\n")
	s.WriteString("  ║     🎮 CONNECT FOUR - 6x7 🎮          ║\n")
	s.WriteString("  ╚════════════════════════════════════════╝\n\n")
	s.WriteString("  Selecione o nível de dificuldade:\n\n")

	levels := []string{"🟢 Iniciante", "🟠 Intermediário", "🔴 Profissional"}
	for i, level := range levels {
		cursor := " "
		if m.cursor == i {
			cursor = "→"
		}
		s.WriteString(fmt.Sprintf("    %s %s\n", cursor, level))
	}

	s.WriteString("\n  Use ↑/↓ para navegar, Enter para selecionar\n")
	s.WriteString("  Pressione 'q' para sair\n")
	return s.String()
}

func (m Model) viewPlaying() string {
	s := strings.Builder{}
	s.WriteString("\n")
	s.WriteString("  ╔════════════════════════════════════════╗\n")
	s.WriteString("  ║     🎮 CONNECT FOUR - 6x7 🎮          ║\n")
	s.WriteString("  ╚════════════════════════════════════════╝\n\n")

	s.WriteString("     A   S   D   F   G   H   J\n")
	s.WriteString("  ┌───┬───┬───┬───┬───┬───┬───┐\n")

	for i := 0; i < 6; i++ {
		s.WriteString("  │")
		for j := 0; j < 7; j++ {
			cell := m.board.Grid[i][j]
			symbol := "   "
			switch cell {
			case 1:
				symbol = " 🔴"
			case 2:
				symbol = " 🟡"
			}
			s.WriteString(symbol)
			s.WriteString("│")
		}
		s.WriteString("\n")
		if i < 5 {
			s.WriteString("  ├───┼───┼───┼───┼───┼───┼───┤\n")
		}
	}
	s.WriteString("  └───┴───┴───┴───┴───┴───┴───┘\n")
	s.WriteString("    0   1   2   3   4   5   6\n\n")

	s.WriteString(fmt.Sprintf("  %s\n\n", m.message))

	if m.currentTurn == 1 && m.lastAITime > 0 {
		s.WriteString(fmt.Sprintf("  ⏱️  Tempo da última jogada da IA: %v\n", m.lastAITime))
		s.WriteString(fmt.Sprintf("  📊 Avaliação do estado: %d\n\n", m.lastAIScore))
	}

	s.WriteString("  Pressione 'q' para sair\n")
	return s.String()
}

func (m Model) viewGameOver() string {
	s := strings.Builder{}
	s.WriteString("\n")
	s.WriteString("  ╔════════════════════════════════════════╗\n")
	s.WriteString("  ║          🎮 FIM DE JOGO 🎮            ║\n")
	s.WriteString("  ╚════════════════════════════════════════╝\n\n")

	s.WriteString("     A   S   D   F   G   H   J\n")
	s.WriteString("  ┌───┬───┬───┬───┬───┬───┬───┐\n")

	for i := 0; i < 6; i++ {
		s.WriteString("  │")
		for j := 0; j < 7; j++ {
			cell := m.board.Grid[i][j]
			symbol := "   "
			switch cell {
			case 1:
				symbol = " 🔴"
			case 2:
				symbol = " 🟡"
			}
			s.WriteString(symbol)
			s.WriteString("│")
		}
		s.WriteString("\n")
		if i < 5 {
			s.WriteString("  ├───┼───┼───┼───┼───┼───┼───┤\n")
		}
	}
	s.WriteString("  └───┴───┴───┴───┴───┴───┴───┘\n")
	s.WriteString("    0   1   2   3   4   5   6\n\n")

	s.WriteString(fmt.Sprintf("  %s\n\n", m.message))

	if m.lastAITime > 0 {
		s.WriteString(fmt.Sprintf("  ⏱️  Tempo da última jogada da IA: %v\n", m.lastAITime))
		s.WriteString(fmt.Sprintf("  📊 Avaliação final: %d\n\n", m.lastAIScore))
	}

	s.WriteString("  Pressione 'r' para jogar novamente\n")
	s.WriteString("  Pressione 'q' para sair\n")
	return s.String()
}
