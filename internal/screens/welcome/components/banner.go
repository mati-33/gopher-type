package components

import (
	"strings"

	"charm.land/lipgloss/v2"
)

type BannerStyles struct {
	Ascii   lipgloss.Style
	Descr   lipgloss.Style
	Version lipgloss.Style
}

type Banner struct {
	Styles          BannerStyles
	GopherTypeAscii string
	descr           string
	version         string
}

func NewBanner(version string) Banner {
	ascii := `
▄▀▀ █▀▄ █▀▄ █▄█ █▀▀ █▀▄   ▀█▀ ▀▄▀ █▀▄ █▀▀
█▄█ █▄█ █▀▀ █ █ ██▄ █▀▄    █   █  █▀▀ ██▄`
	ascii = strings.TrimSpace(ascii)

	return Banner{
		Styles: BannerStyles{
			Ascii:   lipgloss.NewStyle().Padding(1, 4).Background(lipgloss.Color("#303030")),
			Descr:   lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#aaaaaa")),
			Version: lipgloss.NewStyle().MarginBottom(2).Foreground(lipgloss.Color("#aaaaaa")),
		},
		GopherTypeAscii: ascii,
		descr:           "typing practise app for the terminal",
		version:         version,
	}
}

func (b Banner) View() string {
	width := lipgloss.Width(b.GopherTypeAscii)

	return lipgloss.JoinVertical(lipgloss.Center,
		b.Styles.Ascii.Render(b.GopherTypeAscii),
		lipgloss.PlaceHorizontal(width, lipgloss.Right, b.Styles.Version.Render(b.version)),
		b.Styles.Descr.Render(b.descr),
	)
}
