package screens

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	comp "github.com/mati-33/gopher-type/internal/components"
	"github.com/mati-33/gopher-type/internal/config"
	"github.com/mati-33/gopher-type/internal/themes"
	"github.com/mati-33/gopher-type/internal/version"
)

type welcomeScreen struct {
	config   config.Config
	width    int
	height   int
	banner   comp.Banner
	menu     comp.Menu
	modeName string
	keybinds welcomeKeybinds
	theme    themes.Theme
	info     comp.MenuInfo
}

func NewWelcomeScreen(config config.Config, theme themes.Theme, width, height int) welcomeScreen {
	keybinds := newWelcomeKeybind()
	banner := comp.NewBanner(theme, version.Version)
	menu := comp.NewMenu(theme, []comp.Keybind{
		keybinds.Practise,
		keybinds.Mode,
		keybinds.Theme,
		keybinds.Quit,
	}, lipgloss.Width(banner.GopherTypeAscii))

	return welcomeScreen{
		width:    width,
		height:   height,
		banner:   banner,
		menu:     menu,
		config:   config,
		modeName: config.InitMode,
		keybinds: keybinds,
		theme:    theme,
		info:     comp.NewMenuInfo(theme, config.InitMode, config.InitTheme, lipgloss.Width(banner.GopherTypeAscii)),
	}
}

func (s welcomeScreen) Init() tea.Cmd {
	return nil
}

func (s welcomeScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
		return s, nil

	case ChangeProvider:
		s.modeName = msg.Name

	case themes.Theme:
		s.theme = msg

	case tea.KeyMsg:
		switch msg.String() {

		case s.keybinds.Quit.Key:
			return s, tea.Quit

		case s.keybinds.Mode.Key:
			return s, func() tea.Msg {
				return PushScreen{
					Screen: NewModeScreen(s.config, s.theme, s.width, s.height),
				}
			}

		case s.keybinds.Theme.Key:
			return s, func() tea.Msg {
				return PushScreen{
					Screen: NewThemeChangeScreen(s.config, s.theme),
				}
			}

		case s.keybinds.Practise.Key:
			return s, func() tea.Msg {
				return PushScreen{
					Screen: NewTypingScreen(
						s.config,
						s.theme,
						s.width,
						s.height,
					),
				}
			}
		}

	}

	cmds := []tea.Cmd{
		s.banner.Update(msg),
		s.menu.Update(msg),
		s.info.Update(msg),
	}

	return s, tea.Batch(cmds...)
}

func (s welcomeScreen) View() tea.View {
	bannerView := s.banner.View()
	menuView := s.menu.View()
	infoView := s.info.View()
	screen := lipgloss.JoinVertical(lipgloss.Center, bannerView, "", "", menuView, infoView)

	return tea.NewView(lipgloss.Place(
		s.width, s.height,
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
