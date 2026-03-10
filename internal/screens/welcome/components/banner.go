package components

import (
	"image/color"
	"strings"

	"charm.land/lipgloss/v2"
)

type BannerStyles struct {
	Ascii                     lipgloss.Style
	Descr                     lipgloss.Style
	Version                   lipgloss.Style
	VersionGradientColorStart color.Color
	VersionGradientColorEnd   color.Color
}

type Banner struct {
	Styles          BannerStyles
	gopherTypeAscii string
	descr           string
	version         string
}

func NewBanner() Banner {
	ascii := `
▄▀▀ █▀▄ █▀▄ █▄█ █▀▀ █▀▄   ▀█▀ ▀▄▀ █▀▄ █▀▀
█▄█ █▄█ █▀▀ █ █ ██▄ █▀▄    █   █  █▀▀ ██▄`
	ascii = strings.TrimSpace(ascii)

	return Banner{
		Styles: BannerStyles{
			Ascii:                     lipgloss.NewStyle().MarginBottom(1),
			Descr:                     lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#aaaaaa")),
			Version:                   lipgloss.NewStyle().MarginBottom(2),
			VersionGradientColorStart: lipgloss.Color("#12b3eb"),
			VersionGradientColorEnd:   lipgloss.Color("#5460f9"),
		},
		gopherTypeAscii: ascii,
		descr:           "typing practise app for the terminal",
		version:         "v0.3.0",
	}
}

func (b Banner) View() string {
	width := lipgloss.Width(b.gopherTypeAscii)
	versionRaw := lipgloss.PlaceHorizontal(width, 0.15, b.version)
	gradient := lipgloss.Blend1D(
		lipgloss.Width(versionRaw),
		b.Styles.VersionGradientColorStart,
		b.Styles.VersionGradientColorEnd,
	)

	version := strings.Builder{}

	for i, ch := range versionRaw {
		version.WriteString(lipgloss.NewStyle().Background(gradient[i]).Render(string(ch)))
	}

	return lipgloss.JoinVertical(lipgloss.Center,
		b.Styles.Ascii.Render(b.gopherTypeAscii),
		b.Styles.Version.Render(version.String()),
		b.Styles.Descr.Render(b.descr),
	)
}
