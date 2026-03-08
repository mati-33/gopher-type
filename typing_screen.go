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
	white  = lipgloss.Color("#ffffff")
	grey   = lipgloss.Color("#bbbbbb")
	red    = lipgloss.Color("#ff005f")
	blue   = lipgloss.Color("#0087ff")
	yellow = lipgloss.Color("#ffff30")
	green  = lipgloss.Color("#3d9937")

	cursorStyle = lipgloss.NewStyle().Underline(true).Foreground(grey)
	beforeStyle = lipgloss.NewStyle().Foreground(grey)
	afterStyle  = lipgloss.NewStyle().Foreground(white)
	errorStyle  = lipgloss.NewStyle().Foreground(red)
	textStyle   = lipgloss.NewStyle()
	lineStyle   = lipgloss.NewStyle().MarginTop(1)
	iconStyle   = lipgloss.NewStyle()
	diffStyle   = lipgloss.NewStyle()
)

type result struct {
	wpm      int
	accuracy float64
}

type StatsComponent struct {
	CurrentResult  *result
	PreviousResult *result
}

func (s StatsComponent) View() string {
	if s.CurrentResult == nil {
		return fmt.Sprintf("%s %s  %s %s",
			iconStyle.Foreground(yellow).Render("󱐋"),
			beforeStyle.Render("speed: -"),
			iconStyle.Foreground(blue).Render("󰣉"),
			beforeStyle.Render("accuracy: -"),
		)
	}
	if s.PreviousResult == nil {
		return fmt.Sprintf("%s %s %dwpm  %s %s %.2f%%",
			iconStyle.Foreground(yellow).Render("󱐋"),
			beforeStyle.Render("speed:"),
			s.CurrentResult.wpm,
			iconStyle.Foreground(blue).Render("󰣉"),
			beforeStyle.Render("accuracy:"),
			100.0*s.CurrentResult.accuracy,
		)
	}

	var wpmDiffStr string
	wpmDiff := s.CurrentResult.wpm - s.PreviousResult.wpm
	if wpmDiff == 0 {
		wpmDiffStr = ""
	} else {
		wpmDiffSign := "+"
		wpmDiffColor := green
		if wpmDiff < 0.0 {
			wpmDiffSign = ""
			wpmDiffColor = red
		}
		wpmDiffStr = diffStyle.Foreground(wpmDiffColor).Render(fmt.Sprintf(" (%s%dwpm)", wpmDiffSign, wpmDiff))
	}

	var accuracyDiffStr string
	accuracyDiff := s.CurrentResult.accuracy - s.PreviousResult.accuracy
	if accuracyDiff <= (1./1000.) && accuracyDiff >= (-1./1000.) {
		accuracyDiffStr = ""
	} else {
		accuracyDiffSign := "+"
		accuracyDiffColor := green
		if accuracyDiff < 0.0 {
			accuracyDiffSign = ""
			accuracyDiffColor = red
		}
		accuracyDiffStr = diffStyle.Foreground(accuracyDiffColor).Render(fmt.Sprintf(" (%s%.2f%%)", accuracyDiffSign, accuracyDiff*100.0))
	}

	return fmt.Sprintf("%s %s %dwpm%s  %s %s %.2f%%%s",
		iconStyle.Foreground(yellow).Render("󱐋"),
		beforeStyle.Render("speed:"),
		s.CurrentResult.wpm,
		wpmDiffStr,
		iconStyle.Foreground(blue).Render("󰣉"),
		beforeStyle.Render("accuracy:"),
		100.0*s.CurrentResult.accuracy,
		accuracyDiffStr,
	)
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
	stopwatch    stopwatch.Model
	textLen      int
	stats        StatsComponent
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
		stats:        StatsComponent{},
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
				result := calculateResult(len(s.text), len(s.errors), s.stopwatch.Elapsed())
				previousResult := s.stats.CurrentResult
				s.stats.CurrentResult = &result
				s.stats.PreviousResult = previousResult
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
	textWidth := int(float32(s.width) * 0.7)
	resultOffset := int(float32(s.height) * 0.2)
	if s.height < 14 {
		resultOffset = 0
	}

	linesStr := []string{}
	lines := splitText(s.text, textWidth)
	i := 0

	for _, line := range lines {
		b := strings.Builder{}
		for _, rune := range line {
			char := string(rune)
			var style lipgloss.Style

			switch {
			case i < s.cursor && slices.Contains(s.errors, i):
				if char == " " {
					char = "."
				}
				style = errorStyle
			case i < s.cursor:
				style = afterStyle
			case i == s.cursor:
				style = cursorStyle
			default:
				style = beforeStyle
			}

			b.WriteString(style.Render(char))
			i++
		}

		linesStr = append(linesStr, lineStyle.Render(b.String()))
	}

	resultView := textStyle.
		Width(s.width).
		Align(lipgloss.Center).
		Render(s.stats.View())

	textView := lipgloss.JoinVertical(0, linesStr...)
	textLayer := lipgloss.NewLayer(lipgloss.Place(
		s.width, s.height,
		lipgloss.Center, lipgloss.Center,
		textView,
	),
		lipgloss.NewLayer(resultView).Y(resultOffset),
	)

	c := lipgloss.NewCompositor(textLayer)
	return tea.NewView(c.Render())
}

func splitText(text []rune, width int) [][]rune {
	words := strings.Fields(string(text))

	if len(words) == 0 {
		return nil
	}

	var lines [][]rune

	i := 0
	for i < len(words) {
		var line []rune

		for i < len(words) {
			word := []rune(words[i] + " ")

			if len(line)+len(word) <= width || len(line) == 0 {
				line = append(line, word...)
				i++
			} else {
				break
			}
		}

		lines = append(lines, line)
	}

	return lines
}

func calculateResult(runesNo, errorsNo int, elapsed time.Duration) result {
	wpm := calculateWpm(runesNo, elapsed)
	acc := calculateAccuracy(runesNo, errorsNo)
	return result{
		wpm:      wpm,
		accuracy: acc,
	}
}

func calculateWpm(runesNo int, elapsed time.Duration) int {
	return int(math.Round(float64(runesNo) / (5.0 * elapsed.Seconds()) * 60.0))
}

func calculateAccuracy(runesNo, errorsNo int) float64 {
	return math.Floor(float64(runesNo-errorsNo)/float64(runesNo)*10000) / 10000
}
