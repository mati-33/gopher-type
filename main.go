package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/mati-33/gopher-type/internal/ctx"
	"github.com/mati-33/gopher-type/internal/modes"
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
	ctx         *ctx.Context
	screenStack []screens.Interface
}

func newModel() model {
	return model{
		ctx:         ctx.New(),
		screenStack: []screens.Interface{},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.ctx.Width = msg.Width
		m.ctx.Height = msg.Height

		if len(m.screenStack) > 0 {
			return m, nil
		}

		if len(m.screenStack) == 0 {
			return m, func() tea.Msg {
				sc := screens.NewWelcomeScreen(m.ctx)
				return screens.PushScreen{
					Screen: &sc,
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
		return m, nil

	case screens.PopScreen:
		m.screenStack = m.screenStack[:len(m.screenStack)-1]
		return m, msg.Command

	case screens.ChangeProvider:
		m.ctx.Mode = modes.MustGetMode(msg.Name)
		return m, tea.Batch(m.passToAll(msg)...)

	case themes.Theme:
		m.ctx.Theme = msg
		return m, tea.Batch(m.passToAll(msg)...)

	case themes.ToggleTransparency:
		m.ctx.Config.Transparent = !m.ctx.Config.Transparent
		return m, nil
	}

	if len(m.screenStack) == 0 {
		return m, nil
	}

	currentScreen := m.screenStack[len(m.screenStack)-1]
	cmd := currentScreen.Update(msg)

	return m, cmd
}

func (m model) View() tea.View {
	if len(m.screenStack) > 0 {
		view := m.screenStack[len(m.screenStack)-1].View()
		view.AltScreen = true
		if !m.ctx.Config.Transparent {
			view.BackgroundColor = m.ctx.Theme.Background
		}
		return view
	}

	return tea.NewView("")
}

func (m model) passToAll(msg tea.Msg) []tea.Cmd {
	cmds := make([]tea.Cmd, 0, len(m.screenStack))

	for _, s := range m.screenStack {
		cmds = append(cmds, s.Update(msg))
	}

	return cmds
}
