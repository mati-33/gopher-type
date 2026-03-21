package welcome

import "github.com/mati-33/gopher-type/internal/screens"

type keybinds struct {
	Practise screens.Keybind
	Mode     screens.Keybind
	Theme    screens.Keybind
	Quit     screens.Keybind
}

func newKeybind() keybinds {
	return keybinds{
		Practise: screens.Keybind{Key: "enter", Desc: "practise"},
		Mode:     screens.Keybind{Key: "m", Desc: "select mode"},
		Theme:    screens.Keybind{Key: "t", Desc: "change theme"},
		Quit:     screens.Keybind{Key: "q", Desc: "quit"},
	}
}
