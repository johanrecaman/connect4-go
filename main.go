package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/johanrecaman/connect4-go/ui"
)

func main() {
	p := tea.NewProgram(ui.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro ao iniciar o jogo: %v", err)
		os.Exit(1)
	}
}
