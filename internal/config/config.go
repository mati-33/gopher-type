package config

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"slices"

	"github.com/mati-33/gopher-type/internal/modes"
	"github.com/mati-33/gopher-type/internal/themes"
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

func getConfigFilepath() (string, error) {
	if gtc := os.Getenv(envvar); gtc != "" {
		return path.Join(gtc, filename), nil
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return path.Join(configDir, filedir, filename), nil

}

func loadUserConfigFile() ([]byte, error) {
	fp, err := getConfigFilepath()
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(fp)
	if err != nil {
		return nil, ErrConfigNotFound
	}

	return content, nil
}
