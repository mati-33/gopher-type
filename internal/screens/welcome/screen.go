package welcome

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/screens"
	"github.com/mati-33/gopher-type/internal/screens/mode"
	"github.com/mati-33/gopher-type/internal/screens/typing"
	"github.com/mati-33/gopher-type/internal/screens/welcome/components"
	"github.com/mati-33/gopher-type/internal/version"
)

type welcomeScreen struct {
	width        int
	height       int
	banner       components.Banner
	menu         components.Menu
	providerName string
}

func NewWelcomeScreen(width, height int) welcomeScreen {
	banner := components.NewBanner(version.Version)
	menu := components.NewMenu([]components.MenuOption{
		{Key: "enter", Description: "practise"},
		{Key: "m", Description: "show modes"},
		{Key: "t", Description: "change theme"},
		{Key: "q", Description: "quit"},
	}, lipgloss.Width(banner.GopherTypeAscii))

	return welcomeScreen{
		width:        width,
		height:       height,
		banner:       banner,
		menu:         menu,
		providerName: "english",
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
		s.providerName = msg.Name

	case tea.KeyMsg:
		switch msg.String() {

		case "q":
			return s, tea.Quit

		case "m":
			return s, func() tea.Msg {
				return screens.PushScreen{
					Screen: mode.NewModeScreen(s.width, s.height),
				}
			}

		case "enter":
			return s, func() tea.Msg {
				return screens.PushScreen{
					Screen: typing.NewTypingScreen(
						s.providerName,
						15,
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
