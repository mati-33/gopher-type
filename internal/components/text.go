package components

import (
	"math"
	"slices"
	"strings"
	"time"

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
	Text      []rune
	Started   bool
	Styles    TextStyles
	Cursor    int
	Errors    []int
	startedAt time.Time
}

func NewText(theme themes.Theme, text []rune) Text {
	return Text{
		Started: false,
		Styles:  NewTextStyles(theme),
		Cursor:  0,
		Text:    text,
		Errors:  []int{},
	}
}

func (t *Text) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {

	case themes.Theme:
		t.Styles = NewTextStyles(msg)

	case tea.KeyMsg:
		if t.Cursor == 0 {
			t.Started = true
			t.startedAt = time.Now()
		}

		got := msg.String()
		expected := string(t.Text[t.Cursor])

		if got == "space" {
			got = " "
		}

		if got != expected {
			t.Errors = append(t.Errors, t.Cursor)
		}

		if t.Cursor < len(t.Text)-1 {
			t.Cursor++
		} else {
			wpm := calculateWpm(len(t.Text), time.Since(t.startedAt))
			acc := calculateAccuracy(len(t.Text), len(t.Errors))
			t.Reset()

			return func() tea.Msg {
				return TextResult{
					Wpm:      wpm,
					Accuracy: acc,
				}
			}
		}

	}

	return nil
}

func (t *Text) View(width int) string {
	linesStr := []string{}
	lines := splitText(t.Text, width)
	i := 0

	for _, line := range lines {
		b := strings.Builder{}
		for _, rune := range line {
			char := string(rune)
			var style lipgloss.Style

			switch {
			case i < t.Cursor && slices.Contains(t.Errors, i):
				if char == " " {
					char = "."
				}
				style = t.Styles.Error
			case i < t.Cursor:
				style = t.Styles.After
			case i == t.Cursor:
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

func (t *Text) Reset() {
	t.Cursor = 0
	t.Errors = []int{}
	t.Started = false
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
