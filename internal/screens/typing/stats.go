package typing

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/themes"
)

type StatsStyles struct {
	SpeedIcon    lipgloss.Style
	AccuracyIcon lipgloss.Style
	Text         lipgloss.Style
	Value        lipgloss.Style
	PositiveDiff lipgloss.Style
	NegativeDiff lipgloss.Style
}

func NewStatsStyles(theme themes.Theme) StatsStyles {
	return StatsStyles{
		SpeedIcon:    lipgloss.NewStyle().Foreground(theme.Primary).SetString("󱐋"),
		AccuracyIcon: lipgloss.NewStyle().Foreground(theme.Primary).SetString("󰣉"),
		Text:         lipgloss.NewStyle().Foreground(theme.Text),
		Value:        lipgloss.NewStyle().Foreground(theme.Text),
	}
}

type Stats struct {
	Wpm      int
	Accuracy float64
	Styles   StatsStyles
}

func NewStats(theme themes.Theme) Stats {
	return Stats{
		Styles: NewStatsStyles(theme),
	}
}

func (s Stats) View() string {
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
