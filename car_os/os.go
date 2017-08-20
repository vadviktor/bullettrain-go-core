package carOs

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"runtime"

	"github.com/mgutz/ansi"
)

const carPaint = "white:cyan"
const symbolPaint = "white:cyan"

// Os car
type Os struct {
	paint string
}

// GetPaint returns the calculated end paint string for the car.
func (t *Os) GetPaint() string {
	if t.paint = os.Getenv("BULLETTRAIN_CAR_OS_PAINT"); t.paint == "" {
		t.paint = carPaint
	}

	return t.paint
}

// CanShow decides if this car needs to be displayed.
func (t *Os) CanShow() bool {
	s := false
	if e := os.Getenv("BULLETTRAIN_CAR_OS_SHOW"); e == "true" {
		s = true
	}

	return s
}

func paintedSymbol(osName string) string {
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
		"opensuse":   "",
		"raspbian":   "",
		"redhat":     "",
		"sabayon":    "",
		"slackware":  "",
		"ubuntu":     "",
		"tux":        ""}

	var symbol string
	if symbol = os.Getenv("BULLETTRAIN_CAR_TIME_SYMBOL_ICON"); symbol == "" {
		var present bool
		symbol, present = osSymbols[osName]
		if !present {
			symbol = osSymbols["tux"]
		}
		symbol = fmt.Sprintf(" %s ", symbol)
	}

	var osSymbolPaint string
	if osSymbolPaint = os.Getenv("BULLETTRAIN_CAR_TIME_SYMBOL_PAINT"); osSymbolPaint == "" {
		osSymbolPaint = symbolPaint
	}

	return ansi.Color(symbol, osSymbolPaint)
}

func findOutOs() string {
	// We know it's a Mac.
	if runtime.GOOS == "darwin" {
		return "darwin"
	}

	fName := "/etc/os-release"
	fBody, fErr := ioutil.ReadFile(fName)
	if fErr != nil {
		panic("/etc/os-release could not be read.")
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
func (t *Os) Render(out chan<- string) {
	defer close(out)
	carPaint := ansi.ColorFunc(t.GetPaint())

	var n string
	if n = os.Getenv("BULLETTRAIN_CAR_OS_NAME"); n == "" {
		n = findOutOs()
	}

	out <- fmt.Sprintf("%s%s",
		paintedSymbol(n),
		carPaint(n))
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (t *Os) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_OS_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (t *Os) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_OS_SEPARATOR_SYMBOL")
}
