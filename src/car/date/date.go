package carDate

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/bullettrain-sh/bullettrain-go-core/src/ansi"
)

const (
	carPaint    = "white:black"
	symbolIcon  = ""
	symbolPaint = "red:black"
	// language=GoTemplate
	carTemplate = `{{.Icon | printf "%s" | cs}}{{.Date | c}}`
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

	var dateSymbol string
	if dateSymbol = os.Getenv("BULLETTRAIN_CAR_DATE_SYMBOL_ICON"); dateSymbol == "" {
		dateSymbol = symbolIcon
	}

	var dateSymbolPaint string
	if dateSymbolPaint = os.Getenv("BULLETTRAIN_CAR_DATE_SYMBOL_PAINT"); dateSymbolPaint == "" {
		dateSymbolPaint = symbolPaint
	}

	y, m, d := time.Now().Date()
	dateText := fmt.Sprintf("%02d-%02d-%02d", y, m, d)

	var t string
	if t = os.Getenv("BULLETTRAIN_CAR_DATE_TEMPLATE"); t == "" {
		t = carTemplate
	}

	funcMap := template.FuncMap{
		// Pipeline functions for colouring.
		"c":  func(t string) string { return ansi.Color(t, c.GetPaint()) },
		"cs": func(t string) string { return ansi.Color(t, dateSymbolPaint) },
	}

	tpl := template.Must(template.New("date").Funcs(funcMap).Parse(t))
	data := struct {
		Icon string
		Date string
	}{Icon: dateSymbol, Date: dateText}
	dateFromTpl := new(bytes.Buffer)
	err := tpl.Execute(dateFromTpl, data)
	if err != nil {
		log.Fatalf("Can't generate the date template: %s", err.Error())
	}

	out <- dateFromTpl.String()
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

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
	return os.Getenv("BULLETTRAIN_CAR_DATE_SEPARATOR_TEMPLATE")
}
