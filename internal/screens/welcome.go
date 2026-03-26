package screens

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	comp "github.com/mati-33/gopher-type/internal/components"
	"github.com/mati-33/gopher-type/internal/ctx"
	"github.com/mati-33/gopher-type/internal/themes"
	"github.com/mati-33/gopher-type/internal/version"
)

type welcomeScreen struct {
	ctx      *ctx.Context
	banner   comp.Banner
	menu     comp.Menu
	info     comp.MenuInfo
	keybinds welcomeKeybinds
}

func NewWelcomeScreen(ctx *ctx.Context) welcomeScreen {
	keybinds := newWelcomeKeybind()
	banner := comp.NewBanner(ctx.Theme, version.Version)
	menu := comp.NewMenu(ctx.Theme, []comp.Keybind{
		keybinds.Practise,
		keybinds.Mode,
		keybinds.Theme,
		keybinds.Quit,
	}, lipgloss.Width(banner.GopherTypeAscii))

	return welcomeScreen{
		ctx:      ctx,
		keybinds: keybinds,
		banner:   banner,
		menu:     menu,
		info:     comp.NewMenuInfo(ctx.Theme, ctx.Mode.Name(), ctx.Theme.Name, lipgloss.Width(banner.GopherTypeAscii)),
	}
}

func (s *welcomeScreen) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {

	case ChangeProvider:
		s.info.ModeName = msg.Name

	case themes.Theme:
		s.info.ThemeName = msg.Name

	case tea.KeyMsg:
		switch msg.String() {

		case s.keybinds.Quit.Key:
			return tea.Quit

		case s.keybinds.Mode.Key:
			return pushScreen(NewModeScreen(s.ctx))

		case s.keybinds.Theme.Key:
			return pushScreen(NewThemeChangeScreen(s.ctx))

		case s.keybinds.Practise.Key:
			return pushScreen(NewTypingScreen(s.ctx))
		}

	}

	cmds := []tea.Cmd{
		s.banner.Update(msg),
		s.menu.Update(msg),
		s.info.Update(msg),
	}

	return tea.Batch(cmds...)
}

func (s *welcomeScreen) View() tea.View {
	bannerView := s.banner.View()
	menuView := s.menu.View()
	infoView := s.info.View()
	screen := lipgloss.JoinVertical(lipgloss.Center, bannerView, "", "", menuView, infoView)

	return tea.NewView(lipgloss.Place(
		s.ctx.Width, s.ctx.Height,
		lipgloss.Center, 0.7,
		screen,
	))
}

type welcomeKeybinds struct {
	Practise comp.Keybind
	Mode     comp.Keybind
	Theme    comp.Keybind
	Quit     comp.Keybind
}

func newWelcomeKeybind() welcomeKeybinds {
	return welcomeKeybinds{
		Practise: comp.Keybind{Key: "enter", Desc: "practise"},
		Mode:     comp.Keybind{Key: "m", Desc: "select mode"},
		Theme:    comp.Keybind{Key: "t", Desc: "change theme"},
		Quit:     comp.Keybind{Key: "q", Desc: "quit"},
	}
}
