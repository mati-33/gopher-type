package textproviders

import "slices"

type Provider interface {
	Provide(wordCount int) []rune
}

type entry struct {
	Priority int
	Factory  func() Provider
}

var registry = map[string]entry{
	"english":           {1, newEn1kProvider},
	"polish":            {10, newPl2kProvider},
	"polish diacritics": {20, newPlDiacriticsProvider},
	"numbers":           {30, newNumberProvider},
}

func GetProviderNames() []string {
	names := make([]string, 0, len(registry))

	for n := range registry {
		names = append(names, n)
	}

	slices.SortFunc(names, func(a, b string) int {
		return registry[a].Priority - registry[b].Priority
	})

	return names
}

func MustGetProvider(name string) Provider {
	entry, ok := registry[name]
	if !ok {
		panic("no provider with that name")
	}

	return entry.Factory()
}
