package config

type Config struct {
	InitTheme     string
	InitMode      string
	InitWordCount int
	PreviewSize   int
	Transparent   bool

	SpeedIcon     string
	AccuracyIncon string
	ModeIcon      string
	WordCountIcon string
	PreviewIcon   string
}

func NewDefault() *Config {
	return &Config{
		InitTheme:     "gopher type",
		InitMode:      "english",
		InitWordCount: 15,
		PreviewSize:   15,
		Transparent:   false,
		SpeedIcon:     "󱐋",
		AccuracyIncon: "󰣉",
		ModeIcon:      "󰦨",
		WordCountIcon: "",
		PreviewIcon:   "󱎸",
	}
}
