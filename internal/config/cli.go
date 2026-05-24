package config

import (
	"flag"
	"fmt"
	"slices"

	"github.com/mati-33/gopher-type/internal/modes"
	"github.com/mati-33/gopher-type/internal/themes"
)

type cliConfig struct {
	Theme       string
	Mode        string
	Transparent bool
	NoIcons     bool
}

func parseCliConfig() (*cliConfig, error) {
	var cc cliConfig

	flag.StringVar(&cc.Theme, "theme", "", "set color theme")
	flag.StringVar(&cc.Mode, "mode", "", "set initial practise mode")
	flag.BoolVar(&cc.Transparent, "transparent", false, "enable transparent background")
	flag.BoolVar(&cc.NoIcons, "no-icons", false, "don't use nerd font icons")

	flag.Parse()

	modes := modes.GetModeNames()
	if cc.Mode != "" && !slices.Contains(modes, cc.Mode) {
		return nil, fmt.Errorf("invalid mode: %s", cc.Mode)
	}

	themes := themes.GetThemeNames()
	if cc.Theme != "" && !slices.Contains(themes, cc.Theme) {
		return nil, fmt.Errorf("ivalid theme: %s", cc.Theme)
	}

	return &cc, nil
}
