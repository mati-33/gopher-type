package components

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/themes"
)

type DetailFieldStyles struct {
	Label lipgloss.Style
	Icon  lipgloss.Style
	Value lipgloss.Style
}

func newDetailFieldStyles(theme themes.Theme) DetailFieldStyles {
	return DetailFieldStyles{
		Label: lipgloss.NewStyle().Foreground(theme.Text),
		Icon:  lipgloss.NewStyle().Foreground(theme.Primary),
		Value: lipgloss.NewStyle().Foreground(theme.Text),
	}
}

type DetailField struct {
	Styles DetailFieldStyles
	Label  string
	Icon   string
	Value  string
}

func NewDetailField(theme themes.Theme, label, icon, value string) DetailField {
	return DetailField{
		Styles: newDetailFieldStyles(theme),
		Label:  label,
		Icon:   icon,
		Value:  value,
	}
}

func (df DetailField) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("%s %s", df.Styles.Icon.Render(df.Icon), df.Styles.Label.Render(df.Label)),
		fmt.Sprintf("  %s", df.Styles.Value.Render(df.Value)),
	)
}
