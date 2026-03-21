package mode

import "github.com/mati-33/gopher-type/internal/screens"

type keybinds struct {
	Next     screens.Keybind
	Previous screens.Keybind
	Refresh  screens.Keybind
	Choose   screens.Keybind
	Cancel   screens.Keybind
	Help     screens.Keybind
}

func newKeybinds() keybinds {
	return keybinds{
		Next:     screens.Keybind{Key: "j", Desc: "next"},
		Previous: screens.Keybind{Key: "k", Desc: "previous"},
		Refresh:  screens.Keybind{Key: "r", Desc: "refresh"},
		Choose:   screens.Keybind{Key: "enter", Desc: "choose"},
		Cancel:   screens.Keybind{Key: "esc", Desc: "cancel"},
		Help:     screens.Keybind{Key: "f1", Desc: "close help"},
	}
}
