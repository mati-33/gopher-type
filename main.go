package main

import (
	"flag"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/mati-33/gopher-type/internal/app"
	"github.com/mati-33/gopher-type/internal/appcontex"
	"github.com/mati-33/gopher-type/internal/config"
)

const version = "v1.0.0"

func main() {
	var versionCalled bool
	flag.BoolVar(&versionCalled, "v", false, "show version")
	flag.BoolVar(&versionCalled, "version", false, "show version")

	cfg, err := config.New()

	if versionCalled {
		fmt.Println(version)
		return
	}

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
