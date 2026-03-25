package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/mati-33/gopher-type/internal/config"
	"github.com/mati-33/gopher-type/internal/screens"
	"github.com/mati-33/gopher-type/internal/themes"
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
	config      config.Config
	theme       themes.Theme
}

func newModel() model {
	config := config.NewDefault()
	return model{
		screenStack: []tea.Model{},
		config:      config,
		theme:       themes.MustGetTheme(config.InitTheme),
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
				return screens.PushScreen{
					Screen: screens.NewWelcomeScreen(m.config, m.theme, m.width, m.height),
				}
			}
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

	case screens.PushScreen:
		m.screenStack = append(m.screenStack, msg.Screen)
		return m, msg.Screen.Init()

	case screens.PopScreen:
		m.screenStack = m.screenStack[:len(m.screenStack)-1]
		return m, msg.Command

	case themes.Theme:
		m.theme = msg
		screens := make([]tea.Model, 0, len(m.screenStack))
		cmds := make([]tea.Cmd, 0, len(m.screenStack))
		for _, s := range m.screenStack {
			ns, cmd := s.Update(msg)
			screens = append(screens, ns)
			cmds = append(cmds, cmd)
		}
		m.screenStack = screens
		return m, tea.Batch(cmds...)

	case themes.ToggleTransparency:
		m.config.Transparent = !m.config.Transparent
		return m, nil
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
		if !m.config.Transparent {
			view.BackgroundColor = m.theme.Background
		}
		return view
	}

	return tea.NewView("")
}
