package typing

import "github.com/mati-33/gopher-type/internal/screens"

type keybinds struct {
	IncWordCount screens.Keybind
	DecWordCount screens.Keybind
	ChangeMode   screens.Keybind
	GoBack       screens.Keybind
	Help         screens.Keybind
}

func newKeybinds() keybinds {
	return keybinds{
		IncWordCount: screens.Keybind{Key: "ctrl+o", Desc: "increase word count"},
		DecWordCount: screens.Keybind{Key: "ctrl+p", Desc: "decrease word count"},
		ChangeMode:   screens.Keybind{Key: "ctrl+n", Desc: "change mode"},
		GoBack:       screens.Keybind{Key: "esc", Desc: "go back"},
		Help:         screens.Keybind{Key: "f1", Desc: "close help"},
	}
}
