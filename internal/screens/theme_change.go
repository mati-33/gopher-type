package screens

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/appcontex"
	comp "github.com/mati-33/gopher-type/internal/components"
	"github.com/mati-33/gopher-type/internal/themes"
)

type themeChange struct {
	ctx           *appcontex.AppContext
	initialTheme  themes.Theme
	keybinds      themeChangeKeybinds
	picker        comp.Select
	help          comp.Help
	mockSpeed     comp.DetailField
	mockAccuracy  comp.DetailField
	mockMode      comp.DetailField
	mockWordCount comp.DetailField
	mockText      comp.Text
}

func NewThemeChange(ctx *appcontex.AppContext) *themeChange {
	picker := comp.NewSelect(ctx.Theme, themes.GetThemeNames(), "themes:", ctx.Config.ThemeIcon)
	picker.SetSelected(ctx.Theme.Name)
	keybinds := newThemeChangeKeybinds()

	return &themeChange{
		ctx:    ctx,
		picker: picker,
		help: comp.NewHelp(ctx.Theme, []comp.Keybind{
			keybinds.Next,
			keybinds.Previous,
			keybinds.Choose,
			keybinds.Cancel,
			keybinds.Help,
		}),
		keybinds:      keybinds,
		initialTheme:  ctx.Theme,
		mockSpeed:     comp.NewDetailField(ctx.Theme, "speed", ctx.Config.SpeedIcon, "64wpm"),
		mockAccuracy:  comp.NewDetailField(ctx.Theme, "accuracy", ctx.Config.AccuracyIncon, "97.23%"),
		mockMode:      comp.NewDetailField(ctx.Theme, "mode", ctx.Config.ModeIcon, "english"),
		mockWordCount: comp.NewDetailField(ctx.Theme, "word count", ctx.Config.WordCountIcon, "15"),
		mockText:      comp.NewText(ctx.Theme, []rune("hello world foo bar hehe"), 30, 2),
	}
}

func (t *themeChange) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {

	case comp.SelectChanged:
		theme := themes.MustGetTheme(msg.Option)
		return tea.Batch(t.mockText.Update(theme), func() tea.Msg { return theme })

	case tea.KeyMsg:
		switch msg.String() {

		case t.keybinds.ToggleTransparency.Key:
			return func() tea.Msg { return themes.ToggleTransparency{} }

		case t.keybinds.Choose.Key:
			return pop(nil)

		case t.keybinds.Cancel.Key:
			return pop(func() tea.Msg { return t.initialTheme })

		case t.keybinds.Help.Key:
			t.help.Toggle()
			return nil
		}
	}

	cmds := []tea.Cmd{
		t.picker.Update(msg),
		t.help.Update(msg),
		t.mockSpeed.Update(msg),
		t.mockAccuracy.Update(msg),
		t.mockMode.Update(msg),
		t.mockWordCount.Update(msg),
	}

	return tea.Batch(cmds...)
}

func (t *themeChange) View() string {
	helpView := t.help.View()
	helpXOffset := max(0, t.ctx.Width-lipgloss.Width(helpView)-2)
	helpYOffset := max(0, t.ctx.Height-lipgloss.Height(helpView))

	bannerView := lipgloss.JoinHorizontal(lipgloss.Top,
		t.mockSpeed.View(),
		"  ",
		t.mockAccuracy.View(),
		"     ",
		t.mockMode.View(),
		"  ",
		t.mockWordCount.View(),
	)

	mockTextView := t.mockText.View()

	l := lipgloss.NewLayer(lipgloss.Place(
		t.ctx.Width, t.ctx.Height,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Top,
			t.picker.View(),
			strings.Repeat(" ", 10),
			lipgloss.JoinVertical(lipgloss.Center, bannerView, mockTextView),
		),
	),
		lipgloss.NewLayer(helpView).Y(helpYOffset).X(helpXOffset),
	)

	c := lipgloss.NewCompositor(l)

	return c.Render()
}

type themeChangeKeybinds struct {
	Next               comp.Keybind
	Previous           comp.Keybind
	Choose             comp.Keybind
	ToggleTransparency comp.Keybind
	Cancel             comp.Keybind
	Help               comp.Keybind
}

func newThemeChangeKeybinds() themeChangeKeybinds {
	return themeChangeKeybinds{
		Next:               comp.Keybind{Key: "j", Desc: "next"},
		Previous:           comp.Keybind{Key: "k", Desc: "previous"},
		Choose:             comp.Keybind{Key: "enter", Desc: "choose"},
		ToggleTransparency: comp.Keybind{Key: "l", Desc: "toggle transparency"},
		Cancel:             comp.Keybind{Key: "esc", Desc: "cancel"},
		Help:               comp.Keybind{Key: "f1", Desc: "close help"},
	}
}
