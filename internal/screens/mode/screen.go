package mode

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/config"
	"github.com/mati-33/gopher-type/internal/modes"
	"github.com/mati-33/gopher-type/internal/screens"
	"github.com/mati-33/gopher-type/internal/themes"
)

type modeScreen struct {
	config   config.Config
	width    int
	height   int
	preview  Preview
	choices  Choices
	help     screens.Help
	keybinds keybinds
}

func NewModeScreen(config config.Config, theme themes.Theme, width, height int) modeScreen {
	choices := NewChoices(theme, modes.GetModeNames())
	mode := modes.MustGetMode(choices.Selected())
	preview := NewPreview(theme, int(float32(width)*0.55), string(mode.Generate(config.PreviewSize)))
	keybinds := newKeybinds()

	return modeScreen{
		config:   config,
		width:    width,
		height:   height,
		preview:  preview,
		choices:  choices,
		keybinds: keybinds,
		help: screens.NewHelp(theme, []screens.Keybind{
			keybinds.Next,
			keybinds.Previous,
			keybinds.Refresh,
			keybinds.Choose,
			keybinds.Cancel,
			keybinds.Help,
		}),
	}
}

func (m modeScreen) Init() tea.Cmd {
	return nil
}

func (m modeScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ChoiceChanged:
		mode := modes.MustGetMode(msg.Name)
		m.preview.Text = string(mode.Generate(m.config.PreviewSize))

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.preview.Width = int(float32(msg.Width) * 0.55)

	case tea.KeyMsg:
		switch msg.String() {

		case m.keybinds.Refresh.Key:
			name := m.choices.Selected()
			mode := modes.MustGetMode(name)
			m.preview.Text = string(mode.Generate(m.config.PreviewSize))

		case m.keybinds.Choose.Key:
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

		case m.keybinds.Cancel.Key:
			return m, func() tea.Msg {
				return screens.PopScreen{}
			}

		case m.keybinds.Help.Key:
			m.help.Toggle()
			return m, nil
		}
	}

	choices, cmd := m.choices.Update(msg)
	m.choices = choices

	return m, cmd
}

func (m modeScreen) View() tea.View {
	choicesView := m.choices.View()
	previewView := m.preview.View()
	helpView := m.help.View()
	helpOffset := max(0, m.width-lipgloss.Width(helpView)-2)

	layer := lipgloss.NewLayer(lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Top, choicesView, "     ", previewView),
	),
		lipgloss.NewLayer(helpView).Y(m.height-lipgloss.Height(helpView)).X(helpOffset),
	)

	c := lipgloss.NewCompositor(layer)
	return tea.NewView(c.Render())
}
