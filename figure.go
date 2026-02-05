package figure

import (
	"log"
	"reflect"
	"strings"
)

const (
	ascii_offset = 32
	first_ascii  = ' '
	last_ascii   = '~'
)

type Figure struct {
	phrase string
	font
	strict bool
	color  string
}

// NewFigure converts a phrase to a Figure that can then
// be printed to an io.Writer (such as io.Stdout).
//
// font must be loaded beforehand. The default is "standard"
// (provided it is already loaded).
//
// If strict mode is enabled, then any unprintable ASCII
// characters will produce a panic. Otherwise, it will be
// replaced by a '?'.
func NewFigure(phrase, font string, strict bool) Figure {
	f := newFont(font)
	if f.reverse {
		phrase = reverse(phrase)
	}
	return Figure{phrase: phrase, font: f, strict: strict}
}

// NewColorFigure converts a phrase to a Figure that can then
// be printed to an io.Writer (such as io.Stdout).
//
// font must be loaded beforehand. The default is "standard"
// (provided it is already loaded).
//
// color can be Red, Green, Yellow, Blue, Purple, Cyan, Gray or White.
//
// If strict mode is enabled, then any unprintable ASCII
// characters will produce a panic. Otherwise, it will be
// replaced by a '?'.
func NewColorFigure(phrase, font string, color string, strict bool) Figure {
	fig := NewFigure(phrase, font, strict)
	if color != Black {
		color = strings.ToLower(color)
		if _, found := colors[color]; !found {
			log.Fatalf("invalid color. must be one of: %s", reflect.ValueOf(colors).MapKeys())
		}
		fig.color = color
	}
	return fig
}

// Slicify returns a slice of the figure.
func (figure Figure) Slicify() (rows []string) {
	for r := 0; r < figure.font.height; r++ {
		var printRow string
		for _, char := range figure.phrase {
			if char < first_ascii || char > last_ascii {
				if figure.strict {
					log.Fatal("invalid input.")
				} else {
					char = '?'
				}
			}
			fontIndex := char - ascii_offset
			charRowText := scrub(figure.font.letters[fontIndex][r], figure.font.hardblank)
			printRow += charRowText
		}
		if r < figure.font.baseline || len(strings.TrimSpace(printRow)) > 0 {
			rows = append(rows, strings.TrimRight(printRow, " "))
		}
	}
	return rows
}

func scrub(text string, char byte) string {
	return strings.Replace(text, string(char), " ", -1)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
