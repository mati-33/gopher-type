package components

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/themes"
)

type ChoicesStyles struct {
	Normal   lipgloss.Style
	Selected lipgloss.Style
	Cursor   lipgloss.Style
	Label    lipgloss.Style
	Icon     lipgloss.Style
}

func newChoicesStyles(theme themes.Theme) ChoicesStyles {
	return ChoicesStyles{
		Normal:   lipgloss.NewStyle().Foreground(theme.TextMuted),
		Selected: lipgloss.NewStyle().Foreground(theme.Text),
		Cursor:   lipgloss.NewStyle().SetString(">").PaddingRight(1).Foreground(theme.Secondary).Bold(true),
		Label:    lipgloss.NewStyle().Foreground(theme.Text).MarginBottom(1),
		Icon:     lipgloss.NewStyle().PaddingRight(1).Foreground(theme.Primary),
	}

}

type Choices struct {
	Styles  ChoicesStyles
	choices []string
	cursor  int
	label   string
	icon    string
}

func NewChoices(theme themes.Theme, choices []string, label, icon string) Choices {
	return Choices{
		Styles:  newChoicesStyles(theme),
		choices: choices,
		label:   label,
		icon:    icon,
	}
}

func (c Choices) Update(msg tea.Msg) (Choices, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			if c.cursor < len(c.choices)-1 {
				c.cursor++
				return c, func() tea.Msg { return ChoiceChanged{c.choices[c.cursor]} }
			}
		case "k":
			if c.cursor > 0 {
				c.cursor--
				return c, func() tea.Msg { return ChoiceChanged{c.choices[c.cursor]} }
			}
		}
	}
	return c, nil
}

func (c Choices) View() string {
	b := strings.Builder{}
	cursorWidth := lipgloss.Width(c.Styles.Cursor.Render())

	for i, choice := range c.choices {
		cursor := strings.Repeat(" ", cursorWidth)
		style := c.Styles.Normal
		if i == c.cursor {
			cursor = c.Styles.Cursor.Render()
			style = c.Styles.Selected
		}
		fmt.Fprintf(&b, "%s%s\n", cursor, style.Render(choice))
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("%s%s", c.Styles.Icon.Render(c.icon), c.Styles.Label.Render(c.label)),
		b.String(),
	)
}

func (c Choices) Selected() string {
	return c.choices[c.cursor]
}

type ChoiceChanged struct {
	Name string
}
