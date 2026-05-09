package config

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"

	"os"
	"slices"

	"github.com/mati-33/gopher-type/internal/modes"
	"github.com/mati-33/gopher-type/internal/themes"
)

var ErrConfigNotFound = errors.New("config file not found")

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

type userConfig struct {
	Theme       string `json:"theme"`
	Mode        string `json:"mode"`
	Transparent *bool  `json:"transparent"`
	Icons       *bool  `json:"icons"`
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

func New(userConfig *userConfig) *Config {
	c := newDefault()
	if userConfig == nil {
		return c
	}

	c.InitMode = cmp.Or(userConfig.Mode, c.InitMode)
	c.InitTheme = cmp.Or(userConfig.Theme, c.InitTheme)

	if userConfig.Transparent != nil {
		c.Transparent = *userConfig.Transparent
	}

	if userConfig.Icons != nil && !*userConfig.Icons {
		c.Icons = newEmptyIcons()
	}

	return c
}

func LoadUserConfig() (*userConfig, error) {
	fileContent, err := loadUserConfigFile()
	if err != nil {
		return nil, err
	}

	var uc userConfig
	if err := json.Unmarshal(fileContent, &uc); err != nil {
		return nil, err
	}

	modes := modes.GetModeNames()
	if uc.Mode != "" && !slices.Contains(modes, uc.Mode) {
		return nil, fmt.Errorf("invalid mode: %s", uc.Mode)
	}

	themes := themes.GetThemeNames()
	if uc.Theme != "" && !slices.Contains(themes, uc.Theme) {
		return nil, fmt.Errorf("ivalid theme: %s", uc.Theme)
	}

	return &uc, nil
}

func loadUserConfigFile() ([]byte, error) {
	var (
		content []byte
		err     error
	)

	for _, path := range []string{
		os.Getenv("GOPHER_CONFIG"),
		fmt.Sprintf("%s/gopher-type/config.json", os.Getenv("XDG_CONFIG_HOME")),
		fmt.Sprintf("%s/.config/gopher-type/config.json", os.Getenv("HOME")),
	} {
		content, err = os.ReadFile(path)
		if err == nil {
			return content, nil
		}
	}

	return nil, ErrConfigNotFound
}
