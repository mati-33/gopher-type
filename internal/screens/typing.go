package screens

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	comp "github.com/mati-33/gopher-type/internal/components"
	"github.com/mati-33/gopher-type/internal/config"
	"github.com/mati-33/gopher-type/internal/modes"
	"github.com/mati-33/gopher-type/internal/themes"
)

type typingScreen struct {
	config       config.Config
	theme        themes.Theme
	mode         modes.Mode
	providerName string
	width        int
	height       int
	wordCount    int
	stats        comp.TypingStats
	text         comp.Text
	info         comp.TypingInfo
	help         comp.Help
	keybinds     typingKeybinds
}

func NewTypingScreen(config config.Config, theme themes.Theme, width, height int) typingScreen {
	mode := modes.MustGetMode(config.InitMode)
	wc := config.InitWordCount
	keybinds := newTypingKeybinds()

	return typingScreen{
		mode:      mode,
		width:     width,
		height:    height,
		wordCount: config.InitWordCount,
		stats:     comp.NewTypingStats(theme),
		text:      comp.NewText(theme, mode.Generate(wc), int(float32(width)*0.7), height),
		info:      comp.NewTypingInfo(theme, wc, mode.Name()),
		config:    config,
		theme:     theme,
		keybinds:  keybinds,
		help: comp.NewHelp(theme, []comp.Keybind{
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

	case comp.TextResult:
		s.stats.Wpm = msg.Wpm
		s.stats.Accuracy = msg.Accuracy
		s.text.Text = s.mode.Generate(s.wordCount)

	case ChangeProvider:
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
				return s, func() tea.Msg { return PopScreen{} }
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
				return PushScreen{
					Screen: NewModeScreen(s.config, s.theme, s.width, s.height),
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

type typingKeybinds struct {
	IncWordCount comp.Keybind
	DecWordCount comp.Keybind
	ChangeMode   comp.Keybind
	GoBack       comp.Keybind
	Help         comp.Keybind
}

func newTypingKeybinds() typingKeybinds {
	return typingKeybinds{
		IncWordCount: comp.Keybind{Key: "ctrl+o", Desc: "increase word count"},
		DecWordCount: comp.Keybind{Key: "ctrl+p", Desc: "decrease word count"},
		ChangeMode:   comp.Keybind{Key: "ctrl+n", Desc: "change mode"},
		GoBack:       comp.Keybind{Key: "esc", Desc: "go back"},
		Help:         comp.Keybind{Key: "f1", Desc: "close help"},
	}
}
