package components

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/themes"
)

type PreviewStyles struct {
	Title     lipgloss.Style
	Text      lipgloss.Style
	TitleIcon lipgloss.Style
}

func newPreviewStyles(theme themes.Theme) PreviewStyles {
	return PreviewStyles{
		Title:     lipgloss.NewStyle().Foreground(theme.Text).MarginBottom(1),
		Text:      lipgloss.NewStyle().Align(lipgloss.Left).Foreground(theme.TextMuted).MarginLeft(2),
		TitleIcon: lipgloss.NewStyle().SetString("󱎸").PaddingRight(1).Foreground(theme.Primary),
	}
}

type Preview struct {
	Styles PreviewStyles
	Text   string
	Width  int
}

func NewPreview(theme themes.Theme, width int, text string) Preview {
	return Preview{
		Styles: newPreviewStyles(theme),
		Width:  width,
		Text:   text,
	}
}

func (p Preview) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("%s%s", p.Styles.TitleIcon.Render(), p.Styles.Title.Render("preview:")),
		p.Styles.Text.Width(p.Width).Render(p.Text),
	)
}
