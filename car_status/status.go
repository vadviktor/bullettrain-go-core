package carStatus

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

const carPaint = "white:red"
const symbolIcon = ""
const symbolPaint = "yellow:red"

// Status Car
type Status struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (t *Status) GetPaint() string {
	if t.paint = os.Getenv("BULLETTRAIN_CAR_STATUS_PAINT"); t.paint == "" {
		t.paint = carPaint
	}

	return t.paint
}

func paintedSymbol() string {
	var timeSymbol string
	if timeSymbol = os.Getenv("BULLETTRAIN_CAR_STATUS_SYMBOL_ICON"); timeSymbol == "" {
		timeSymbol = symbolIcon
	}

	var timeSymbolPaint string
	if timeSymbolPaint = os.Getenv("BULLETTRAIN_CAR_STATUS_SYMBOL_PAINT"); timeSymbolPaint == "" {
		timeSymbolPaint = symbolPaint
	}

	return ansi.Color(timeSymbol, timeSymbolPaint)
}

// CanShow decides if this car needs to be displayed.
func (t *Status) CanShow() bool {
	if len(os.Args) > 1 {
		return os.Args[1] != "" && os.Args[1] != "0"
	}

	return false
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (t *Status) Render(out chan<- string) {
	defer close(out)
	carPaint := ansi.ColorFunc(t.GetPaint())

	out <- fmt.Sprintf("%s%s",
		paintedSymbol(),
		carPaint(os.Args[1]))
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (t *Status) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_STATUS_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (t *Status) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_STATUS_SEPARATOR_SYMBOL")
}
