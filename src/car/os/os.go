package carOs

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime"
	"text/template"

	"github.com/bullettrain-sh/bullettrain-go-core/src/ansi"
)

const (
	carPaint    = "white:cyan"
	symbolPaint = "white:cyan"
	carTemplate = `{{.Icon | printf "%s " | cs}}{{.Name | c}}`
)

// Os car
type Car struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_OS_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	s := false
	if e := os.Getenv("BULLETTRAIN_CAR_OS_SHOW"); e == "true" {
		s = true
	}

	return s
}

func symbol(osName string) string {
	osSymbols := map[string]string{
		"arch":       "",
		"centos":     "",
		"coreos":     "",
		"darwin":     "",
		"debian":     "",
		"elementary": "",
		"fedora":     "",
		"freebsd":    "",
		"gentoo":     "",
		"linuxmint":  "",
		"mageia":     "",
		"mandriva":   "",
		"manjaro":    "",
		"opensuse":   "",
		"raspbian":   "",
		"redhat":     "",
		"sabayon":    "",
		"slackware":  "",
		"ubuntu":     "",
		"tux":        ""}

	var symbol string
	if symbol = os.Getenv("BULLETTRAIN_CAR_OS_SYMBOL_ICON"); symbol == "" {
		var present bool
		symbol, present = osSymbols[osName]
		if !present {
			symbol = osSymbols["tux"]
		}
	}

	return symbol
}

func FindOutOs() string {
	// We know it's a Mac.
	if runtime.GOOS == "darwin" {
		return "darwin"
	}

	fName := "/etc/os-release"
	fBody, fErr := ioutil.ReadFile(fName)
	if fErr != nil {
		log.Fatalln("/etc/os-release could not be read.")
	}

	flavourExp := regexp.MustCompile(`ID=([a-z]+)`)
	flavourParts := flavourExp.FindStringSubmatch(string(fBody))

	flavour := "tux"
	if len(flavourParts) == 2 && flavourParts[1] != "" {
		flavour = flavourParts[1]
	}

	return flavour
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out)

	var osSymbolPaint string
	if osSymbolPaint = os.Getenv("BULLETTRAIN_CAR_OS_SYMBOL_PAINT"); osSymbolPaint == "" {
		osSymbolPaint = symbolPaint
	}

	var s string
	if s = os.Getenv("BULLETTRAIN_CAR_OS_TEMPLATE"); s == "" {
		s = carTemplate
	}

	funcMap := template.FuncMap{
		// Pipeline functions for colouring.
		"c":  func(t string) string { return ansi.Color(t, c.GetPaint()) },
		"cs": func(t string) string { return ansi.Color(t, osSymbolPaint) },
	}

	osName := FindOutOs()
	tpl := template.Must(template.New("os").Funcs(funcMap).Parse(s))
	data := struct {
		Icon string
		Name string
	}{Icon: symbol(osName), Name: osName}
	fromTpl := new(bytes.Buffer)
	err := tpl.Execute(fromTpl, data)
	if err != nil {
		log.Fatalf("Can't generate the OS template: %s", err.Error())
	}

	out <- fromTpl.String()
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_OS_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_OS_SEPARATOR_SYMBOL")
}

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
	return os.Getenv("BULLETTRAIN_CAR_OS_SEPARATOR_TEMPLATE")
}
