package carContext

import (
	"fmt"
	"os"
	"os/user"

	"github.com/mgutz/ansi"
)

// Context car
type Context struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (t *Context) GetPaint() string {
	if t.paint = os.Getenv("BULLETTRAIN_CAR_CONTEXT_PAINT"); t.paint == "" {
		t.paint = "black:white"
	}

	return t.paint
}

// CanShow decides if this car needs to be displayed.
func (t *Context) CanShow() bool {
	s := true
	if e := os.Getenv("BULLETTRAIN_CAR_CONTEXT_SHOW"); e == "false" {
		s = false
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (t *Context) Render(out chan<- string) {
	defer close(out)

	var username string
	u, e := user.Current()
	if e == nil {
		username = u.Username
	} else {
		username = "---"
	}

	hostname, e := os.Hostname()
	if e != nil {
		hostname = "---"
	}

	out <- ansi.Color(fmt.Sprintf(" %s@%s ", username, hostname),
		t.GetPaint())
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (t *Context) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_CONTEXT_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (t *Context) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_CONTEXT_SEPARATOR_SYMBOL")
}
