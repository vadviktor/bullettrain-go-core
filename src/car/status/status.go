package carStatus

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"github.com/bullettrain-sh/bullettrain-go-core/src/ansi"
)

const (
	carPaint    = "255:160"
	symbolIcon  = ""
	symbolPaint = "220:160"
	carTemplate = `{{.Icon | printf "%s " | cs}}{{.Code | c}}`
)

// Status Car
type Car struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_STATUS_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	if len(os.Args) > 1 {
		return os.Args[1] != "" && os.Args[1] != "0"
	}

	return false
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out)

	var statusSymbol string
	if statusSymbol = os.Getenv("BULLETTRAIN_CAR_STATUS_SYMBOL_ICON"); statusSymbol == "" {
		statusSymbol = symbolIcon
	}

	var statusSymbolPaint string
	if statusSymbolPaint = os.Getenv("BULLETTRAIN_CAR_STATUS_SYMBOL_PAINT"); statusSymbolPaint == "" {
		statusSymbolPaint = symbolPaint
	}

	var s string
	if s = os.Getenv("BULLETTRAIN_CAR_STATUS_TEMPLATE"); s == "" {
		s = carTemplate
	}

	funcMap := template.FuncMap{
		// Pipeline functions for colouring.
		"c":  func(t string) string { return ansi.Color(t, c.GetPaint()) },
		"cs": func(t string) string { return ansi.Color(t, statusSymbolPaint) },
	}

	tpl := template.Must(template.New("status").Funcs(funcMap).Parse(s))
	data := struct {
		Icon string
		Code string
	}{Icon: statusSymbol, Code: os.Args[1]}
	fromTpl := new(bytes.Buffer)
	err := tpl.Execute(fromTpl, data)
	if err != nil {
		log.Fatalf("Can't generate the user template: %s", err.Error())
	}

	out <- fromTpl.String()
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_STATUS_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_STATUS_SEPARATOR_SYMBOL")
}

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
	return os.Getenv("BULLETTRAIN_CAR_STATUS_SEPARATOR_TEMPLATE")
}
