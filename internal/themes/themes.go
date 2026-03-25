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

var registry = []Theme{
	{
		Name:       "gopher type",
		Primary:    lipgloss.Color("#214d7a"),
		Secondary:  lipgloss.Color("#156868"),
		Text:       lipgloss.Color("#ffffff"),
		TextMuted:  lipgloss.Color("#919191"),
		Error:      lipgloss.Color("#b7355a"),
		Background: lipgloss.Color("#141618"),
	},
	{
		Name:       "gruvbox",
		Primary:    lipgloss.Color("#b8bb26"),
		Secondary:  lipgloss.Color("#85A598"),
		Text:       lipgloss.Color("#fbf1c7"),
		TextMuted:  lipgloss.Color("#9e9984"),
		Error:      lipgloss.Color("#fb4934"),
		Background: lipgloss.Color("#282828"),
	},
	{
		Name:       "dracula",
		Primary:    lipgloss.Color("#BD93F9"),
		Secondary:  lipgloss.Color("#FF79C6"),
		Text:       lipgloss.Color("#F8F8F2"),
		TextMuted:  lipgloss.Color("#9b9b9b"),
		Error:      lipgloss.Color("#FF5555"),
		Background: lipgloss.Color("#282A36"),
	},

	{
		Name:       "catppuccin",
		Primary:    lipgloss.Color("#F5C2E7"),
		Secondary:  lipgloss.Color("#cba6f7"),
		Text:       lipgloss.Color("#cdd6f4"),
		TextMuted:  lipgloss.Color("#9399b2"),
		Error:      lipgloss.Color("#F28FAD"),
		Background: lipgloss.Color("#181825"),
	},
	{
		Name:       "rose pine",
		Primary:    lipgloss.Color("#c4a7e7"),
		Secondary:  lipgloss.Color("#31748f"),
		Text:       lipgloss.Color("#e0def4"),
		TextMuted:  lipgloss.Color("#6e6a86"),
		Error:      lipgloss.Color("#eb6f92"),
		Background: lipgloss.Color("#191724"),
	},
}

func MustGetTheme(name string) Theme {
	for _, t := range registry {
		if name == t.Name {
			return t
		}
	}
	panic("no theme with that name")
}

func GetThemeNames() []string {
	names := make([]string, 0, len(registry))

	for _, p := range registry {
		names = append(names, p.Name)
	}

	return names
}

type ToggleTransparency struct{}
