package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	m := newModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start the program: %v", err)
		os.Exit(1)
	}
}

const header = `
▄▀▀ █▀▄ █▀▄ █▄█ █▀▀ █▀▄   ▀█▀ ▀▄▀ █▀▄ █▀▀
█▄█ █▄█ █▀▀ █ █ ██▄ █▀▄    █   █  █▀▀ ██▄
	`

type model struct {
	text   []rune
	width  int
	height int
}

func newModel() model {
	return model{
		text: []rune("in this world is the destiny of mankind controlled by some transcendental entity or law"),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return newTypingScreen(m.text, m.width, m.height), nil
		}
	}
	return m, nil
}

func (m model) View() string {
	h := lipgloss.Place(m.width, 6, lipgloss.Center, lipgloss.Center, header)
	return lipgloss.JoinVertical(
		lipgloss.Center,
		h,
		"Press enter to start",
	)
}
