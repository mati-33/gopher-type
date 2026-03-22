package components

import (
	"math"
	"slices"
	"strings"
	"time"

	"charm.land/bubbles/v2/stopwatch"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/themes"
)

type TextStyles struct {
	Before lipgloss.Style
	After  lipgloss.Style
	Cursor lipgloss.Style
	Error  lipgloss.Style
	Lines  lipgloss.Style
}

func NewTextStyles(theme themes.Theme) TextStyles {
	return TextStyles{
		Before: lipgloss.NewStyle().Foreground(theme.TextMuted),
		After:  lipgloss.NewStyle().Foreground(theme.Text),
		Cursor: lipgloss.NewStyle().Underline(true).Foreground(theme.TextMuted),
		Error:  lipgloss.NewStyle().Foreground(theme.Error),
		Lines:  lipgloss.NewStyle().MarginTop(1).Align(lipgloss.Left),
	}
}

type Text struct {
	Width     int
	Height    int
	Text      []rune
	Started   bool
	Styles    TextStyles
	cursor    int
	errors    []int
	stopwatch stopwatch.Model
}

func NewText(theme themes.Theme, text []rune, width, height int) Text {
	return Text{
		Width:     width,
		Height:    height,
		Started:   false,
		Styles:    NewTextStyles(theme),
		cursor:    0,
		Text:      text,
		errors:    []int{},
		stopwatch: stopwatch.New(stopwatch.WithInterval(time.Millisecond)),
	}
}

func (t *Text) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if t.cursor == 0 {
			cmds = append(cmds, t.stopwatch.Start())
			t.Started = true
		}

		got := msg.String()
		expected := string(t.Text[t.cursor])

		if got == "space" {
			got = " "
		}

		if got != expected {
			t.errors = append(t.errors, t.cursor)
		}

		if t.cursor < len(t.Text)-1 {
			t.cursor++
		} else {
			wpm := calculateWpm(len(t.Text), t.stopwatch.Elapsed())
			acc := calculateAccuracy(len(t.Text), len(t.errors))

			cmd := func() tea.Msg {
				return TextResult{
					Wpm:      wpm,
					Accuracy: acc,
				}
			}
			cmds = append(cmds, t.Reset()...)
			cmds = append(cmds, cmd)
		}

	}

	stopwatch, cmd := t.stopwatch.Update(msg)
	t.stopwatch = stopwatch
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

func (t *Text) View() string {
	linesStr := []string{}
	lines := splitText(t.Text, t.Width)
	i := 0

	for _, line := range lines {
		b := strings.Builder{}
		for _, rune := range line {
			char := string(rune)
			var style lipgloss.Style

			switch {
			case i < t.cursor && slices.Contains(t.errors, i):
				if char == " " {
					char = "."
				}
				style = t.Styles.Error
			case i < t.cursor:
				style = t.Styles.After
			case i == t.cursor:
				style = t.Styles.Cursor
			default:
				style = t.Styles.Before
			}

			b.WriteString(style.Render(char))
			i++
		}

		linesStr = append(linesStr, t.Styles.Lines.Render(b.String()))
	}
	return lipgloss.JoinVertical(0, linesStr...)
}

func (t *Text) Reset() []tea.Cmd {
	t.cursor = 0
	t.errors = []int{}
	t.Started = false

	return []tea.Cmd{
		t.stopwatch.Stop(),
		t.stopwatch.Reset(),
	}
}

type TextResult struct {
	Wpm      int
	Accuracy float64
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

func calculateWpm(runesNo int, elapsed time.Duration) int {
	return int(math.Round(float64(runesNo) / (5.0 * elapsed.Seconds()) * 60.0))
}

func calculateAccuracy(runesNo, errorsNo int) float64 {
	return math.Floor(float64(runesNo-errorsNo)/float64(runesNo)*10000) / 10000
}
