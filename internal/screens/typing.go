package screens

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	comp "github.com/mati-33/gopher-type/internal/components"
	"github.com/mati-33/gopher-type/internal/config"
	"github.com/mati-33/gopher-type/internal/modes"
	"github.com/mati-33/gopher-type/internal/themes"
)

type typingScreen struct {
	config         config.Config
	theme          themes.Theme
	keybinds       typingKeybinds
	mode           modes.Mode
	width          int
	height         int
	wordCount      int
	text           comp.Text
	help           comp.Help
	speedField     comp.DetailField
	accuracyField  comp.DetailField
	modeField      comp.DetailField
	wordCountField comp.DetailField
}

func NewTypingScreen(config config.Config, theme themes.Theme, width, height int) typingScreen {
	mode := modes.MustGetMode(config.InitMode)
	wc := config.InitWordCount
	keybinds := newTypingKeybinds()

	return typingScreen{
		mode:      mode,
		width:     width,
		height:    height,
		wordCount: config.InitWordCount,
		text:      comp.NewText(theme, mode.Generate(wc), int(float32(width)*0.7), height),
		config:    config,
		theme:     theme,
		keybinds:  keybinds,
		help: comp.NewHelp(theme, []comp.Keybind{
			keybinds.IncWordCount,
			keybinds.DecWordCount,
			keybinds.ChangeMode,
			keybinds.ChangeTheme,
			keybinds.GoBack,
			keybinds.Help,
		}),
		speedField:     comp.NewDetailField(theme, "speed", config.SpeedIcon, "-"),
		accuracyField:  comp.NewDetailField(theme, "accuracy", config.AccuracyIncon, "-"),
		modeField:      comp.NewDetailField(theme, "mode", config.ModeIcon, mode.Name()),
		wordCountField: comp.NewDetailField(theme, "word count", config.WordCountIcon, "15"),
	}
}

func (s *typingScreen) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
		s.text.Width = int(float32(s.width) * 0.7)

	case comp.TextResult:
		s.speedField.Value = fmt.Sprintf("%dwpm", msg.Wpm)
		s.accuracyField.Value = fmt.Sprintf("%.2f%%", 100.0*msg.Accuracy)
		s.text.Text = s.mode.Generate(s.wordCount)

	case themes.Theme:
		s.theme = msg

	case ChangeProvider:
		s.mode = modes.MustGetMode(msg.Name)
		s.modeField.Value = s.mode.Name()
		s.text.Text = s.mode.Generate(s.wordCount)
		return tea.Batch(s.text.Reset()...)

	case tea.KeyMsg:
		switch msg.String() {

		case s.keybinds.GoBack.Key:
			if s.text.Started {
				s.text.Text = s.mode.Generate(s.wordCount)
				return tea.Batch(s.text.Reset()...)
			} else {
				return popScreen(nil)
			}

		case s.keybinds.IncWordCount.Key:
			s.wordCount++
			s.wordCountField.Value = fmt.Sprintf("%d", s.wordCount)
			s.text.Text = s.mode.Generate(s.wordCount)
			return tea.Batch(s.text.Reset()...)

		case s.keybinds.DecWordCount.Key:
			if s.wordCount > 1 {
				s.wordCount--
				s.wordCountField.Value = fmt.Sprintf("%d", s.wordCount)
				s.text.Text = s.mode.Generate(s.wordCount)
				return tea.Batch(s.text.Reset()...)
			} else {
				return nil
			}

		case s.keybinds.ChangeMode.Key:
			screen := NewModeScreen(s.config, s.theme, s.width, s.height)
			return pushScreen(&screen)

		case s.keybinds.ChangeTheme.Key:
			screen := NewThemeChangeScreen(s.config, s.theme)
			return pushScreen(&screen)

		case s.keybinds.Help.Key:
			s.help.Toggle()
			return nil
		}
	}

	cmds = append(cmds,
		s.text.Update(msg),
		s.speedField.Update(msg),
		s.accuracyField.Update(msg),
		s.modeField.Update(msg),
		s.wordCountField.Update(msg),
		s.help.Update(msg),
	)

	return tea.Batch(cmds...)
}

func (s typingScreen) View() tea.View {
	bannerOffset := int(float32(s.height) * 0.2)

	if s.height < 14 {
		bannerOffset = 0
	}

	helpView := s.help.View()
	helpOffset := max(0, s.width-lipgloss.Width(helpView)-2)

	bannerView := lipgloss.PlaceHorizontal(
		s.width,
		lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Top,
			s.speedField.View(),
			"  ",
			s.accuracyField.View(),
			"     ",
			s.modeField.View(),
			"  ",
			s.wordCountField.View(),
		),
	)

	textView := s.text.View()
	textLayer := lipgloss.NewLayer(lipgloss.Place(
		s.width, s.height,
		lipgloss.Center, lipgloss.Center,
		textView,
	),
		lipgloss.NewLayer(bannerView).Y(bannerOffset),
		lipgloss.NewLayer(helpView).Y(s.height-lipgloss.Height(helpView)).X(helpOffset),
	)

	c := lipgloss.NewCompositor(textLayer)
	return tea.NewView(c.Render())
}

type typingKeybinds struct {
	IncWordCount comp.Keybind
	DecWordCount comp.Keybind
	ChangeMode   comp.Keybind
	ChangeTheme  comp.Keybind
	GoBack       comp.Keybind
	Help         comp.Keybind
}

func newTypingKeybinds() typingKeybinds {
	return typingKeybinds{
		IncWordCount: comp.Keybind{Key: "ctrl+o", Desc: "increase word count"},
		DecWordCount: comp.Keybind{Key: "ctrl+p", Desc: "decrease word count"},
		ChangeMode:   comp.Keybind{Key: "ctrl+n", Desc: "change mode"},
		ChangeTheme:  comp.Keybind{Key: "ctrl+t", Desc: "change theme"},
		GoBack:       comp.Keybind{Key: "esc", Desc: "go back"},
		Help:         comp.Keybind{Key: "f1", Desc: "close help"},
	}
}
