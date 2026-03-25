package components

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/themes"
)

type Keybind struct {
	Key  string
	Desc string
}

type HelpStyles struct {
	Key  lipgloss.Style
	Desc lipgloss.Style
}

func newHelpStyles(theme themes.Theme) HelpStyles {
	return HelpStyles{
		Key:  lipgloss.NewStyle().Foreground(theme.Secondary).Bold(true),
		Desc: lipgloss.NewStyle().Foreground(theme.Text),
	}
}

type Help struct {
	Styles   HelpStyles
	Expanded bool
	Keybinds []Keybind
}

func NewHelp(theme themes.Theme, keybinds []Keybind) Help {
	return Help{
		Styles:   newHelpStyles(theme),
		Keybinds: keybinds,
	}
}

func (h *Help) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case themes.Theme:
		h.Styles = newHelpStyles(msg)
	}

	return nil
}

func (h Help) View() string {
	if !h.Expanded {
		return fmt.Sprintf("%s %s", h.Styles.Key.Render("f1"), h.Styles.Desc.Render("help"))
	}

	keys := make([]string, 0, len(h.Keybinds))
	descs := make([]string, 0, len(h.Keybinds))

	for _, k := range h.Keybinds {
		keys = append(keys, h.Styles.Key.Render(k.Key))
		descs = append(descs, h.Styles.Desc.Render(k.Desc))
	}

	keysView := lipgloss.JoinVertical(lipgloss.Left, keys...)
	descsView := lipgloss.JoinVertical(lipgloss.Left, descs...)

	return lipgloss.JoinHorizontal(lipgloss.Top, keysView, "  ", descsView)
}

func (h *Help) Toggle() {
	h.Expanded = !h.Expanded
}
