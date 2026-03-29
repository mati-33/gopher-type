package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/mati-33/gopher-type/internal/app"
)

func main() {
	f, err := tea.LogToFile("debug.log", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup debug logging: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	app := app.New()
	p := tea.NewProgram(app)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start the program: %v", err)
		os.Exit(1)
	}
}
