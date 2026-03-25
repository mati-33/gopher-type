package screens

import (
	tea "charm.land/bubbletea/v2"
	comp "github.com/mati-33/gopher-type/internal/components"
	"github.com/mati-33/gopher-type/internal/config"
	"github.com/mati-33/gopher-type/internal/themes"
)

type themeChangeScreen struct {
	config  config.Config
	theme   themes.Theme
	choices comp.Select
}

func NewThemeChangeScreen(config config.Config, theme themes.Theme) themeChangeScreen {
	return themeChangeScreen{
		config:  config,
		theme:   theme,
		choices: comp.NewSelect(theme, themes.GetThemeNames(), "themes:", "A"),
	}
}

func (t themeChangeScreen) Init() tea.Cmd {
	return nil
}

func (t themeChangeScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case comp.ChoiceChanged:
		return t, func() tea.Msg { return themes.MustGetTheme(msg.Name) }
	case tea.KeyMsg:
		switch msg.String() {
		case "l":
			return t, func() tea.Msg { return themes.ToggleTransparency{} }
		case "enter":
			return t, func() tea.Msg { return PopScreen{} }
		}
	}

	cmd := t.choices.Update(msg)
	return t, cmd
}

func (t themeChangeScreen) View() tea.View {
	return tea.NewView(t.choices.View())
}
