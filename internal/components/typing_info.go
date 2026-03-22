package components

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/themes"
)

type TypingInfoStyles struct {
	Text          lipgloss.Style
	Value         lipgloss.Style
	WordCountIcon lipgloss.Style
	ModeIcon      lipgloss.Style
}

func NewTypingInfoStyle(theme themes.Theme) TypingInfoStyles {
	return TypingInfoStyles{
		Text:          lipgloss.NewStyle().Foreground(theme.Text),
		Value:         lipgloss.NewStyle().Foreground(theme.Text),
		WordCountIcon: lipgloss.NewStyle().SetString("").Foreground(theme.Primary),
		ModeIcon:      lipgloss.NewStyle().SetString("󰦨").Foreground(theme.Primary),
	}
}

type TypingInfo struct {
	WordCount int
	Mode      string
	Styles    TypingInfoStyles
}

func NewTypingInfo(theme themes.Theme, wordCount int, mode string) TypingInfo {
	return TypingInfo{
		WordCount: wordCount,
		Mode:      mode,
		Styles:    NewTypingInfoStyle(theme),
	}
}

func (i TypingInfo) View() string {
	count := lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("%s %s", i.Styles.WordCountIcon.Render(), i.Styles.Text.Render("word count")),
		fmt.Sprintf("  %s", i.Styles.Value.Render(fmt.Sprintf("%d", i.WordCount))),
	)
	mode := lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("%s %s", i.Styles.ModeIcon.Render(), i.Styles.Text.Render("mode")),
		fmt.Sprintf("  %s", i.Styles.Value.Render(i.Mode)),
	)

	return lipgloss.JoinHorizontal(lipgloss.Top, mode, "   ", count)
}
