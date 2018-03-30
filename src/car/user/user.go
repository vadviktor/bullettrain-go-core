package carUser

import (
	"bytes"
	"log"
	"os"
	"os/user"
	"text/template"

	"github.com/bullettrain-sh/bullettrain-go-core/src/ansi"
)

const (
	carPaint    = "black:white"
	carTemplate = `{{.User | c}}`
)

// User car
type Car struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_USER_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	s := true
	if e := os.Getenv("BULLETTRAIN_CAR_USER_SHOW"); e == "false" {
		s = false
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out)

	var username string
	u, e := user.Current()
	if e == nil {
		username = u.Username
	}

	var s string
	if s = os.Getenv("BULLETTRAIN_CAR_USER_TEMPLATE"); s == "" {
		s = carTemplate
	}

	funcMap := template.FuncMap{
		// Pipeline functions for colouring.
		"c": func(t string) string { return ansi.Color(t, c.GetPaint()) },
	}

	tpl := template.Must(template.New("user").Funcs(funcMap).Parse(s))
	data := struct{ User string }{User: username}
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
	return os.Getenv("BULLETTRAIN_CAR_USER_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_USER_SEPARATOR_SYMBOL")
}

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
	return os.Getenv("BULLETTRAIN_CAR_USER_SEPARATOR_TEMPLATE")
}
