package welcome

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/themes"
)

type infoStyles struct {
	text  lipgloss.Style
	value lipgloss.Style
}

func newInfoStyles(theme themes.Theme) infoStyles {
	return infoStyles{
		text:  lipgloss.NewStyle().Foreground(theme.Secondary),
		value: lipgloss.NewStyle().Foreground(theme.TextMuted),
	}
}

type info struct {
	styles    infoStyles
	modeName  string
	themeName string
	width     int
}

func newInfo(theme themes.Theme, modeName, themeName string, width int) info {
	return info{
		styles:    newInfoStyles(theme),
		modeName:  modeName,
		themeName: themeName,
		width:     width,
	}
}

func (i info) View() string {
	ml := i.styles.text.Render("mode")
	mv := i.styles.value.Render(i.modeName)
	m := fmt.Sprintf("%s%s%s", ml, strings.Repeat(" ", i.width-lipgloss.Width(ml)-lipgloss.Width(mv)), mv)
	tl := i.styles.text.Render("theme")
	tv := i.styles.value.Render(i.themeName)
	t := fmt.Sprintf("%s%s%s", tl, strings.Repeat(" ", i.width-lipgloss.Width(tl)-lipgloss.Width(tv)), tv)

	return lipgloss.JoinVertical(lipgloss.Left, m, t)

}
