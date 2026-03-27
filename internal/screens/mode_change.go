package screens

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/appcontex"
	comp "github.com/mati-33/gopher-type/internal/components"
	"github.com/mati-33/gopher-type/internal/modes"
)

type modeScreen struct {
	ctx      *appcontex.AppContext
	preview  comp.Preview
	choices  comp.Select
	help     comp.Help
	keybinds modeChangeKeybinds
}

func NewModeScreen(ctx *appcontex.AppContext) *modeScreen {
	choices := comp.NewSelect(ctx.Theme, modes.GetModeNames(), "modes:", ctx.Config.ModeIcon)
	mode := modes.MustGetMode(choices.Selected())
	preview := comp.NewPreview(ctx.Theme, int(float32(ctx.Width)*0.55), string(mode.Generate(ctx.Config.PreviewSize)))
	keybinds := newModeChangeKeybinds()

	return &modeScreen{
		ctx:     ctx,
		preview: preview,
		choices: choices,
		help: comp.NewHelp(ctx.Theme, []comp.Keybind{
			keybinds.Next,
			keybinds.Previous,
			keybinds.Refresh,
			keybinds.Choose,
			keybinds.ThemeChange,
			keybinds.Cancel,
			keybinds.Help,
		}),
		keybinds: keybinds,
	}
}

func (m *modeScreen) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case comp.ChoiceChanged:
		mode := modes.MustGetMode(msg.Name)
		m.preview.Text = string(mode.Generate(m.ctx.Config.PreviewSize))

	case tea.WindowSizeMsg:
		m.preview.Width = int(float32(msg.Width) * 0.55)

	case tea.KeyMsg:
		switch msg.String() {

		case m.keybinds.Refresh.Key:
			name := m.choices.Selected()
			mode := modes.MustGetMode(name)
			m.preview.Text = string(mode.Generate(m.ctx.Config.PreviewSize))

		case m.keybinds.Choose.Key:
			name := m.choices.Selected()
			return popScreen(func() tea.Msg {
				return ChangeProvider{Name: name}
			})

		case m.keybinds.ThemeChange.Key:
			return pushScreen(NewThemeChangeScreen(m.ctx))

		case m.keybinds.Cancel.Key:
			return popScreen(nil)

		case m.keybinds.Help.Key:
			m.help.Toggle()
			return nil
		}
	}

	cmds := []tea.Cmd{
		m.choices.Update(msg),
		m.preview.Update(msg),
		m.help.Update(msg),
	}

	return tea.Batch(cmds...)
}

func (m *modeScreen) View() tea.View {
	choicesView := m.choices.View()
	previewView := m.preview.View()
	helpView := m.help.View()
	helpOffset := max(0, m.ctx.Width-lipgloss.Width(helpView)-2)

	layer := lipgloss.NewLayer(lipgloss.Place(
		m.ctx.Width, m.ctx.Height,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Top, choicesView, "     ", previewView),
	),
		lipgloss.NewLayer(helpView).Y(m.ctx.Height-lipgloss.Height(helpView)).X(helpOffset),
	)

	c := lipgloss.NewCompositor(layer)
	return tea.NewView(c.Render())
}

type modeChangeKeybinds struct {
	Next        comp.Keybind
	Previous    comp.Keybind
	Refresh     comp.Keybind
	Choose      comp.Keybind
	ThemeChange comp.Keybind
	Cancel      comp.Keybind
	Help        comp.Keybind
}

func newModeChangeKeybinds() modeChangeKeybinds {
	return modeChangeKeybinds{
		Next:        comp.Keybind{Key: "j", Desc: "next"},
		Previous:    comp.Keybind{Key: "k", Desc: "previous"},
		Refresh:     comp.Keybind{Key: "r", Desc: "refresh"},
		Choose:      comp.Keybind{Key: "enter", Desc: "choose"},
		ThemeChange: comp.Keybind{Key: "ctrl+t", Desc: "change theme"},
		Cancel:      comp.Keybind{Key: "esc", Desc: "cancel"},
		Help:        comp.Keybind{Key: "f1", Desc: "close help"},
	}
}
