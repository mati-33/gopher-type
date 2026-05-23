package config

import (
	"cmp"
	"errors"
)

var ErrConfigNotFound = errors.New("config file not found")

var (
	filename = "config.json"
	filedir  = "gopher-type"
	envvar   = "GOPHER_TYPE_CONFIG"
)

type Config struct {
	InitTheme     string
	InitMode      string
	InitWordCount int
	PreviewSize   int
	Transparent   bool
	Icons         Icons
}

type Icons struct {
	Speed     string
	Accuracy  string
	Mode      string
	WordCount string
	Preview   string
	Theme     string
}

func newIcons() Icons {
	return Icons{
		Speed:     "󱐋",
		Accuracy:  "󰣉",
		Mode:      "󰦨",
		WordCount: "",
		Preview:   "󱎸",
		Theme:     "",
	}
}

func newEmptyIcons() Icons {
	return Icons{
		Speed:     " ",
		Accuracy:  " ",
		Mode:      " ",
		WordCount: " ",
		Preview:   " ",
		Theme:     " ",
	}
}

func newDefault() *Config {
	return &Config{
		InitTheme:     "gopher type",
		InitMode:      "english",
		InitWordCount: 15,
		PreviewSize:   15,
		Transparent:   false,
		Icons:         newIcons(),
	}
}

func New(fc *fileConfig) *Config {
	c := newDefault()
	if fc == nil {
		return c
	}

	c.InitMode = cmp.Or(fc.Mode, c.InitMode)
	c.InitTheme = cmp.Or(fc.Theme, c.InitTheme)

	if fc.Transparent != nil {
		c.Transparent = *fc.Transparent
	}

	if fc.Icons != nil && !*fc.Icons {
		c.Icons = newEmptyIcons()
	}

	return c
}
