package asciiart

import (
	_ "embed"

	"github.com/admpub/go-figure"
)

//go:embed font.flf
var font []byte

func init() {
	figure.RegisterFont("hollywood", font)
}
