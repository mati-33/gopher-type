package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	f, err := tea.LogToFile("debug.log", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup debug logging: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	m := newModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start the program: %v", err)
		os.Exit(1)
	}
}

type model struct {
	screenStack []tea.Model
	width       int
	height      int
}

func newModel() model {
	return model{
		screenStack: []tea.Model{},
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

		if len(m.screenStack) == 0 {
			return m, func() tea.Msg {
				return PushScreen{
					newWelcomeScreen(m.width, m.height),
				}
			}
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

	case PushScreen:
		m.screenStack = append(m.screenStack, msg.screen)
		return m, msg.screen.Init()

	case PopScreen:
		m.screenStack = m.screenStack[:len(m.screenStack)-1]
		return m, tea.ClearScreen
	}

	if len(m.screenStack) == 0 {
		return m, nil
	}

	currentScreen := m.screenStack[len(m.screenStack)-1]
	currentScreen, cmd := currentScreen.Update(msg)
	m.screenStack[len(m.screenStack)-1] = currentScreen

	return m, cmd
}

func (m model) View() tea.View {
	if len(m.screenStack) > 0 {
		view := m.screenStack[len(m.screenStack)-1].View()
		view.AltScreen = true
		return view
	}

	return tea.NewView("")
}

type PushScreen struct {
	screen tea.Model
}

type PopScreen struct{}
