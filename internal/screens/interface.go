package screens

import tea "charm.land/bubbletea/v2"

type Interface interface {
	Update(tea.Msg) tea.Cmd
	View() tea.View
}
