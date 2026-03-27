package screens

import tea "charm.land/bubbletea/v2"

func push(screen Interface) tea.Cmd {
	return func() tea.Msg {
		return Push{Screen: screen}
	}
}

func pop(command tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		return Pop{Command: command}
	}
}
