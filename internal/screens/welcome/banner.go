package welcome

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/themes"
)

type BannerStyles struct {
	Ascii   lipgloss.Style
	Descr   lipgloss.Style
	Version lipgloss.Style
}

func newBannerStyles(theme themes.Theme) BannerStyles {
	return BannerStyles{
		Ascii:   lipgloss.NewStyle().Padding(1, 4).Foreground(theme.Text),
		Descr:   lipgloss.NewStyle().Italic(true).Foreground(theme.TextMuted),
		Version: lipgloss.NewStyle().MarginBottom(2).Foreground(theme.TextMuted),
	}
}

type Banner struct {
	Styles          BannerStyles
	GopherTypeAscii string
	descr           string
	version         string
}

func NewBanner(theme themes.Theme, version string) Banner {
	ascii := `
▄▀▀ █▀▄ █▀▄ █▄█ █▀▀ █▀▄   ▀█▀ ▀▄▀ █▀▄ █▀▀
█▄█ █▄█ █▀▀ █ █ ██▄ █▀▄    █   █  █▀▀ ██▄`
	ascii = strings.TrimSpace(ascii)

	return Banner{
		Styles:          newBannerStyles(theme),
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
