package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strings"
	"text/template"

	"github.com/bullettrain-sh/bullettrain-go-core/src/ansi"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/custom"
)

type carRenderer interface {
	// Render builds and passes the end product of a completely composed car onto
	// the channel.
	Render(out chan<- string)

	// GetPaint returns the calculated end paint string for the car.
	GetPaint() string

	// CanShow decides if this car needs to be displayed.
	CanShow() bool

	// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
	// separator through ENV variable.
	GetSeparatorPaint() string

	// GetSeparatorSymbol overrides the symbol of the right hand side
	// separator through ENV variable.
	GetSeparatorSymbol() string

	// GetSeparatorTemplate overrides the template of the right hand side
	// separator through ENV variable.
	GetSeparatorTemplate() string
}

type separator string

// init defines some steps that affects running the program and needs provisioning.
func init() {
	if d := os.Getenv("BULLETTRAIN_NO_PAINT"); d == "true" {
		ansi.DisableColors(true)
	}
}

func main() {
	// List of cars available for use.
	trailers := carsToRender()

	// Create a channel for each car.
	noOfCars := len(trailers) * 2
	chans := make([]chan string, noOfCars)
	for i := range chans {
		chans[i] = make(chan string)
	}

	// Spin off a goroutine for each car.
	var lastSeparator bool
	paintFlipper := flipPaint()
	for j, k := 0, 0; j < noOfCars; j, k = j+2, k+1 {
		// Render car.
		go trailers[k].Render(chans[j])

		// Render separator.
		var newPaint string
		if newPaint = trailers[k].GetSeparatorPaint(); newPaint == "" {
			lastSeparator = j+2 == noOfCars
			if lastSeparator {
				newPaint = paintFlipper(
					trailers[k].GetPaint(),
					"default:default")
			} else {
				newPaint = paintFlipper(
					trailers[k].GetPaint(),
					trailers[k+1].GetPaint())
			}
		}

		sep := new(separator)
		go sep.Render(chans[j+1], newPaint,
			trailers[k].GetSeparatorSymbol(),
			trailers[k].GetSeparatorTemplate())
	}

	var n bytes.Buffer
	// Gather each goroutine's response through their channels,
	// keeping their order.
	for _, c := range chans {
		n.WriteString(<-c)
	}

	if l := os.Getenv("BULLETTRAIN_CARS_SEPARATE_LINE"); l != "false" {
		n.WriteRune('\n')
	}

	n.WriteString(lineEnding())

	if d := os.Getenv("BULLETTRAIN_DEBUG"); d == "true" {
		fmt.Printf("%+ x", n.String())
		fmt.Println("")
		fmt.Printf("%+q", n.String())
	} else {
		fmt.Print(n.String())
	}
}

// pwd returns the current directory path.
func pwd() string {
	cmd := exec.Command("pwd", "-P")
	pwd, err := cmd.Output()
	var d string
	if err == nil {
		d = strings.Trim(string(pwd), "\n")
	}

	return d
}

func carsOrder() (o []string) {
	if envOrder := os.Getenv("BULLETTRAIN_CARS"); envOrder == "" {
		o = strings.Fields(defaultCarOrder)
	} else {
		o = strings.Fields(envOrder)
	}

	return
}

func carsToRender() []carRenderer {
	trailers := trailers(pwd())

	var carsToRender []carRenderer
	for _, car := range carsOrder() {
		c, ex := trailers[car]
		if ex {
			if c.CanShow() {
				carsToRender = append(carsToRender, c)
			}
		} else {
			customCar := new(carCustom.Car)
			customCar.SetCallword(car)
			if customCar.CanShow() {
				carsToRender = append(carsToRender, customCar)
			}
		}
	}

	return carsToRender
}

func lineEnding() string {
	u, e := user.Current()
	if e != nil {
		log.Fatalf("Can't figure out current username: %s\n", e.Error())
	}

	var c, s, t string
	if u.Username == "root" {
		if s = os.Getenv("BULLETTRAIN_PROMPT_CHAR_ROOT"); s == "" {
			s = "#"
		}

		if c = os.Getenv("BULLETTRAIN_PROMPT_CHAR_ROOT_PAINT"); c == "" {
			c = "red"
		}

		if t = os.Getenv("BULLETTRAIN_PROMPT_CHAR_ROOT_TEMPLATE"); t == "" {
			t = promptCharTemplate
		}
	} else {
		if s = os.Getenv("BULLETTRAIN_PROMPT_CHAR"); s == "" {
			s = "$"
		}

		if c = os.Getenv("BULLETTRAIN_PROMPT_CHAR_PAINT"); c == "" {
			c = "green"
		}

		if t = os.Getenv("BULLETTRAIN_PROMPT_CHAR_TEMPLATE"); t == "" {
			t = promptCharTemplate
		}
	}

	funcMap := template.FuncMap{
		// Pipeline function for colouring.
		"c": func(t string) string { return ansi.Color(t, c) },
	}

	tpl := template.Must(template.New("promptChar").Funcs(funcMap).Parse(t))
	d := struct{ Icon string }{Icon: s}
	symbolFromTpl := new(bytes.Buffer)
	err := tpl.Execute(symbolFromTpl, d)
	if err != nil {
		log.Fatalf("Can't generate the prompt char template: %s", err.Error())
	}

	return symbolFromTpl.String()
}

// flipPaint flips the FG and BG setup in colour strings of cars for a separator.
// Use it as a closure.
func flipPaint() func(string, string) string {
	// foregroundColor+attributes:backgroundColor+attributes
	colourExp := regexp.MustCompile(`\w*\+?\w*:?(\w*)\+?\w?`)

	flipped := func(currentPaint, nextPaint string) string {
		currentParts := colourExp.FindStringSubmatch(currentPaint)
		nextParts := colourExp.FindStringSubmatch(nextPaint)

		newFg := "default"
		if len(currentParts) == 2 && currentParts[1] != "" {
			newFg = currentParts[1]
		}

		newBg := "default"
		if len(nextParts) == 2 && nextParts[1] != "" {
			newBg = nextParts[1]
		}

		return fmt.Sprintf("%s:%s", newFg, newBg)
	}

	return flipped
}

func (s *separator) Render(out chan<- string, paint, symbolOverride, templateOverride string) {
	defer close(out)

	var symbol string
	if symbolOverride != "" {
		symbol = symbolOverride
	} else if symbol = os.Getenv("BULLETTRAIN_SEPARATOR_ICON"); symbol == "" {
		symbol = separatorSymbol
	}

	var t string
	if templateOverride != "" {
		t = templateOverride
	} else if t = os.Getenv("BULLETTRAIN_SEPARATOR_TEMPLATE"); t == "" {
		t = separatorTemplate
	}

	funcMap := template.FuncMap{
		// Pipeline function for colouring.
		"c": func(t string) string { return ansi.Color(t, paint) },
	}

	tpl := template.Must(template.New("separator").Funcs(funcMap).Parse(t))
	d := struct{ Icon string }{Icon: symbol}
	symbolFromTpl := new(bytes.Buffer)
	err := tpl.Execute(symbolFromTpl, d)
	if err != nil {
		log.Fatalf("Can't generate the separator template: %s", err.Error())
	}

	out <- symbolFromTpl.String()
}
