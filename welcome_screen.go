package main

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	textproviders "github.com/mati-33/gopher-type/text_providers"
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
	headerStyle = lipgloss.NewStyle().Align(lipgloss.Center).MarginTop(1).MarginBottom(3)
)

type welcomeScreen struct {
	width  int
	height int
}

func newWelcomeScreen(width, height int) welcomeScreen {
	return welcomeScreen{
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

		case "enter":
			return s, func() tea.Msg {
				p := textproviders.NewWordArrayProviderFromTxtFile(textproviders.Eng1k)
				return PushScreen{
					screen: newTypingScreen(
						p,
						150,
						s.width,
						s.height,
					),
				}
			}
		}
	}
	return s, nil
}

func (s welcomeScreen) View() tea.View {
	headerView := headerStyle.Width(s.width).Render(header)
	welcomeTextView := "Press ENTER to start"

	return tea.NewView(lipgloss.JoinVertical(lipgloss.Center, headerView, welcomeTextView))
}
