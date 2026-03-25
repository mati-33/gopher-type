package screens

import (
	tea "charm.land/bubbletea/v2"
)

type PushScreen struct {
	Screen Interface
}

type PopScreen struct {
	Command tea.Cmd
}

type ChangeProvider struct {
	Name string
}
