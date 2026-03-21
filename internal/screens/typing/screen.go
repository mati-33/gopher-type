package typing

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/config"
	"github.com/mati-33/gopher-type/internal/modes"
	"github.com/mati-33/gopher-type/internal/screens"
	"github.com/mati-33/gopher-type/internal/screens/mode"
)

type typingScreen struct {
	config       config.Config
	mode         modes.Mode
	providerName string
	width        int
	height       int
	wordCount    int
	stats        Stats
	text         Text
	info         Info
	help         screens.Help
	keybinds     keybinds
}

func NewTypingScreen(config config.Config, width, height int) typingScreen {
	mode := modes.MustGetMode(config.InitMode)
	wc := config.InitWordCount
	keybinds := newKeybinds()

	return typingScreen{
		mode:      mode,
		width:     width,
		height:    height,
		wordCount: config.InitWordCount,
		stats:     NewStats(),
		text:      NewText(mode.Generate(wc), int(float32(width)*0.7), height),
		info:      NewInfo(wc, mode.Name()),
		config:    config,
		keybinds:  keybinds,
		help: screens.NewHelp([]screens.Keybind{
			keybinds.IncWordCount,
			keybinds.DecWordCount,
			keybinds.ChangeMode,
			keybinds.GoBack,
			keybinds.Help,
		}),
	}
}

func (s typingScreen) Init() tea.Cmd {
	return tea.ClearScreen
}

func (s typingScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
		s.text.Width = int(float32(s.width) * 0.7)

	case TextResult:
		s.stats.Wpm = msg.Wpm
		s.stats.Accuracy = msg.Accuracy
		s.text.Text = s.mode.Generate(s.wordCount)

	case screens.ChangeProvider:
		s.mode = modes.MustGetMode(msg.Name)
		s.info.Mode = s.mode.Name()
		s.text.Text = s.mode.Generate(s.wordCount)
		return s, tea.Batch(s.text.Reset()...)

	case tea.KeyMsg:
		switch msg.String() {

		case s.keybinds.GoBack.Key:
			if s.text.Started {
				s.text.Text = s.mode.Generate(s.wordCount)
				return s, tea.Batch(s.text.Reset()...)
			} else {
				return s, func() tea.Msg { return screens.PopScreen{} }
			}

		case s.keybinds.IncWordCount.Key:
			s.wordCount++
			s.info.WordCount = s.wordCount
			s.text.Text = s.mode.Generate(s.wordCount)
			return s, tea.Batch(s.text.Reset()...)

		case s.keybinds.DecWordCount.Key:
			if s.wordCount > 1 {
				s.wordCount--
				s.info.WordCount = s.wordCount
				s.text.Text = s.mode.Generate(s.wordCount)
				return s, tea.Batch(s.text.Reset()...)
			} else {
				return s, nil
			}

		case s.keybinds.ChangeMode.Key:
			return s, func() tea.Msg {
				return screens.PushScreen{
					Screen: mode.NewModeScreen(s.config, s.width, s.height),
				}
			}

		case s.keybinds.Help.Key:
			s.help.Toggle()
			return s, nil
		}
	}

	cmd := s.text.Update(msg)
	cmds = append(cmds, cmd)
	return s, tea.Batch(cmds...)
}

func (s typingScreen) View() tea.View {
	bannerOffset := int(float32(s.height) * 0.2)

	if s.height < 14 {
		bannerOffset = 0
	}

	statsView := s.stats.View()
	infoView := s.info.View()
	helpView := s.help.View()

	helpOffset := max(0, s.width-lipgloss.Width(helpView)-2)

	bannerView := lipgloss.PlaceHorizontal(
		s.width,
		lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Top, statsView, "     ", infoView),
	)

	textView := s.text.View()
	textLayer := lipgloss.NewLayer(lipgloss.Place(
		s.width, s.height,
		lipgloss.Center, lipgloss.Center,
		textView,
	),
		lipgloss.NewLayer(bannerView).Y(bannerOffset),
		lipgloss.NewLayer(helpView).Y(s.height-lipgloss.Height(helpView)).X(helpOffset),
	)

	c := lipgloss.NewCompositor(textLayer)
	return tea.NewView(c.Render())
}
