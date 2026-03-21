package welcome

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/screens"
)

type MenuStyles struct {
	Key         lipgloss.Style
	Description lipgloss.Style
}

type Menu struct {
	Styles  MenuStyles
	Options []screens.Keybind
	Width   int
}

func NewMenu(options []screens.Keybind, width int) Menu {
	return Menu{
		Styles: MenuStyles{
			Key:         lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")),
			Description: lipgloss.NewStyle().Foreground(lipgloss.Color("#aaaaaa")),
		},
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
