package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type choice int

const (
	practise choice = iota
	settings
	about
	exit
)

const header = `
▄▀▀ █▀▄ █▀▄ █▄█ █▀▀ █▀▄   ▀█▀ ▀▄▀ █▀▄ █▀▀
█▄█ █▄█ █▀▀ █ █ ██▄ █▀▄    █   █  █▀▀ ██▄
	`

var (
	headerStyle        = lipgloss.NewStyle().Align(lipgloss.Center).MarginTop(1).MarginBottom(3)
	choiceStyle        = lipgloss.NewStyle()
	focusedChoiceStyle = lipgloss.NewStyle().Background(lipgloss.Color("#303234"))
)

type welcomeScreen struct {
	choices []string
	cursor  int
	width   int
	height  int
}

func newWelcomeScreen(width, height int) welcomeScreen {
	return welcomeScreen{
		choices: []string{
			"Practise",
			"Settings",
			"About",
			"Exit",
		},
		width:  width,
		height: height,
	}
}

func (s welcomeScreen) Init() tea.Cmd {
	return nil
}

func (s welcomeScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
		return s, nil

	case tea.KeyMsg:
		switch msg.String() {

		case "j":
			if s.cursor < len(s.choices)-1 {
				s.cursor++
			}

		case "k":
			if s.cursor > 0 {
				s.cursor--
			}

		case "enter":
			switch s.cursor {
			case int(practise):
				return s, func() tea.Msg {
					return PushScreen{
						screen: newTypingScreen(
							[]rune("in this world is the destiny of mankind controlled by some transcendental entity or law"),
							s.width,
							s.height,
						),
					}
				}
			case int(settings):
			case int(about):
			case int(exit):
				return s, tea.Quit
			}
		}
	}
	return s, nil
}

func (s welcomeScreen) View() string {
	choicesWidth := lipgloss.Width(header)
	b := strings.Builder{}

	for idx, choice := range s.choices {
		cursorChar := " "
		style := choiceStyle
		if idx == s.cursor {
			style = focusedChoiceStyle
			cursorChar = ">"
		}
		b.WriteString(style.Width(choicesWidth).Render(fmt.Sprintf("%s %s", cursorChar, choice)))
		b.WriteString("\n\n")
	}

	headerView := headerStyle.Width(s.width).Render(header)
	choicesView := lipgloss.Place(
		s.width, s.height-lipgloss.Height(headerView),
		lipgloss.Center, 0.8,
		b.String(),
	)

	return lipgloss.JoinVertical(lipgloss.Center, headerView, choicesView)
}
