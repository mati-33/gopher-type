package components

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/themes"
)

type MenuInfoStyles struct {
	text  lipgloss.Style
	value lipgloss.Style
}

func newInfoStyles(theme themes.Theme) MenuInfoStyles {
	return MenuInfoStyles{
		text:  lipgloss.NewStyle().Foreground(theme.Secondary),
		value: lipgloss.NewStyle().Foreground(theme.TextMuted),
	}
}

type MenuInfo struct {
	styles    MenuInfoStyles
	ModeName  string
	ThemeName string
	width     int
}

func NewMenuInfo(theme themes.Theme, modeName, themeName string, width int) MenuInfo {
	return MenuInfo{
		styles:    newInfoStyles(theme),
		ModeName:  modeName,
		ThemeName: themeName,
		width:     width,
	}
}

func (m *MenuInfo) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case themes.Theme:
		m.styles = newInfoStyles(msg)
	}

	return nil
}

func (i *MenuInfo) View() string {
	ml := i.styles.text.Render("mode")
	mv := i.styles.value.Render(i.ModeName)
	m := fmt.Sprintf("%s%s%s", ml, strings.Repeat(" ", i.width-lipgloss.Width(ml)-lipgloss.Width(mv)), mv)
	tl := i.styles.text.Render("theme")
	tv := i.styles.value.Render(i.ThemeName)
	t := fmt.Sprintf("%s%s%s", tl, strings.Repeat(" ", i.width-lipgloss.Width(tl)-lipgloss.Width(tv)), tv)

	return lipgloss.JoinVertical(lipgloss.Left, m, t)
}
