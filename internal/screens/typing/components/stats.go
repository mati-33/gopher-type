package components

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

type StatsStyles struct {
	SpeedIcon    lipgloss.Style
	AccuracyIcon lipgloss.Style
	Text         lipgloss.Style
	Value        lipgloss.Style
	PositiveDiff lipgloss.Style
	NegativeDiff lipgloss.Style
}

func NewStatsStyles() StatsStyles {
	yellow := lipgloss.Color("#ffff30")
	blue := lipgloss.Color("#0087ff")
	green := lipgloss.Color("#3d9937")
	grey := lipgloss.Color("#bbbbbb")
	white := lipgloss.Color("#ffffff")
	red := lipgloss.Color("#ff005f")

	return StatsStyles{
		SpeedIcon:    lipgloss.NewStyle().Foreground(yellow),
		AccuracyIcon: lipgloss.NewStyle().Foreground(blue),
		Text:         lipgloss.NewStyle().Foreground(grey),
		Value:        lipgloss.NewStyle().Foreground(white),
		PositiveDiff: lipgloss.NewStyle().Foreground(green),
		NegativeDiff: lipgloss.NewStyle().Foreground(red),
	}
}

type StatsValues struct {
	Wpm      int
	Accuracy float64
}

type Stats struct {
	CurrentResult  *StatsValues
	PreviousResult *StatsValues
	Styles         StatsStyles
}

func NewStats() Stats {
	return Stats{
		Styles: NewStatsStyles(),
	}
}

func (s *Stats) UpdateStats(r *StatsValues) {
	prev := s.CurrentResult
	s.CurrentResult = r
	s.PreviousResult = prev
}

func (s Stats) View() string {
	if s.CurrentResult == nil {
		return fmt.Sprintf("%s %s  %s %s",
			s.Styles.SpeedIcon.Render("󱐋"),
			s.Styles.Text.Render("speed: -"),
			s.Styles.AccuracyIcon.Render("󰣉"),
			s.Styles.Text.Render("accuracy: -"),
		)
	}
	if s.PreviousResult == nil {
		return fmt.Sprintf("%s %s %swpm  %s %s %s",
			s.Styles.SpeedIcon.Render("󱐋"),
			s.Styles.Text.Render("speed:"),
			s.Styles.Value.Render(fmt.Sprintf("%d", s.CurrentResult.Wpm)),
			s.Styles.AccuracyIcon.Render("󰣉"),
			s.Styles.Text.Render("accuracy:"),
			s.Styles.Value.Render(fmt.Sprintf("%.2f%%", 100.0*s.CurrentResult.Accuracy)),
		)
	}

	wpmDiffStr := s.getWpmDiffStr()
	accuracyDiffStr := s.getAccuracyDiffStr()

	return fmt.Sprintf("%s %s %swpm%s  %s %s %s%s",
		s.Styles.SpeedIcon.Render("󱐋"),
		s.Styles.Text.Render("speed:"),
		s.Styles.Value.Render(fmt.Sprintf("%d", s.CurrentResult.Wpm)),
		wpmDiffStr,
		s.Styles.AccuracyIcon.Render("󰣉"),
		s.Styles.Text.Render("accuracy:"),
		s.Styles.Value.Render(fmt.Sprintf("%.2f%%", 100.0*s.CurrentResult.Accuracy)),
		accuracyDiffStr,
	)
}

func (s Stats) getWpmDiffStr() string {
	wpmDiff := s.CurrentResult.Wpm - s.PreviousResult.Wpm
	if wpmDiff == 0 {
		return ""
	}

	sign := "+"
	style := s.Styles.PositiveDiff
	if wpmDiff < 0.0 {
		sign = ""
		style = s.Styles.NegativeDiff
	}
	return style.Render(fmt.Sprintf(" (%s%dwpm)", sign, wpmDiff))
}

func (s Stats) getAccuracyDiffStr() string {
	accuracyDiff := s.CurrentResult.Accuracy - s.PreviousResult.Accuracy
	if accuracyDiff <= (1./1000.) && accuracyDiff >= (-1./1000.) {
		return ""
	}

	sign := "+"
	style := s.Styles.PositiveDiff
	if accuracyDiff < 0.0 {
		sign = ""
		style = s.Styles.NegativeDiff
	}
	return style.Render(fmt.Sprintf(" (%s%.2f%%)", sign, accuracyDiff*100.0))
}
