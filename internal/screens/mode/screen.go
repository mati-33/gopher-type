package mode

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/screens"
	"github.com/mati-33/gopher-type/internal/screens/mode/components"
	"github.com/mati-33/gopher-type/internal/text_providers"
)

type modeScreen struct {
	width   int
	height  int
	preview components.Preview
	choices components.Choices
}

func NewModeScreen(width, height int) modeScreen {
	choices := components.NewChoices(textproviders.GetProviderNames())
	provider := textproviders.MustGetProvider(choices.Selected())
	preview := components.NewPreview(int(float32(width)*0.55), provider.Preview())

	return modeScreen{
		width:   width,
		height:  height,
		preview: preview,
		choices: choices,
	}
}

func (m modeScreen) Init() tea.Cmd {
	return nil
}

func (m modeScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case components.ChoiceChanged:
		p := textproviders.MustGetProvider(msg.Name)
		m.preview.Text = p.Preview()

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.preview.Width = int(float32(msg.Width) * 0.55)

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			name := m.choices.Selected()
			return m, func() tea.Msg {
				return screens.PopScreen{
					Command: func() tea.Msg {
						return screens.ChangeProvider{
							Name: name,
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

	choices, cmd := m.choices.Update(msg)
	m.choices = choices

	return m, cmd
}

func (m modeScreen) View() tea.View {
	choicesView := m.choices.View()
	previewView := m.preview.View()

	return tea.NewView(lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Top, choicesView, "     ", previewView),
	))
}
