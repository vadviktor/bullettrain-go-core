package carStatus

import (
	"fmt"
	"os"

	"github.com/bullettrain-sh/bullettrain-go-core/ansi"
)

const (
	carPaint    = "white:red"
	symbolIcon  = ""
	symbolPaint = "yellow:red"
)

// Status Car
type Car struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_STATUS_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
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
func (c *Car) CanShow() bool {
	if len(os.Args) > 1 {
		return os.Args[1] != "" && os.Args[1] != "0"
	}

	return false
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out)
	carPaint := ansi.ColorFunc(c.GetPaint())

	if n := os.Getenv("BULLETTRAIN_CAR_STATUS_CODE_SHOW"); n == "false" {
		out <- fmt.Sprintf("%s", paintedSymbol())
	} else {
		out <- fmt.Sprintf("%s%s", paintedSymbol(), carPaint(os.Args[1]))
	}
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_STATUS_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_STATUS_SEPARATOR_SYMBOL")
}
