package carDate

import (
	"fmt"
	"os"
	"time"

	"github.com/bullettrain-sh/bullettrain-go-core/ansi"
)

const (
	carPaint    = "white:black"
	symbolIcon  = " "
	symbolPaint = "red:black"
)

// Date car
type Car struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_DATE_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

func paintedSymbol() string {
	var dateSymbol string
	if dateSymbol = os.Getenv("BULLETTRAIN_CAR_DATE_SYMBOL_ICON"); dateSymbol == "" {
		dateSymbol = symbolIcon
	}

	var dateSymbolPaint string
	if dateSymbolPaint = os.Getenv("BULLETTRAIN_CAR_DATE_SYMBOL_PAINT"); dateSymbolPaint == "" {
		dateSymbolPaint = symbolPaint
	}

	return ansi.Color(dateSymbol, dateSymbolPaint)
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	s := false
	if e := os.Getenv("BULLETTRAIN_CAR_DATE_SHOW"); e == "true" {
		s = true
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out)
	carPaint := ansi.ColorFunc(c.GetPaint())

	y, m, d := time.Now().Date()
	out <- fmt.Sprintf("%s%s",
		paintedSymbol(),
		carPaint(fmt.Sprintf("%02d-%02d-%02d", y, m, d)))
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_DATE_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_DATE_SEPARATOR_SYMBOL")
}
