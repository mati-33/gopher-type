package screens

import (
	tea "charm.land/bubbletea/v2"
)

type Push struct {
	Screen Interface
}

type Pop struct {
	Command tea.Cmd
}

type ChangeProvider struct {
	Name string
}
