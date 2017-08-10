package carDate

import (
	"fmt"
	"os"
	"time"

	"github.com/mgutz/ansi"
)

type Date struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (t *Date) GetPaint() string {
	if t.paint = os.Getenv("BULLETTRAIN_CAR_DATE_PAINT"); t.paint == "" {
		t.paint = "white:black"
	}

	return t.paint
}

func paintedSymbol() string {
	var symbol string
	if symbol = os.Getenv("BULLETTRAIN_CAR_DATE_SYMBOL_ICON"); symbol == "" {
		symbol = "  "
	}

	var symbolPaint string
	if symbolPaint = os.Getenv("BULLETTRAIN_CAR_DATE_SYMBOL_PAINT"); symbolPaint == "" {
		symbolPaint = "red:black"
	}

	return ansi.Color(symbol, symbolPaint)
}

// CanShow decides if this car needs to be displayed.
func (t *Date) CanShow() bool {
	s := true
	if e := os.Getenv("BULLETTRAIN_CAR_DATE_SHOW"); e == "false" {
		s = false
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (t *Date) Render(out chan<- string) {
	defer close(out)
	carPaint := ansi.ColorFunc(t.GetPaint())

	y, m, d := time.Now().Date()
	out <- fmt.Sprintf("%s%s%s",
		paintedSymbol(),
		carPaint(fmt.Sprintf("%02d-%02d-%02d", y, m, d)),
		carPaint(" "))
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
