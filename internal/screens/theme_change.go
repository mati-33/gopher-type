package screens

import (
	tea "charm.land/bubbletea/v2"
	comp "github.com/mati-33/gopher-type/internal/components"
	"github.com/mati-33/gopher-type/internal/ctx"
	"github.com/mati-33/gopher-type/internal/themes"
)

type themeChangeScreen struct {
	ctx     *ctx.Context
	choices comp.Select
}

func NewThemeChangeScreen(ctx *ctx.Context) *themeChangeScreen {
	return &themeChangeScreen{
		ctx:     ctx,
		choices: comp.NewSelect(ctx.Theme, themes.GetThemeNames(), "themes:", "A"),
	}
}

func (t *themeChangeScreen) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case comp.ChoiceChanged:
		return func() tea.Msg { return themes.MustGetTheme(msg.Name) }
	case tea.KeyMsg:
		switch msg.String() {
		case "l":
			return func() tea.Msg { return themes.ToggleTransparency{} }
		case "enter":
			return func() tea.Msg { return PopScreen{} }
		}
	}

	cmd := t.choices.Update(msg)
	return cmd
}

func (t *themeChangeScreen) View() tea.View {
	return tea.NewView(t.choices.View())
}
