# Details to how to create a new car

Here is a template:

```go
package carDemo

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

const carPaint = "black:white"
const symbolIcon = "€"
const symbolPaint = "black:white"

// Demo Car
type Demo struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (t *Demo) GetPaint() string {
	if t.paint = os.Getenv("BULLETTRAIN_CAR_DEMO_PAINT"); t.paint == "" {
		t.paint = carPaint
	}

	return t.paint
}

func paintedSymbol() string {
	var timeSymbol string
	if timeSymbol = os.Getenv("BULLETTRAIN_CAR_DEMO_SYMBOL_ICON"); timeSymbol == "" {
		timeSymbol = symbolIcon
	}

	var timeSymbolPaint string
	if timeSymbolPaint = os.Getenv("BULLETTRAIN_CAR_DEMO_SYMBOL_PAINT"); timeSymbolPaint == "" {
		timeSymbolPaint = symbolPaint
	}

	return ansi.Color(timeSymbol, timeSymbolPaint)
}

// CanShow decides if this car needs to be displayed.
func (t *Demo) CanShow() bool {
	s := false
	if e := os.Getenv("BULLETTRAIN_CAR_DEMO_SHOW"); e == "true" {
		s = true
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (t *Demo) Render(out chan<- string) {
	defer close(out)
	carPaint := ansi.ColorFunc(t.GetPaint())

	
	out <- fmt.Sprintf("%s%s", paintedSymbol(), carPaint("Demo text"))
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (t *Demo) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_DEMO_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (t *Demo) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_DEMO_SEPARATOR_SYMBOL")
}
```
