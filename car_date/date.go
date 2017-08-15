package carDate

import (
	"fmt"
	"os"
	"time"

	"github.com/mgutz/ansi"
)

const carPaint = "white:black"
const symbolIcon = ""
const symbolPaint = "red:black"

// Date car
type Date struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (t *Date) GetPaint() string {
	if t.paint = os.Getenv("BULLETTRAIN_CAR_DATE_PAINT"); t.paint == "" {
		t.paint = carPaint
	}

	return t.paint
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
func (t *Date) CanShow() bool {
	s := false
	if e := os.Getenv("BULLETTRAIN_CAR_DATE_SHOW"); e == "true" {
		s = true
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (t *Date) Render(out chan<- string) {
	defer close(out)
	carPaint := ansi.ColorFunc(t.GetPaint())

	y, m, d := time.Now().Date()
	out <- fmt.Sprintf("%s%s",
		paintedSymbol(),
		carPaint(fmt.Sprintf("%02d-%02d-%02d", y, m, d)))
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (t *Date) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_DATE_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables
func (t *Date) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_DATE_SEPARATOR_SYMBOL")
}
