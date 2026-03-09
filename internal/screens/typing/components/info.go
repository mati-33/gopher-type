package components

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

type InfoStyles struct {
	Text          lipgloss.Style
	Value         lipgloss.Style
	WordCountIcon lipgloss.Style
}

func NewInfoStyle() InfoStyles {
	grey := lipgloss.Color("#bbbbbb")
	white := lipgloss.Color("#ffffff")
	violet := lipgloss.Color("#5f00ff")

	return InfoStyles{
		Text:          lipgloss.NewStyle().Foreground(grey),
		Value:         lipgloss.NewStyle().Foreground(white),
		WordCountIcon: lipgloss.NewStyle().SetString("").Foreground(violet),
	}
}

type Info struct {
	WordCount int
	Styles    InfoStyles
}

func NewInfo(wordCount int) Info {
	return Info{
		WordCount: wordCount,
		Styles:    NewInfoStyle(),
	}
}

func (i Info) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("%s %s", i.Styles.WordCountIcon.Render(), i.Styles.Text.Render("word count")),
		fmt.Sprintf("  %d", i.WordCount),
	)
}
