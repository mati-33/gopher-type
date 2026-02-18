package main

import (
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	white = lipgloss.Color("#ffffff")
	grey  = lipgloss.Color("#bbbbbb")
	red   = lipgloss.Color("#ff005f")

	cursorStyle = lipgloss.NewStyle().Underline(true).Foreground(grey)
	beforeStyle = lipgloss.NewStyle().Foreground(grey)
	afterStyle  = lipgloss.NewStyle().Foreground(white)
	errorStyle  = lipgloss.NewStyle().Foreground(red)
)

type typingScreen struct {
	text     []rune
	errors   []int
	cursor   int
	width    int
	height   int
	finished bool
}

func newTypingScreen(text []rune, width, height int) typingScreen {
	return typingScreen{
		text:     text,
		errors:   []int{},
		cursor:   0,
		width:    width,
		height:   height,
		finished: false,
	}
}

func (s typingScreen) Init() tea.Cmd {
	return nil
}

func (s typingScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return s, tea.Quit
		case "enter":
			s.cursor = 0
			s.errors = []int{}
			s.finished = false
		default:
			current := string(s.text[s.cursor])
			if msg.String() != current {
				s.errors = append(s.errors, s.cursor)
			}
			if s.cursor < len(s.text)-1 {
				s.cursor++
			} else {
				s.finished = true
			}
			return s, nil
		}
	}
	return s, nil
}

func (s typingScreen) View() string {
	var view string
	if s.finished {
		view = "finish"

	} else {
		b := strings.Builder{}

		for idx, ch := range s.text[:s.cursor] {
			if slices.Contains(s.errors, idx) {
				b.WriteString(errorStyle.Render(string(ch)))
			} else {
				b.WriteString(afterStyle.Render(string(ch)))
			}
		}

		b.WriteString(cursorStyle.Render(string(s.text[s.cursor])))
		b.WriteString(beforeStyle.Render(string(s.text[s.cursor+1:])))

		view = b.String()

	}
	return lipgloss.Place(s.width, s.height, lipgloss.Center, lipgloss.Center, view)
}
