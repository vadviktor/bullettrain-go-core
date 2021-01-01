package carPhp

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/bullettrain-sh/bullettrain-go-core/src/ansi"
)

const (
	carPaint          = "white:69"
	pythonSymbolPaint = "black:69"
	pythonSymbolIcon  = "\uE608" // 
	// language=GoTemplate
	carTemplate = `{{.Icon | printf "%s " | cs}}{{.Info | c}}`
)

// Car for PHP
type Car struct {
	paint string
	// Current directory
	Pwd string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_PHP_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	if e := os.Getenv("BULLETTRAIN_CAR_PHP_SHOW"); e == "true" {
		return true
	}

	var d string
	if d = c.Pwd; d == "" {
		return false
	}

	// Show when .php files exist in current directory
	phpPattern := fmt.Sprintf("%s%s*.php", d, string(os.PathSeparator))
	phpFiles, _ := filepath.Glob(phpPattern)
	if phpFiles != nil {
		return true
	}

	return false
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out) // Always close the channel!

	var symbolIcon string
	if symbolIcon = os.Getenv("BULLETTRAIN_CAR_PHP_SYMBOL"); symbolIcon == "" {
		symbolIcon = pythonSymbolIcon
	}

	var symbolPaint string
	if symbolPaint = os.Getenv("BULLETTRAIN_CAR_PHP_SYMBOL_PAINT"); symbolPaint == "" {
		symbolPaint = pythonSymbolPaint
	}

	var s string
	if s = os.Getenv("BULLETTRAIN_CAR_USER_TEMPLATE"); s == "" {
		s = carTemplate
	}

	funcMap := template.FuncMap{
		// Pipeline functions for colouring.
		"c":  func(t string) string { return ansi.Color(t, c.GetPaint()) },
		"cs": func(t string) string { return ansi.Color(t, symbolPaint) },
	}

	var version string
	cmd := exec.Command("php", "--version")
	cmdOut, err := cmd.CombinedOutput()
	if err == nil {
		// language=GoRegExp
		re := regexp.MustCompile(`([0-9.]+)`)
		versionArr := re.FindStringSubmatch(string(cmdOut))
		if len(versionArr) > 0 {
			version = versionArr[1]
		}
	} else {
		version = strings.Replace(string(cmdOut), "\n", " ", -1)
	}

	tpl := template.Must(template.New("php").Funcs(funcMap).Parse(s))
	data := struct {
		Icon string
		Info string
	}{Icon: symbolIcon, Info: version}
	fromTpl := new(bytes.Buffer)
	err = tpl.Execute(fromTpl, data)
	if err != nil {
		log.Fatalf("Can't generate the php template: %s", err.Error())
	}

	out <- fromTpl.String()
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_PHP_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_PHP_SEPARATOR_SYMBOL")
}

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
	return os.Getenv("BULLETTRAIN_CAR_PHP_SEPARATOR_TEMPLATE")
}
