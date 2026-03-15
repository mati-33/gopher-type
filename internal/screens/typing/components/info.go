package components

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

type InfoStyles struct {
	Text          lipgloss.Style
	Value         lipgloss.Style
	WordCountIcon lipgloss.Style
	ModeIcon      lipgloss.Style
}

func NewInfoStyle() InfoStyles {
	grey := lipgloss.Color("#bbbbbb")
	white := lipgloss.Color("#ffffff")
	violet := lipgloss.Color("#5f00ff")
	pink := lipgloss.Color("#d7005f")

	return InfoStyles{
		Text:          lipgloss.NewStyle().Foreground(grey),
		Value:         lipgloss.NewStyle().Foreground(white),
		WordCountIcon: lipgloss.NewStyle().SetString("").Foreground(violet),
		ModeIcon:      lipgloss.NewStyle().SetString("󰦨").Foreground(pink),
	}
}

type Info struct {
	WordCount int
	Mode      string
	Styles    InfoStyles
}

func NewInfo(wordCount int, mode string) Info {
	return Info{
		WordCount: wordCount,
		Mode:      mode,
		Styles:    NewInfoStyle(),
	}
}

func (i Info) View() string {
	count := lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("%s %s", i.Styles.WordCountIcon.Render(), i.Styles.Text.Render("word count")),
		fmt.Sprintf("  %d", i.WordCount),
	)
	mode := lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("%s %s", i.Styles.ModeIcon.Render(), i.Styles.Text.Render("mode")),
		fmt.Sprintf("  %s", i.Mode),
	)

	return lipgloss.JoinHorizontal(lipgloss.Top, mode, "   ", count)
}
