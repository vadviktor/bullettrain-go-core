package carTime

import (
	"fmt"
	"os"
	"time"

	"github.com/mgutz/ansi"
)

type Time struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (t *Time) GetPaint() string {
	if t.paint = os.Getenv("BULLETTRAIN_CAR_TIME_PAINT"); t.paint == "" {
		t.paint = "white:black"
	}

	return t.paint
}

func paintedSymbol() string {
	var symbol string
	if symbol = os.Getenv("BULLETTRAIN_CAR_TIME_SYMBOL_ICON"); symbol == "" {
		symbol = "  "
	}

	var symbolPaint string
	if symbolPaint = os.Getenv("BULLETTRAIN_CAR_TIME_SYMBOL_PAINT"); symbolPaint == "" {
		symbolPaint = "white:black"
	}

	return ansi.Color(symbol, symbolPaint)
}

// CanShow decides if this car needs to be displayed.
func (t *Time) CanShow() bool {
	s := true
	if e := os.Getenv("BULLETTRAIN_CAR_TIME_SHOW"); e == "false" {
		s = false
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (t *Time) Render(out chan<- string) {
	defer close(out)
	carPaint := ansi.ColorFunc(t.GetPaint())

	n := time.Now()
	out <- fmt.Sprintf("%s%s%s",
		paintedSymbol(),
		carPaint(fmt.Sprintf("%02d:%02d:%02d",
			n.Hour(), n.Minute(), n.Second())),
		carPaint(" "))
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (t *Time) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_TIME_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (t *Time) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_TIME_SEPARATOR_SYMBOL")
}
