package asciiart

import (
	_ "embed"

	asciiart "github.com/romance-dev/ascii-art"
)

//go:embed font.flf
var font []byte

func init() {
	asciiart.RegisterFont("calligraphy", font)
}
