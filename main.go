package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	white = lipgloss.Color("#ffffff")
	grey  = lipgloss.Color("#bbbbbb")

	cursorStyle = lipgloss.NewStyle().Underline(true).Foreground(grey)
	beforeStyle = lipgloss.NewStyle().Foreground(grey)
	afterStyle  = lipgloss.NewStyle().Foreground(white)
	textStyle   = lipgloss.NewStyle().Align(lipgloss.Center).Height(3)
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
	text   string
	cursor int
	width  int
	height int
}

func newModel() model {
	return model{
		text: "in this world is the destiny of mankind controlled by some transcendental entity or law",
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
			m.cursor = 0
		default:
			if m.cursor < len(m.text) {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	h := lipgloss.Place(m.width, 6, lipgloss.Center, lipgloss.Center, header)
	var t string
	if m.cursor >= len(m.text) {
		t = "Speed: 65 wpm\nAccuracy: 100%"
	} else {
		t = fmt.Sprintf("%s%s%s",
			afterStyle.Render(m.text[:m.cursor]),
			cursorStyle.Render(string(m.text[m.cursor])),
			beforeStyle.Render(m.text[m.cursor+1:]),
		)
	}
	return lipgloss.JoinVertical(
		lipgloss.Center,
		h,
		lipgloss.Place(
			m.width,
			m.height-lipgloss.Height(h),
			lipgloss.Center,
			0.8,
			textStyle.Width(int(float32(m.width)*0.7)).Render(t),
		),
	)
}
