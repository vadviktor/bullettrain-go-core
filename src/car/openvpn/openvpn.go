package carOpenvpn

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"text/template"

	"github.com/bullettrain-sh/bullettrain-go-core/src/ansi"
)

const (
	carPaint            = "208:black"
	symbolPaint         = "208:black"
	symbolIcon          = ""
	symbolLocked        = ""
	symbolLockedPaint   = "green:black"
	symbolUnlocked      = ""
	symbolUnlockedPaint = "red:black"
	// language=GoTemplate
	carTemplate = `{{.Icon | printf "%s " | cs}}{{range $name, $status := .Statuses}}{{printStatus $name $status}}{{end}}`
)

// Car for Openvpn
type Car struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_OPENVPN_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	return true
}

func vpnStatuses() map[string]string {
	var serviceHandle string
	serviceHandle = "openvpn-client@"

	cmd := exec.Command("systemctl", "list-units", "-a",
		fmt.Sprintf("%s*", serviceHandle))
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to get systemd statuses: %s", err.Error())
	}

	// language=GoRegExp
	re := regexp.MustCompile(fmt.Sprintf(
		`%s(.*)\.service.*(running|failed|dead|auto-restart)`, serviceHandle))
	matches := re.FindAllStringSubmatch(string(out), -1)
	statuses := make(map[string]string)
	for _, match := range matches {
		statuses[match[1]] = match[2]
	}

	return statuses
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out) // Always close the channel!

	var si string
	if si = os.Getenv("BULLETTRAIN_CAR_OPENVPN_SYMBOL_ICON"); si == "" {
		si = symbolIcon
	}

	var siu string
	if siu = os.Getenv("BULLETTRAIN_CAR_OPENVPN_SYMBOL_ICON_UNLOCKED"); siu == "" {
		siu = symbolUnlocked
	}

	var spu string
	if spu = os.Getenv("BULLETTRAIN_CAR_OPENVPN_SYMBOL_PAINT_UNLOCKED"); spu == "" {
		spu = symbolUnlockedPaint
	}

	var sil string
	if sil = os.Getenv("BULLETTRAIN_CAR_OPENVPN_SYMBOL_ICON_LOCKED"); sil == "" {
		sil = symbolLocked
	}

	var spl string
	if spl = os.Getenv("BULLETTRAIN_CAR_OPENVPN_SYMBOL_PAINT_LOCKED"); spl == "" {
		spl = symbolLockedPaint
	}

	var sp string
	if sp = os.Getenv("BULLETTRAIN_CAR_OPENVPN_SYMBOL_PAINT"); sp == "" {
		sp = symbolPaint
	}

	funcMap := template.FuncMap{
		// Pipeline functions for colouring.
		"c":  func(t string) string { return ansi.Color(t, c.GetPaint()) },
		"cs": func(t string) string { return ansi.Color(t, sp) },
		"printStatus": func(name string, status string) string {
			statusIcon := siu
			statusColour := spu
			if status == "running" {
				statusIcon = sil
				statusColour = spl
			}

			carPaint := c.GetPaint()
			return fmt.Sprintf("%s%s%s",
				ansi.Color(statusIcon, statusColour),
				ansi.Color(name, carPaint),
				ansi.Color(" ", carPaint))
		},
	}

	tpl := template.Must(
		template.New("openvpn").Funcs(funcMap).Parse(carTemplate))
	data := struct {
		Icon     string
		Statuses map[string]string
	}{Icon: si, Statuses: vpnStatuses()}
	fromTpl := new(bytes.Buffer)
	err := tpl.Execute(fromTpl, data)
	if err != nil {
		log.Fatalf("Can't generate the openvpn template: %s", err.Error())
	}

	out <- fromTpl.String()
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_OPENVPN_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_OPENVPN_SEPARATOR_SYMBOL")
}

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
	return os.Getenv("BULLETTRAIN_CAR_OPENVPN_SEPARATOR_TEMPLATE")
}
