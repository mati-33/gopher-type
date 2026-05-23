package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"slices"

	"github.com/mati-33/gopher-type/internal/modes"
	"github.com/mati-33/gopher-type/internal/themes"
)

var errConfigNotFound = errors.New("config file not found")

type fileConfig struct {
	Theme       string `json:"theme"`
	Mode        string `json:"mode"`
	Transparent *bool  `json:"transparent"`
	Icons       *bool  `json:"icons"`
}

func parseFileConfig() (*fileConfig, error) {
	fileContent, err := loadConfigFile()
	if err != nil {
		return nil, err
	}

	var uc fileConfig
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

func loadConfigFile() ([]byte, error) {
	fp, err := getConfigFilepath()
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(fp)
	if err != nil {
		return nil, errConfigNotFound
	}

	return content, nil
}
