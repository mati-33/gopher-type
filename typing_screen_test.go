package main

import (
	"fmt"
	"slices"
	"strings"
	"testing"
	"time"
)

func Test_calculateWpm(t *testing.T) {
	tests := []struct {
		name    string
		runesNo int
		elapsed time.Duration
		want    int
	}{
		{
			name:    "case1",
			runesNo: 50,
			elapsed: time.Second * 60,
			want:    10,
		},
		{
			name:    "case2",
			runesNo: 250,
			elapsed: time.Second * 33,
			want:    91,
		},
		{
			name:    "case3",
			runesNo: 133,
			elapsed: time.Second * 66,
			want:    24,
		},
		{
			name:    "case4",
			runesNo: 200,
			elapsed: time.Second * 5,
			want:    480,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateWpm(tt.runesNo, tt.elapsed)
			if got != tt.want {
				t.Errorf("calculateWpm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateAccuracy(t *testing.T) {
	tests := []struct {
		name     string
		runesNo  int
		errorsNo int
		want     float64
	}{
		{
			name:     "case1",
			runesNo:  100,
			errorsNo: 1,
			want:     0.99,
		},
		{
			name:     "case2",
			runesNo:  100,
			errorsNo: 0,
			want:     1.0,
		},
		{
			name:     "case3",
			runesNo:  100,
			errorsNo: 10,
			want:     0.9,
		},
		{
			name:     "case4",
			runesNo:  133,
			errorsNo: 5,
			want:     0.9624,
		},
		{
			name:     "case5",
			runesNo:  456,
			errorsNo: 59,
			want:     0.8706,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateAccuracy(tt.runesNo, tt.errorsNo)
			if got != tt.want {
				t.Errorf("calculateAccuracy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitText(t *testing.T) {
	tests := []struct {
		name  string
		text  []rune
		width int
		want  [][]rune
	}{
		{
			name:  "case1",
			text:  []rune("gave huge nature has trade as moment modern much did learn catch wonder office offer give anger lift atom cause control here fast sharp example"),
			width: 57,
			want: [][]rune{
				[]rune("gave huge nature has trade as moment modern much did "),
				[]rune("learn catch wonder office offer give anger lift atom "),
				[]rune("cause control here fast sharp example "),
			},
		},
		{
			name:  "avoid infinite loop",
			text:  []rune("foo bar baz"),
			width: 0,
			want: [][]rune{
				[]rune("foo "),
				[]rune("bar "),
				[]rune("baz "),
			},
		},
		{
			name:  "big width",
			text:  []rune("foo bar baz"),
			width: 110,
			want: [][]rune{
				[]rune("foo bar baz "),
			},
		},
		{
			name:  "two lines",
			text:  []rune("hello world foo bar golang gopher python snek"),
			width: 28,
			want: [][]rune{
				[]rune("hello world foo bar golang "),
				[]rune("gopher python snek "),
			},
		},
		{
			name:  "no text",
			text:  []rune(""),
			width: 10,
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitText(tt.text, tt.width)
			if !runeSlicesEqual(got, tt.want) {
				t.Errorf("splitText()\ncase: %s\ngot: %s\nwant: %s", tt.name, printSplittedText(got), printSplittedText(tt.want))
			}
		})
	}
}

func printSplittedText(splitted [][]rune) string {
	if len(splitted) == 0 {
		return "[][]rune{}"
	}

	b := strings.Builder{}
	b.WriteString("[][]rune{\n")

	for _, line := range splitted {
		fmt.Fprintf(&b, "  []rune(\"%s\"),\n", string(line))
	}

	b.WriteString("}\n")

	return b.String()
}

func runeSlicesEqual(a, b [][]rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !slices.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}
