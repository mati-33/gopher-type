package mode

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/mati-33/gopher-type/internal/screens"
	"github.com/mati-33/gopher-type/internal/text_providers"
)

type modeScreen struct {
	width  int
	height int
	modes  []string
	cursor int
}

func NewModeScreen(width, height int) modeScreen {
	return modeScreen{
		width:  width,
		height: height,
		modes:  textproviders.GetProviderNames(),
		cursor: 0,
	}
}

func (m modeScreen) Init() tea.Cmd {
	return nil
}

func (m modeScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			m.cursor++
		case "k":
			m.cursor--
		case "enter":
			return m, func() tea.Msg {
				return screens.PopScreen{
					Command: func() tea.Msg {
						return screens.ChangeProvider{
							Name: m.modes[m.cursor],
						}
					},
				}
			}
		case "esc":
			return m, func() tea.Msg {
				return screens.PopScreen{}
			}
		}
	}
	return m, nil
}

func (m modeScreen) View() tea.View {
	b := strings.Builder{}

	for i, mode := range m.modes {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		fmt.Fprintf(&b, "%s %s\n", cursor, mode)
	}

	return tea.NewView(b.String())
}
