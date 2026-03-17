package modes

import (
	"fmt"
	"math/rand"
	"strings"
)

type numberMode struct {
	max int
}

func (p numberMode) Generate(wordCount int) []rune {
	b := strings.Builder{}

	for i := range wordCount {
		fmt.Fprintf(&b, "%d", p.genNum())

		if i < wordCount-1 {
			b.WriteString(" ")
		}
	}

	return []rune(b.String())
}

func (p numberMode) Name() string {
	return "numbers"
}

func (p numberMode) Preview() string {
	return "1983 421 8723 668 8524 75 49 334 50 33 655 349 4030 94 8 59 141 6721 1801 6080"
}

func (p numberMode) genNum() int {
	x := rand.Intn(4)
	num := 0
	switch x {
	case 0:
		num = rand.Intn(10)
	case 1:
		num = rand.Intn(90) + 10
	case 2:
		num = rand.Intn(900) + 100
	case 3:
		num = rand.Intn(9000) + 1000
	}
	return num
}

func newNumberMode() Mode {
	return numberMode{
		max: 1000,
	}
}
