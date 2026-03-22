package components

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/themes"
)

type MenuStyles struct {
	Key         lipgloss.Style
	Description lipgloss.Style
}

func newMenuStyles(theme themes.Theme) MenuStyles {
	return MenuStyles{
		Key:         lipgloss.NewStyle().Foreground(theme.Primary).Bold(true),
		Description: lipgloss.NewStyle().Foreground(theme.Text),
	}
}

type Menu struct {
	Styles  MenuStyles
	Options []Keybind
	Width   int
}

func NewMenu(theme themes.Theme, options []Keybind, width int) Menu {
	return Menu{
		Styles:  newMenuStyles(theme),
		Options: options,
		Width:   width,
	}
}

func (m Menu) View() string {
	b := strings.Builder{}

	for _, opt := range m.Options {
		k := m.Styles.Key.Render(opt.Key)
		d := m.Styles.Description.Render(opt.Desc)
		spaces := max(1, m.Width-lipgloss.Width(k)-lipgloss.Width(d))
		fmt.Fprintf(&b, "%s%s%s\n", k, strings.Repeat(" ", spaces), d)
	}

	return b.String()
}
