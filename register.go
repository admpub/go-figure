package figure

import "sync"

var (
	driversMu sync.RWMutex
	fonts     = make(map[string][]byte)
)

// RegisterFont registers a new font.
// Builtin fonts can be blank imported from the
// fonts sub-package individually or en masse.
func RegisterFont(name string, font []byte) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if font == nil {
		panic("ascii-art/plugin: Register font is nil")
	}
	if _, dup := fonts[name]; dup {
		panic("ascii-art/plugin: Register called twice for font " + name)
	}
	fonts[name] = font
}
