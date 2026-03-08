package typing

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/screens"
	"github.com/mati-33/gopher-type/internal/screens/typing/components"
)

type TextProvider interface {
	Provide(maxLen int) []rune
}

type typingScreen struct {
	textProvider TextProvider
	width        int
	height       int
	textLen      int
	stats        components.Stats
	text         components.Text
}

func NewTypingScreen(textProvider TextProvider, textLen, width, height int) typingScreen {
	return typingScreen{
		textProvider: textProvider,
		width:        width,
		height:       height,
		textLen:      textLen,
		stats:        components.NewStats(),
		text:         components.NewText(textProvider.Provide(textLen), int(float32(width)*0.7), height),
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
		s.stats.UpdateStats(&components.StatsValues{
			Wpm: msg.Wpm, Accuracy: msg.Accuracy,
		})
		s.text.Text = s.textProvider.Provide(s.textLen)

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if s.text.Started {
				s.text.Text = s.textProvider.Provide(s.textLen)
				return s, tea.Batch(s.text.Reset()...)
			} else {
				return s, func() tea.Msg { return screens.PopScreen{} }
			}
		}
	}

	cmd := s.text.Update(msg)
	cmds = append(cmds, cmd)
	return s, tea.Batch(cmds...)
}

func (s typingScreen) View() tea.View {
	resultOffset := int(float32(s.height) * 0.2)

	if s.height < 14 {
		resultOffset = 0
	}

	statsView := lipgloss.PlaceHorizontal(s.width, lipgloss.Center, s.stats.View())

	textView := s.text.View()
	textLayer := lipgloss.NewLayer(lipgloss.Place(
		s.width, s.height,
		lipgloss.Center, lipgloss.Center,
		textView,
	),
		lipgloss.NewLayer(statsView).Y(resultOffset),
	)

	c := lipgloss.NewCompositor(textLayer)
	return tea.NewView(c.Render())
}
