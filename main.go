package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/mati-33/gopher-type/internal/app"
	"github.com/mati-33/gopher-type/internal/appcontex"
	"github.com/mati-33/gopher-type/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	appctx := appcontex.New(cfg)
	app := app.New(appctx)
	p := tea.NewProgram(app)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start the program: %v", err)
		os.Exit(1)
	}
}
