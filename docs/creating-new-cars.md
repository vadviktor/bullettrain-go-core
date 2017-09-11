# How to create a new car

This document will describe the process to add new car plugin in Go.

## Naming

The car needs to have a fitting package name that matches the current convention: `carSomeName`
The name is prefixed by `car`.

## Template:

Here is a demo template you can use to start off your new `car`. It shows basic implementation of the `carRenderer` interface.

```go
package carDemo

import (
	"fmt"
	"os"

	"github.com/bullettrain-sh/bullettrain-go-core/ansi"
)

const (
    carPaint = "black:white"
    symbolIcon = "€"
    symbolPaint = "black:white"
)

// Demo Car
type Car struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_DEMO_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

func paintedSymbol() string {
	var demoSymbol string
	if demoSymbol = os.Getenv("BULLETTRAIN_CAR_DEMO_SYMBOL_ICON"); demoSymbol == "" {
		demoSymbol = symbolIcon
	}

	var timeSymbolPaint string
	if timeSymbolPaint = os.Getenv("BULLETTRAIN_CAR_DEMO_SYMBOL_PAINT"); timeSymbolPaint == "" {
		timeSymbolPaint = symbolPaint
	}

	return ansi.Color(demoSymbol, timeSymbolPaint)
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	s := false
	if e := os.Getenv("BULLETTRAIN_CAR_DEMO_SHOW"); e == "true" {
		s = true
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out)
	carPaint := ansi.ColorFunc(c.GetPaint())


	out <- fmc.Sprintf("%s%s", paintedSymbol(), carPaint("Demo text"))
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_DEMO_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_DEMO_SEPARATOR_SYMBOL")
}
```


When you may want to just permanently change symbol or paint colours, you can simply change the constants on the top level and compile your custom version:

```go
const (
    carPaint = "black:white"
    symbolIcon = "€"
    symbolPaint = "black:white"
)
```

To use your car, this is a checklist to be done in `bullettrain-go-core/bullettrain.go`:

* TODO


## Documenting cars

The `README` should also be as detailed as the ones already existing.
Each and every env variable must be documented with default values and description.
