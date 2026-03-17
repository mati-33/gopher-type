package config

type Config struct {
	InitMode      string
	InitWordCount int
	PreviewSize   int
}

func NewDefault() Config {
	return Config{
		InitMode:      "english",
		InitWordCount: 15,
		PreviewSize:   15,
	}
}
