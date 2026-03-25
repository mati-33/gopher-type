package screens

import tea "charm.land/bubbletea/v2"

func pushScreen(screen Interface) tea.Cmd {
	return func() tea.Msg {
		return PushScreen{Screen: screen}
	}
}

func popScreen(command tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		return PopScreen{Command: command}
	}
}
