package carCustom

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/bullettrain-sh/bullettrain-go-core/pkg/ansi"
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

	var stuff string
	carSymbol := os.Getenv("BULLETTRAIN_CAR_PLUGIN_" + c.callword + "_SYMBOL_ICON")
	carSymbolPaint := os.Getenv("BULLETTRAIN_CAR_PLUGIN_" + c.callword + "_SYMBOL_PAINT")
	carTemplate := os.Getenv("BULLETTRAIN_CAR_PLUGIN_" + c.callword + "_TEMPLATE")

	cmdElem := strings.Fields(
		os.Getenv("BULLETTRAIN_CAR_PLUGIN_" + c.callword + "_CMD"))

	var cmd *exec.Cmd
	if len(cmdElem) < 2 {
		cmd = exec.Command(cmdElem[0])
	} else {
		cmd = exec.Command(cmdElem[0], cmdElem[1:]...)
	}

	cmdOut, err := cmd.CombinedOutput()
	if err == nil {
		stuff = string(cmdOut)
	} else {
		stuff = err.Error()
	}

	funcMap := template.FuncMap{
		// Pipeline functions for colouring.
		"c":  func(t string) string { return ansi.Color(t, c.GetPaint()) },
		"cs": func(t string) string { return ansi.Color(t, carSymbolPaint) },
	}

	tpl := template.Must(template.New(c.callword).Funcs(funcMap).Parse(carTemplate))
	data := struct {
		Icon string
		Info string
	}{Icon: carSymbol, Info: stuff}
	fromTpl := new(bytes.Buffer)
	tplErr := tpl.Execute(fromTpl, data)
	if tplErr != nil {
		log.Fatalf("Can't generate the user template: %s", tplErr.Error())
	}

	out <- fromTpl.String()
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

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
	return os.Getenv("BULLETTRAIN_CAR_PLUGIN_" + c.callword + "_SEPARATOR_TEMPLATE")
}
