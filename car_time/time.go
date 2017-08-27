package carTime

import (
	"fmt"
	"os"
	"time"

	"github.com/bullettrain-sh/bullettrain-go-core/ansi"
)

const (
	carPaint    = "black:white"
	symbolIcon  = ""
	symbolPaint = "black:white"
)

// Time Car
type Car struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_TIME_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

func paintedSymbol() string {
	var timeSymbol string
	if timeSymbol = os.Getenv("BULLETTRAIN_CAR_TIME_SYMBOL_ICON"); timeSymbol == "" {
		timeSymbol = symbolIcon
	}

	var timeSymbolPaint string
	if timeSymbolPaint = os.Getenv("BULLETTRAIN_CAR_TIME_SYMBOL_PAINT"); timeSymbolPaint == "" {
		timeSymbolPaint = symbolPaint
	}

	return ansi.Color(timeSymbol, timeSymbolPaint)
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	s := false
	if e := os.Getenv("BULLETTRAIN_CAR_TIME_SHOW"); e == "true" {
		s = true
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out)
	carPaint := ansi.ColorFunc(c.GetPaint())

	n := time.Now()
	out <- fmt.Sprintf("%s%s",
		paintedSymbol(),
		carPaint(fmt.Sprintf("%02d:%02d:%02d",
			n.Hour(), n.Minute(), n.Second())))
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_TIME_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_TIME_SEPARATOR_SYMBOL")
}
