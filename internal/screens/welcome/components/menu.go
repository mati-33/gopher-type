package components

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

type MenuOption struct {
	Key         string
	Description string
}

type MenuStyles struct {
	Key         lipgloss.Style
	Description lipgloss.Style
}

type Menu struct {
	Styles  MenuStyles
	Options []MenuOption
	Width   int
}

func NewMenu(options []MenuOption, width int) Menu {
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
		d := m.Styles.Description.Render(opt.Description)
		spaces := max(1, m.Width-lipgloss.Width(k)-lipgloss.Width(d))
		fmt.Fprintf(&b, "%s%s%s\n", k, strings.Repeat(" ", spaces), d)
	}

	return b.String()
}
