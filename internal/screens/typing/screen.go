package typing

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/modes"
	"github.com/mati-33/gopher-type/internal/screens"
	"github.com/mati-33/gopher-type/internal/screens/mode"
	"github.com/mati-33/gopher-type/internal/screens/typing/components"
)

type typingScreen struct {
	mode         modes.Mode
	providerName string
	width        int
	height       int
	wordCount    int
	stats        components.Stats
	text         components.Text
	info         components.Info
}

func NewTypingScreen(modeName string, wordCount, width, height int) typingScreen {
	mode := modes.MustGetMode(modeName)

	return typingScreen{
		mode:      mode,
		width:     width,
		height:    height,
		wordCount: wordCount,
		stats:     components.NewStats(),
		text:      components.NewText(mode.Generate(wordCount), int(float32(width)*0.7), height),
		info:      components.NewInfo(wordCount, mode.Name()),
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

	case components.TextResult:
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
		case "esc":
			if s.text.Started {
				s.text.Text = s.mode.Generate(s.wordCount)
				return s, tea.Batch(s.text.Reset()...)
			} else {
				return s, func() tea.Msg { return screens.PopScreen{} }
			}
		case "ctrl+o":
			s.wordCount++
			s.info.WordCount = s.wordCount
			s.text.Text = s.mode.Generate(s.wordCount)
			return s, tea.Batch(s.text.Reset()...)
		case "ctrl+p":
			if s.wordCount > 1 {
				s.wordCount--
				s.info.WordCount = s.wordCount
				s.text.Text = s.mode.Generate(s.wordCount)
				return s, tea.Batch(s.text.Reset()...)
			} else {
				return s, nil
			}
		case "ctrl+n":
			return s, func() tea.Msg {
				return screens.PushScreen{
					Screen: mode.NewModeScreen(s.width, s.height),
				}
			}
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
	)

	c := lipgloss.NewCompositor(textLayer)
	return tea.NewView(c.Render())
}
