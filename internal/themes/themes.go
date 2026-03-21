package themes

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

type Theme struct {
	Name string

	Primary    color.Color
	Secondary  color.Color
	Text       color.Color
	TextMuted  color.Color
	Error      color.Color
	Background color.Color
}

func NewDefault() Theme {
	return Theme{
		Name:       "gopher type",
		Primary:    lipgloss.Color("#214d7a"),
		Secondary:  lipgloss.Color("#156868"),
		Text:       lipgloss.Color("#ffffff"),
		TextMuted:  lipgloss.Color("#919191"),
		Error:      lipgloss.Color("#b7355a"),
		Background: lipgloss.Color("#141618"),
	}
}
