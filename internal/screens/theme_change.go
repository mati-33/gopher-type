package screens

import (
	tea "charm.land/bubbletea/v2"
	"github.com/mati-33/gopher-type/internal/appcontex"
	comp "github.com/mati-33/gopher-type/internal/components"
	"github.com/mati-33/gopher-type/internal/themes"
)

type themeChange struct {
	ctx     *appcontex.AppContext
	choices comp.Select
}

func NewThemeChange(ctx *appcontex.AppContext) *themeChange {
	return &themeChange{
		ctx:     ctx,
		choices: comp.NewSelect(ctx.Theme, themes.GetThemeNames(), "themes:", "A"),
	}
}

func (t *themeChange) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case comp.ChoiceChanged:
		return func() tea.Msg { return themes.MustGetTheme(msg.Name) }
	case tea.KeyMsg:
		switch msg.String() {
		case "l":
			return func() tea.Msg { return themes.ToggleTransparency{} }
		case "enter":
			return func() tea.Msg { return Pop{} }
		}
	}

	cmd := t.choices.Update(msg)
	return cmd
}

func (t *themeChange) View() tea.View {
	return tea.NewView(t.choices.View())
}
