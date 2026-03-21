package typing

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/themes"
)

type InfoStyles struct {
	Text          lipgloss.Style
	Value         lipgloss.Style
	WordCountIcon lipgloss.Style
	ModeIcon      lipgloss.Style
}

func NewInfoStyle(theme themes.Theme) InfoStyles {
	return InfoStyles{
		Text:          lipgloss.NewStyle().Foreground(theme.Text),
		Value:         lipgloss.NewStyle().Foreground(theme.Text),
		WordCountIcon: lipgloss.NewStyle().SetString("").Foreground(theme.Primary),
		ModeIcon:      lipgloss.NewStyle().SetString("󰦨").Foreground(theme.Primary),
	}
}

type Info struct {
	WordCount int
	Mode      string
	Styles    InfoStyles
}

func NewInfo(theme themes.Theme, wordCount int, mode string) Info {
	return Info{
		WordCount: wordCount,
		Mode:      mode,
		Styles:    NewInfoStyle(theme),
	}
}

func (i Info) View() string {
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
