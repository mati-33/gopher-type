package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

type typingScreen struct {
	text       []rune
	errors     []int
	cursor     int
	width      int
	height     int
	lastResult result
	stopwatch  stopwatch.Model
}

func newTypingScreen(text []rune, width, height int) typingScreen {
	return typingScreen{
		text:      text,
		errors:    []int{},
		cursor:    0,
		width:     width,
		height:    height,
		stopwatch: stopwatch.NewWithInterval(time.Millisecond),
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
				cmds = append(cmds, s.stopwatch.Stop(), s.stopwatch.Reset())
			}

		default:
			if s.cursor == 0 {
				cmds = append(cmds, s.stopwatch.Start())
			}

			current := string(s.text[s.cursor])
			if msg.String() != current {
				s.errors = append(s.errors, s.cursor)
			}

			if s.cursor < len(s.text)-1 {
				s.cursor++
			} else {
				s.lastResult = calculateResult(len(s.text), len(s.errors), s.stopwatch.Elapsed())
				s.cursor = 0
				s.errors = []int{}
				cmds = append(cmds, s.stopwatch.Stop(), s.stopwatch.Reset())
			}
		}
	}

	stopwatch, cmd := s.stopwatch.Update(msg)
	s.stopwatch = stopwatch
	cmds = append(cmds, cmd)
	return s, tea.Batch(cmds...)
}

func (s typingScreen) View() string {
	var view string

	b := strings.Builder{}

	for idx, ch := range s.text[:s.cursor] {
		if slices.Contains(s.errors, idx) {
			if string(ch) == " " {
				b.WriteString(errorStyle.Render(""))
			} else {
				b.WriteString(errorStyle.Render(string(ch)))
			}
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
		Render(s.lastResult.View())
	textView := textStyle.
		Width(textWidth).
		Height(3).
		Align(lipgloss.Center, lipgloss.Center).
		Render(view)
	stopwatchView := fmt.Sprintf("elapsed: %s", s.stopwatch.View())

	return lipgloss.Place(
		s.width, s.height,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Left, resultView, textView, stopwatchView),
	)
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
