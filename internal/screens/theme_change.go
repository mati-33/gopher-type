package screens

import (
	tea "charm.land/bubbletea/v2"
	"github.com/mati-33/gopher-type/internal/appcontex"
	comp "github.com/mati-33/gopher-type/internal/components"
	"github.com/mati-33/gopher-type/internal/themes"
)

type themeChange struct {
	ctx    *appcontex.AppContext
	picker comp.Select
}

func NewThemeChange(ctx *appcontex.AppContext) *themeChange {
	picker := comp.NewSelect(ctx.Theme, themes.GetThemeNames(), "themes:", "A")
	picker.SetSelected(ctx.Theme.Name)

	return &themeChange{
		ctx:    ctx,
		picker: picker,
	}
}

func (t *themeChange) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case comp.SelectChanged:
		return func() tea.Msg { return themes.MustGetTheme(msg.Option) }
	case tea.KeyMsg:
		switch msg.String() {
		case "l":
			return func() tea.Msg { return themes.ToggleTransparency{} }
		case "enter":
			return func() tea.Msg { return Pop{} }
		}
	}

	cmd := t.picker.Update(msg)
	return cmd
}

func (t *themeChange) View() tea.View {
	return tea.NewView(t.picker.View())
}
