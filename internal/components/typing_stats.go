package components

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/themes"
)

type TypingStatsStyles struct {
	SpeedIcon    lipgloss.Style
	AccuracyIcon lipgloss.Style
	Text         lipgloss.Style
	Value        lipgloss.Style
	PositiveDiff lipgloss.Style
	NegativeDiff lipgloss.Style
}

func NewTypingStatsStyles(theme themes.Theme) TypingStatsStyles {
	return TypingStatsStyles{
		SpeedIcon:    lipgloss.NewStyle().Foreground(theme.Primary).SetString("󱐋"),
		AccuracyIcon: lipgloss.NewStyle().Foreground(theme.Primary).SetString("󰣉"),
		Text:         lipgloss.NewStyle().Foreground(theme.Text),
		Value:        lipgloss.NewStyle().Foreground(theme.Text),
	}
}

type TypingStats struct {
	Wpm      int
	Accuracy float64
	Styles   TypingStatsStyles
}

func NewTypingStats(theme themes.Theme) TypingStats {
	return TypingStats{
		Styles: NewTypingStatsStyles(theme),
	}
}

func (s TypingStats) View() string {
	speedLabel := fmt.Sprintf("%s %s", s.Styles.SpeedIcon.Render(), s.Styles.Text.Render("speed"))
	accuracyLabel := fmt.Sprintf("%s %s", s.Styles.AccuracyIcon.Render(), s.Styles.Text.Render("accuracy"))

	if s.Wpm == 0 {
		return lipgloss.JoinHorizontal(lipgloss.Top,
			lipgloss.JoinVertical(lipgloss.Left,
				speedLabel, s.Styles.Text.Render("  -")),
			"   ",
			lipgloss.JoinVertical(lipgloss.Left,
				accuracyLabel, s.Styles.Text.Render("  -")),
		)
	}

	speed := lipgloss.JoinVertical(lipgloss.Left,
		speedLabel, s.Styles.Value.Render(fmt.Sprintf("  %dwpm", s.Wpm)),
	)
	accuracy := lipgloss.JoinVertical(lipgloss.Left,
		accuracyLabel, s.Styles.Value.Render(fmt.Sprintf("  %.2f%%", 100.0*s.Accuracy)),
	)

	return lipgloss.JoinHorizontal(lipgloss.Top, speed, "   ", accuracy)
}
