package carK8s

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/bullettrain-sh/bullettrain-go-core/src/ansi"
)

const (
	carPaint    = "white+h:yellow"
	helmSymbol  = "\uFD31" // ﴱ
	carTemplate = `{{.Icon | printf " %s " | c}}{{.Context | c}}`
)

type Car struct {
	paint string
	// Current directory
	Pwd string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_K8S_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	path, _ := exec.LookPath("kubectl")

	if path == "" {
		return false
	}

	s := true
	if e := os.Getenv("BULLETTRAIN_CAR_K8S_SHOW"); e == "false" {
		s = false
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out)

	cmd := exec.Command("kubectl", "config", "current-context")
	output, err := cmd.Output()
	var context string
	if err == nil {
		context = strings.Trim(string(output), "\n")
	} else {
		context = "ERR!"
	}

	var symbol string
	if symbol = os.Getenv("BULLETTRAIN_CAR_K8S_SYMBOL_ICON"); symbol == "" {
		symbol = helmSymbol
	}

	var s string
	if s = os.Getenv("BULLETTRAIN_CAR_K8S_TEMPLATE"); s == "" {
		s = carTemplate
	}

	funcMap := template.FuncMap{
		// Pipeline functions for colouring.
		"c": func(t string) string { return ansi.Color(t, c.GetPaint()) },
	}

	tpl := template.Must(template.New("k8s").Funcs(funcMap).Parse(s))
	data := struct {
		Icon    string
		Context string
	}{Icon: symbol, Context: context}
	contextFromTpl := new(bytes.Buffer)
	err = tpl.Execute(contextFromTpl, data)
	if err != nil {
		log.Fatalf("Can't generate the time template: %s", err.Error())
	}

	out <- contextFromTpl.String()
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_K8S_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_K8S_SEPARATOR_SYMBOL")
}

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
	return os.Getenv("BULLETTRAIN_CAR_K8S_SEPARATOR_TEMPLATE")
}
