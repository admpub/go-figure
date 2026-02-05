package figure_test

import (
	"testing"

	"github.com/admpub/go-figure"
	_ "github.com/admpub/go-figure/fonts/big"
)

func TestPrint(t *testing.T) {
	myFigure := figure.NewFigure("Hello World", "big", false)
	myFigure.Print()
}
