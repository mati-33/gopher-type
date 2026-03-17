package modes

type Mode interface {
	Generate(wordCount int) []rune
	Name() string
	Preview() string
}

var registry = []Mode{
	newEn1kMode(),
	newPl2kMode(),
	newPlDiacriticsMode(),
	newNumberMode(),
}

func Modes() []Mode {
	return registry
}

func GetModeNames() []string {
	names := make([]string, 0, len(registry))

	for _, p := range registry {
		names = append(names, p.Name())
	}

	return names
}

func MustGetMode(name string) Mode {
	for _, p := range registry {
		if name == p.Name() {
			return p
		}
	}
	panic("no mode with that name")
}
