package screens

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/appcontex"
	comp "github.com/mati-33/gopher-type/internal/components"
	"github.com/mati-33/gopher-type/internal/modes"
)

type modeChange struct {
	ctx      *appcontex.AppContext
	preview  comp.Preview
	picker   comp.Select
	help     comp.Help
	keybinds modeChangeKeybinds
}

func NewModeChange(ctx *appcontex.AppContext) *modeChange {
	picker := comp.NewSelect(ctx.Theme, modes.GetModeNames(), "modes:", ctx.Config.ModeIcon)
	picker.SetSelected(ctx.Mode.Name())
	preview := comp.NewPreview(ctx.Theme, string(ctx.Mode.Generate(ctx.Config.PreviewSize)))
	keybinds := newModeChangeKeybinds()

	return &modeChange{
		ctx:     ctx,
		preview: preview,
		picker:  picker,
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

func (m *modeChange) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case comp.SelectChanged:
		mode := modes.MustGetMode(msg.Option)
		m.preview.Text = string(mode.Generate(m.ctx.Config.PreviewSize))

	case tea.KeyMsg:
		switch msg.String() {

		case m.keybinds.Refresh.Key:
			name := m.picker.Selected()
			mode := modes.MustGetMode(name)
			m.preview.Text = string(mode.Generate(m.ctx.Config.PreviewSize))

		case m.keybinds.Choose.Key:
			name := m.picker.Selected()
			return pop(func() tea.Msg {
				return ChangeProvider{Name: name}
			})

		case m.keybinds.ThemeChange.Key:
			return push(NewThemeChange(m.ctx))

		case m.keybinds.Cancel.Key:
			return pop(nil)

		case m.keybinds.Help.Key:
			m.help.Toggle()
			return nil
		}
	}

	cmds := []tea.Cmd{
		m.picker.Update(msg),
		m.preview.Update(msg),
		m.help.Update(msg),
	}

	return tea.Batch(cmds...)
}

func (m *modeChange) View() string {
	pickerView := m.picker.View()
	previewView := m.preview.View(min(60, int(float32(m.ctx.Width)*0.55)))

	helpView := m.help.View()
	helpOffset := max(0, m.ctx.Width-lipgloss.Width(helpView)-2)

	layer := lipgloss.NewLayer(lipgloss.Place(
		m.ctx.Width, m.ctx.Height,
		lipgloss.Center, 0.6,
		lipgloss.JoinHorizontal(lipgloss.Top,
			pickerView,
			strings.Repeat(" ", 10),
			previewView,
		),
	),
		lipgloss.NewLayer(helpView).Y(m.ctx.Height-lipgloss.Height(helpView)).X(helpOffset),
	)

	c := lipgloss.NewCompositor(layer)

	return c.Render()
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
