package components

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type ChoicesStyles struct {
	Normal    lipgloss.Style
	Selected  lipgloss.Style
	Cursor    lipgloss.Style
	Title     lipgloss.Style
	TitleIcon lipgloss.Style
}

type Choices struct {
	Styles  ChoicesStyles
	choices []string
	cursor  int
}

func NewChoices(choices []string) Choices {
	return Choices{
		Styles: ChoicesStyles{
			Normal:    lipgloss.NewStyle().Foreground(lipgloss.Color("#aaaaaa")),
			Selected:  lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")),
			Cursor:    lipgloss.NewStyle().SetString(">").PaddingRight(1).Foreground(lipgloss.Color("#0087ff")),
			Title:     lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).MarginBottom(1),
			TitleIcon: lipgloss.NewStyle().SetString("󰦨").PaddingRight(1).Foreground(lipgloss.Color("#d7005f")),
		},
		choices: choices,
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
		fmt.Sprintf("%s%s", c.Styles.TitleIcon.Render(), c.Styles.Title.Render("modes:")),
		b.String(),
	)
}

func (c Choices) Selected() string {
	return c.choices[c.cursor]
}

type ChoiceChanged struct {
	Name string
}
