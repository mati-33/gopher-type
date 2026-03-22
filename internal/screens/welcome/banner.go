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
		Ascii: lipgloss.NewStyle().Foreground(theme.Text).
			Border(lipgloss.ThickBorder(), false, false, true, false).
			BorderForeground(theme.Primary),

		Descr:   lipgloss.NewStyle().Italic(true).Foreground(theme.TextMuted),
		Version: lipgloss.NewStyle().Foreground(theme.TextMuted),
	}
}

type Banner struct {
	Styles          BannerStyles
	GopherTypeAscii string
	descr           string
	version         string
	theme           themes.Theme
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
		theme:           theme,
	}
}

func (b Banner) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		b.Styles.Ascii.Render(b.GopherTypeAscii),
		b.Styles.Descr.Render(b.descr),
		b.Styles.Version.Render(b.version),
	)
}
