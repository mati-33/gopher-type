package welcome

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/config"
	"github.com/mati-33/gopher-type/internal/screens"
	"github.com/mati-33/gopher-type/internal/screens/mode"
	"github.com/mati-33/gopher-type/internal/screens/typing"
	"github.com/mati-33/gopher-type/internal/version"
)

type welcomeScreen struct {
	config   config.Config
	width    int
	height   int
	banner   Banner
	menu     Menu
	modeName string
	keybinds keybinds
}

func NewWelcomeScreen(config config.Config, width, height int) welcomeScreen {
	keybinds := newKeybind()
	banner := NewBanner(version.Version)
	menu := NewMenu([]screens.Keybind{
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

	case screens.ChangeProvider:
		s.modeName = msg.Name

	case tea.KeyMsg:
		switch msg.String() {

		case s.keybinds.Quit.Key:
			return s, tea.Quit

		case s.keybinds.Mode.Key:
			return s, func() tea.Msg {
				return screens.PushScreen{
					Screen: mode.NewModeScreen(s.config, s.width, s.height),
				}
			}

		case s.keybinds.Practise.Key:
			return s, func() tea.Msg {
				return screens.PushScreen{
					Screen: typing.NewTypingScreen(
						s.config,
						s.width,
						s.height,
					),
				}
			}
		}

	}
	return s, nil
}

func (s welcomeScreen) View() tea.View {
	bannerView := s.banner.View()
	menuView := s.menu.View()
	screen := lipgloss.JoinVertical(lipgloss.Center, bannerView, "", "", menuView)

	return tea.NewView(lipgloss.Place(
		s.width, s.height,
		lipgloss.Center, 0.7,
		screen,
	))
}
