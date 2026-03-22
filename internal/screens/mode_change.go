package screens

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	comp "github.com/mati-33/gopher-type/internal/components"
	"github.com/mati-33/gopher-type/internal/config"
	"github.com/mati-33/gopher-type/internal/modes"
	"github.com/mati-33/gopher-type/internal/themes"
)

type modeScreen struct {
	config   config.Config
	width    int
	height   int
	preview  comp.Preview
	choices  comp.Select
	help     comp.Help
	keybinds modeChangeKeybinds
}

func NewModeScreen(config config.Config, theme themes.Theme, width, height int) modeScreen {
	choices := comp.NewSelect(theme, modes.GetModeNames(), "modes:", config.ModeIcon)
	mode := modes.MustGetMode(choices.Selected())
	preview := comp.NewPreview(theme, int(float32(width)*0.55), string(mode.Generate(config.PreviewSize)))
	keybinds := newModeChangeKeybinds()

	return modeScreen{
		config:   config,
		width:    width,
		height:   height,
		preview:  preview,
		choices:  choices,
		keybinds: keybinds,
		help: comp.NewHelp(theme, []comp.Keybind{
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
	case comp.ChoiceChanged:
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
				return PopScreen{
					Command: func() tea.Msg {
						return ChangeProvider{Name: name}
					},
				}
			}

		case m.keybinds.Cancel.Key:
			return m, func() tea.Msg {
				return PopScreen{}
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

type modeChangeKeybinds struct {
	Next     comp.Keybind
	Previous comp.Keybind
	Refresh  comp.Keybind
	Choose   comp.Keybind
	Cancel   comp.Keybind
	Help     comp.Keybind
}

func newModeChangeKeybinds() modeChangeKeybinds {
	return modeChangeKeybinds{
		Next:     comp.Keybind{Key: "j", Desc: "next"},
		Previous: comp.Keybind{Key: "k", Desc: "previous"},
		Refresh:  comp.Keybind{Key: "r", Desc: "refresh"},
		Choose:   comp.Keybind{Key: "enter", Desc: "choose"},
		Cancel:   comp.Keybind{Key: "esc", Desc: "cancel"},
		Help:     comp.Keybind{Key: "f1", Desc: "close help"},
	}
}
