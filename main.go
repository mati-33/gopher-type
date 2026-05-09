package main

import (
	"errors"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/mati-33/gopher-type/internal/app"
	"github.com/mati-33/gopher-type/internal/appcontex"
	"github.com/mati-33/gopher-type/internal/config"
)

func main() {
	f, err := tea.LogToFile("debug.log", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup debug logging: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	userConfig, err := config.LoadUserConfig()
	if err != nil && !errors.Is(err, config.ErrConfigNotFound) {
		fmt.Fprintf(os.Stderr, "error in configuration file: %v\n", err)
		os.Exit(1)
	}

	cfg := config.New(userConfig)
	appctx := appcontex.New(cfg)
	app := app.New(appctx)
	p := tea.NewProgram(app)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start the program: %v", err)
		os.Exit(1)
	}
}
