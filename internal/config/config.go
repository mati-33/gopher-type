package config

type Config struct {
	InitMode      string
	InitWordCount int
	PreviewSize   int
	Transparent   bool
}

func NewDefault() Config {
	return Config{
		InitMode:      "english",
		InitWordCount: 15,
		PreviewSize:   15,
		Transparent:   false,
	}
}
