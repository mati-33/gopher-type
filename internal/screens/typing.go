package screens

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mati-33/gopher-type/internal/appcontex"
	comp "github.com/mati-33/gopher-type/internal/components"
	"github.com/mati-33/gopher-type/internal/modes"
)

type typing struct {
	ctx            *appcontex.AppContext
	text           comp.Text
	help           comp.Help
	speedField     comp.DetailField
	accuracyField  comp.DetailField
	modeField      comp.DetailField
	wordCountField comp.DetailField
	wordCount      int
	keybinds       typingKeybinds
}

func NewTyping(ctx *appcontex.AppContext) *typing {
	wc := ctx.Config.InitWordCount
	keybinds := newTypingKeybinds()

	return &typing{
		ctx:  ctx,
		text: comp.NewText(ctx.Theme, ctx.Mode.Generate(wc), int(float32(ctx.Width)*0.7), ctx.Height),
		help: comp.NewHelp(ctx.Theme, []comp.Keybind{
			keybinds.IncWordCount,
			keybinds.DecWordCount,
			keybinds.ChangeMode,
			keybinds.ChangeTheme,
			keybinds.GoBack,
			keybinds.Help,
		}),
		speedField:     comp.NewDetailField(ctx.Theme, "speed", ctx.Config.SpeedIcon, "-"),
		accuracyField:  comp.NewDetailField(ctx.Theme, "accuracy", ctx.Config.AccuracyIncon, "-"),
		modeField:      comp.NewDetailField(ctx.Theme, "mode", ctx.Config.ModeIcon, ctx.Mode.Name()),
		wordCountField: comp.NewDetailField(ctx.Theme, "word count", ctx.Config.WordCountIcon, "15"),
		wordCount:      ctx.Config.InitWordCount,
		keybinds:       keybinds,
	}
}

func (s *typing) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		s.text.Width = int(float32(s.ctx.Width) * 0.7)

	case comp.TextResult:
		s.speedField.Value = fmt.Sprintf("%dwpm", msg.Wpm)
		s.accuracyField.Value = fmt.Sprintf("%.2f%%", 100.0*msg.Accuracy)
		s.text.Text = s.ctx.Mode.Generate(s.wordCount)

	case ChangeProvider:
		s.ctx.Mode = modes.MustGetMode(msg.Name)
		s.modeField.Value = s.ctx.Mode.Name()
		s.text.Text = s.ctx.Mode.Generate(s.wordCount)
		s.text.Reset()
		return nil

	case tea.KeyMsg:
		switch msg.String() {

		case s.keybinds.GoBack.Key:
			if s.text.Started {
				s.text.Text = s.ctx.Mode.Generate(s.wordCount)
				s.text.Reset()
				return nil
			} else {
				return pop(nil)
			}

		case s.keybinds.IncWordCount.Key:
			s.wordCount++
			s.wordCountField.Value = fmt.Sprintf("%d", s.wordCount)
			s.text.Text = s.ctx.Mode.Generate(s.wordCount)
			s.text.Reset()
			return nil

		case s.keybinds.DecWordCount.Key:
			if s.wordCount > 1 {
				s.wordCount--
				s.wordCountField.Value = fmt.Sprintf("%d", s.wordCount)
				s.text.Text = s.ctx.Mode.Generate(s.wordCount)
				s.text.Reset()
				return nil
			} else {
				return nil
			}
		case s.keybinds.ChangeMode.Key:
			return push(NewModeChange(s.ctx))

		case s.keybinds.ChangeTheme.Key:
			return push(NewThemeChange(s.ctx))

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

func (s typing) View() string {
	bannerOffset := int(float32(s.ctx.Height) * 0.2)

	if s.ctx.Height < 14 {
		bannerOffset = 0
	}

	helpView := s.help.View()
	helpOffset := max(0, s.ctx.Width-lipgloss.Width(helpView)-2)

	bannerView := lipgloss.PlaceHorizontal(
		s.ctx.Width,
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
		s.ctx.Width, s.ctx.Height,
		lipgloss.Center, lipgloss.Center,
		textView,
	),
		lipgloss.NewLayer(bannerView).Y(bannerOffset),
		lipgloss.NewLayer(helpView).Y(s.ctx.Height-lipgloss.Height(helpView)).X(helpOffset),
	)

	c := lipgloss.NewCompositor(textLayer)
	return c.Render()
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
