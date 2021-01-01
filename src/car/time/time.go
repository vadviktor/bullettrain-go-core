package carTime

import (
	"bytes"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/bullettrain-sh/bullettrain-go-core/src/ansi"
)

const (
	carPaint    = "black:white"
	symbolIcon  = "\uF43A" // 
	symbolPaint = "black:white"
	// language=GoTemplate
	carTemplate = `{{.Icon | printf " %s " | cs}}{{.Time | c}}`
)

// Time Car
type Car struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_TIME_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	s := false
	if e := os.Getenv("BULLETTRAIN_CAR_TIME_SHOW"); e == "true" {
		s = true
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out)

	var timeSymbol string
	if timeSymbol = os.Getenv("BULLETTRAIN_CAR_TIME_SYMBOL_ICON"); timeSymbol == "" {
		timeSymbol = symbolIcon
	}

	var timeSymbolPaint string
	if timeSymbolPaint = os.Getenv("BULLETTRAIN_CAR_TIME_SYMBOL_PAINT"); timeSymbolPaint == "" {
		timeSymbolPaint = symbolPaint
	}

	n := time.Now()
	t := n.Format("15:04:05")
	if h := os.Getenv("BULLETTRAIN_CAR_TIME_12HR"); h == "true" {
		t = n.Format("3:04:05")
	}

	var s string
	if s = os.Getenv("BULLETTRAIN_CAR_TIME_TEMPLATE"); s == "" {
		s = carTemplate
	}

	funcMap := template.FuncMap{
		// Pipeline functions for colouring.
		"c":  func(t string) string { return ansi.Color(t, c.GetPaint()) },
		"cs": func(t string) string { return ansi.Color(t, timeSymbolPaint) },
	}

	tpl := template.Must(template.New("time").Funcs(funcMap).Parse(s))
	data := struct {
		Icon string
		Time string
	}{Icon: timeSymbol, Time: t}
	timeFromTpl := new(bytes.Buffer)
	err := tpl.Execute(timeFromTpl, data)
	if err != nil {
		log.Fatalf("Can't generate the time template: %s", err.Error())
	}

	out <- timeFromTpl.String()
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_TIME_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_TIME_SEPARATOR_SYMBOL")
}

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
	return os.Getenv("BULLETTRAIN_CAR_TIME_SEPARATOR_TEMPLATE")
}
