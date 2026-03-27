package components

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/themes"
)

type SelectStyles struct {
	Normal   lipgloss.Style
	Selected lipgloss.Style
	Cursor   lipgloss.Style
	Label    lipgloss.Style
	Icon     lipgloss.Style
}

func newSelectStyles(theme themes.Theme) SelectStyles {
	return SelectStyles{
		Normal:   lipgloss.NewStyle().Foreground(theme.TextMuted),
		Selected: lipgloss.NewStyle().Foreground(theme.Text),
		Cursor:   lipgloss.NewStyle().SetString(">").PaddingRight(1).Foreground(theme.Secondary).Bold(true),
		Label:    lipgloss.NewStyle().Foreground(theme.Text).MarginBottom(1),
		Icon:     lipgloss.NewStyle().PaddingRight(1).Foreground(theme.Primary),
	}

}

type Select struct {
	Styles  SelectStyles
	options []string
	cursor  int
	label   string
	icon    string
}

func NewSelect(theme themes.Theme, options []string, label, icon string) Select {
	return Select{
		Styles:  newSelectStyles(theme),
		options: options,
		label:   label,
		icon:    icon,
	}
}

func (s *Select) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {

	case themes.Theme:
		s.Styles = newSelectStyles(msg)
		return nil

	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			if s.cursor < len(s.options)-1 {
				s.cursor++
				return func() tea.Msg { return SelectChanged{s.options[s.cursor]} }
			}
		case "k":
			if s.cursor > 0 {
				s.cursor--
				return func() tea.Msg { return SelectChanged{s.options[s.cursor]} }
			}
		}
	}

	return nil
}

func (s *Select) View() string {
	b := strings.Builder{}
	cursorWidth := lipgloss.Width(s.Styles.Cursor.Render())

	for i, o := range s.options {
		cursor := strings.Repeat(" ", cursorWidth)
		style := s.Styles.Normal
		if i == s.cursor {
			cursor = s.Styles.Cursor.Render()
			style = s.Styles.Selected
		}
		fmt.Fprintf(&b, "%s%s\n", cursor, style.Render(o))
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("%s%s", s.Styles.Icon.Render(s.icon), s.Styles.Label.Render(s.label)),
		b.String(),
	)
}

func (s *Select) Selected() string {
	return s.options[s.cursor]
}

type SelectChanged struct {
	Option string
}
