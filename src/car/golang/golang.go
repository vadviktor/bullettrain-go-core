package carGo

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
	carPaint      = "black:123"
	goSymbolPaint = "black:123"
	goSymbolIcon  = "\uE627" // 
	// language=GoTemplate
	carTemplate = `{{.Icon | printf "%s " | cs}}{{.Info | c}}`
)

// Car for Go
type Car struct {
	paint string
	// Current directory
	Pwd string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_GO_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	if e := os.Getenv("BULLETTRAIN_CAR_GO_SHOW"); e == "true" {
		return true
	}

	var d string
	if d = c.Pwd; d == "" {
		return false
	}

	// Show when .go files exist in current directory
	p := fmt.Sprintf("%s%s*.go", d, string(os.PathSeparator))
	f, _ := filepath.Glob(p)
	if f != nil {
		return true
	}

	return false
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out) // Always close the channel!

	var symbolIcon string
	if symbolIcon = os.Getenv("BULLETTRAIN_CAR_GO_SYMBOL_ICON"); symbolIcon == "" {
		symbolIcon = goSymbolIcon
	}

	var symbolPaint string
	if symbolPaint = os.Getenv("BULLETTRAIN_CAR_GO_SYMBOL_PAINT"); symbolPaint == "" {
		symbolPaint = goSymbolPaint
	}

	cmd := exec.Command("go", "version")
	cmdOut, err := cmd.CombinedOutput()
	var version string
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

	var s string
	if s = os.Getenv("BULLETTRAIN_CAR_GO_TEMPLATE"); s == "" {
		s = carTemplate
	}

	funcMap := template.FuncMap{
		// Pipeline functions for colouring.
		"c":  func(t string) string { return ansi.Color(t, c.GetPaint()) },
		"cs": func(t string) string { return ansi.Color(t, symbolPaint) },
	}

	tpl := template.Must(template.New("go").Funcs(funcMap).Parse(s))
	data := struct {
		Icon string
		Info string
	}{Icon: symbolIcon, Info: version}
	fromTpl := new(bytes.Buffer)
	err = tpl.Execute(fromTpl, data)
	if err != nil {
		log.Fatalf("Can't generate the go template: %s", err.Error())
	}

	out <- fromTpl.String()
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_GO_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_GO_SEPARATOR_SYMBOL")
}

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
	return os.Getenv("BULLETTRAIN_CAR_GO_SEPARATOR_TEMPLATE")
}
