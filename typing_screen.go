package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	"charm.land/bubbles/v2/stopwatch"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	white = lipgloss.Color("#ffffff")
	grey  = lipgloss.Color("#bbbbbb")
	red   = lipgloss.Color("#ff005f")

	cursorStyle = lipgloss.NewStyle().Underline(true).Foreground(grey)
	beforeStyle = lipgloss.NewStyle().Foreground(grey)
	afterStyle  = lipgloss.NewStyle().Foreground(white)
	errorStyle  = lipgloss.NewStyle().Foreground(red)
	textStyle   = lipgloss.NewStyle()
)

type result struct {
	wpm      int
	accuracy float64
	wpmc     int
}

func (r result) View() string {
	if r.wpm == 0 {
		return "wpm: -  accuracy: -  wpmc: -"
	}
	return fmt.Sprintf("wpm: %d  accuracy: %.2f%% wpmc: %d", r.wpm, 100.0*r.accuracy, r.wpmc)
}

type TextProvider interface {
	Provide(maxLen int) []rune
}

type typingScreen struct {
	text         []rune
	textProvider TextProvider
	errors       []int
	cursor       int
	width        int
	height       int
	lastResult   result
	stopwatch    stopwatch.Model
	textLen      int
}

func newTypingScreen(textProvider TextProvider, textLen, width, height int) typingScreen {
	return typingScreen{
		text:         textProvider.Provide(textLen),
		textProvider: textProvider,
		errors:       []int{},
		cursor:       0,
		width:        width,
		height:       height,
		stopwatch:    stopwatch.New(stopwatch.WithInterval(time.Millisecond)),
		textLen:      textLen,
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

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if s.cursor == 0 {
				return s, func() tea.Msg { return PopScreen{} }
			} else {
				s.cursor = 0
				s.errors = []int{}
				s.text = s.textProvider.Provide(s.textLen)
				cmds = append(cmds, s.stopwatch.Stop(), s.stopwatch.Reset())
			}

		default:
			if s.cursor == 0 {
				cmds = append(cmds, s.stopwatch.Start())
			}

			expected := string(s.text[s.cursor])
			got := msg.String()

			if got == "space" {
				got = " "
			}

			if got != expected {
				s.errors = append(s.errors, s.cursor)
			}

			if s.cursor < len(s.text)-1 {
				s.cursor++
			} else {
				s.lastResult = calculateResult(len(s.text), len(s.errors), s.stopwatch.Elapsed())
				s.cursor = 0
				s.errors = []int{}
				s.text = s.textProvider.Provide(s.textLen)
				cmds = append(cmds, s.stopwatch.Stop(), s.stopwatch.Reset())
			}
		}
	}

	stopwatch, cmd := s.stopwatch.Update(msg)
	s.stopwatch = stopwatch
	cmds = append(cmds, cmd)
	return s, tea.Batch(cmds...)
}

func (s typingScreen) View() tea.View {
	var view string

	b := strings.Builder{}

	for idx, ch := range s.text[:s.cursor] {
		if slices.Contains(s.errors, idx) {
			wrongChar := string(ch)
			if wrongChar == " " {
				wrongChar = ""
			}
			b.WriteString(errorStyle.Render(wrongChar))
		} else {
			b.WriteString(afterStyle.Render(string(ch)))
		}
	}

	b.WriteString(cursorStyle.Render(string(s.text[s.cursor])))
	b.WriteString(beforeStyle.Render(string(s.text[s.cursor+1:])))

	view = b.String()

	textWidth := int(float32(s.width) * 0.6)
	resultView := textStyle.
		Width(textWidth).
		Height(1).
		Align(lipgloss.Left).
		MarginBottom(1).
		Render(s.lastResult.View())
	textView := textStyle.
		Width(textWidth).
		Height(3).
		Align(lipgloss.Left, lipgloss.Center).
		Render(view)
	stopwatchView := fmt.Sprintf("elapsed: %s", s.stopwatch.View())

	return tea.NewView(lipgloss.Place(
		s.width, s.height,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Left, resultView, textView, stopwatchView),
	))
}

func calculateResult(runesNo, errorsNo int, elapsed time.Duration) result {
	wpm := calculateWpm(runesNo, elapsed)
	acc := calculateAccuracy(runesNo, errorsNo)
	return result{
		wpm:      wpm,
		accuracy: acc,
		wpmc:     int(float64(wpm) * acc),
	}
}

func calculateWpm(runesNo int, elapsed time.Duration) int {
	return int(math.Round(float64(runesNo) / (5.0 * elapsed.Seconds()) * 60.0))
}

func calculateAccuracy(runesNo, errorsNo int) float64 {
	return math.Floor(float64(runesNo-errorsNo)/float64(runesNo)*10000) / 10000
}
