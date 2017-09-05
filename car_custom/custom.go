package carCustom

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/bullettrain-sh/bullettrain-go-core/ansi"
)

// Custom car
type Car struct {
	// Uppercase
	callword string
	pwd      string
}

// SetCallword makes sure the callword is in the right shape.
func (c *Car) SetCallword(word string) {
	c.callword = strings.ToUpper(word)
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_PLUGIN_" + c.callword + "_PAINT")
}

func (c *Car) paintedSymbol() string {
	return ansi.Color(
		os.Getenv("BULLETTRAIN_CAR_PLUGIN_"+c.callword+"_SYMBOL_ICON"),
		os.Getenv("BULLETTRAIN_CAR_PLUGIN_"+c.callword+"_SYMBOL_PAINT"))
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	s := true
	if e := os.Getenv("BULLETTRAIN_CAR_PLUGIN_" + c.callword + "_SHOW"); e == "false" {
		s = false
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out)
	carPaint := ansi.ColorFunc(c.GetPaint())
	var stuff string

	cmdElem := strings.Fields(
		os.Getenv("BULLETTRAIN_CAR_PLUGIN_" + c.callword + "_CMD"))
	cmd := exec.Command(cmdElem[0], cmdElem[1:]...)
	cmdOut, err := cmd.Output()
	if err == nil {
		stuff = string(cmdOut)
	} else {
		stuff = "xxx"
	}

	out <- fmt.Sprintf("%s%s", c.paintedSymbol(), carPaint(stuff))
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_PLUGIN_" + c.callword + "_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_PLUGIN_" + c.callword + "_SEPARATOR_SYMBOL")
}
