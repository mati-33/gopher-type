package mode

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

type PreviewStyles struct {
	Title     lipgloss.Style
	Text      lipgloss.Style
	TitleIcon lipgloss.Style
}

type Preview struct {
	Styles PreviewStyles
	Text   string
	Width  int
}

func NewPreview(width int, text string) Preview {
	return Preview{
		Styles: PreviewStyles{
			Title:     lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).MarginBottom(1),
			Text:      lipgloss.NewStyle().Align(lipgloss.Left).Foreground(lipgloss.Color("#aaaaaa")).MarginLeft(2),
			TitleIcon: lipgloss.NewStyle().SetString("󱎸").PaddingRight(1).Foreground(lipgloss.Color("#5f00ff")),
		},
		Width: width,
		Text:  text,
	}
}

func (p Preview) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("%s%s", p.Styles.TitleIcon.Render(), p.Styles.Title.Render("preview:")),
		p.Styles.Text.Width(p.Width).Render(p.Text),
	)
}
