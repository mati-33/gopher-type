package textproviders

type Provider interface {
	Provide(wordCount int) []rune
	Name() string
	Preview() string
}

var registry = []Provider{
	newEn1kProvider(),
	newPl2kProvider(),
	newPlDiacriticsProvider(),
	newNumberProvider(),
}

func Providers() []Provider {
	return registry
}

func GetProviderNames() []string {
	names := make([]string, 0, len(registry))

	for _, p := range registry {
		names = append(names, p.Name())
	}

	return names
}

func MustGetProvider(name string) Provider {
	for _, p := range registry {
		if name == p.Name() {
			return p
		}
	}
	panic("no provider with that name")
}
