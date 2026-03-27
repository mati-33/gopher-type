package appcontex

import (
	"github.com/mati-33/gopher-type/internal/config"
	"github.com/mati-33/gopher-type/internal/modes"
	"github.com/mati-33/gopher-type/internal/themes"
)

type AppContext struct {
	Config *config.Config
	Theme  themes.Theme
	Mode   modes.Mode
	Width  int
	Height int
}

func New() *AppContext {
	cfg := config.NewDefault()

	return &AppContext{
		Config: cfg,
		Theme:  themes.MustGetTheme(cfg.InitTheme),
		Mode:   modes.MustGetMode(cfg.InitMode),
	}
}
