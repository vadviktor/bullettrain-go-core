package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"runtime"
	"strings"

	"github.com/bullettrain-sh/bullettrain-go-core/ansi"
	"github.com/bullettrain-sh/bullettrain-go-core/car_context"
	"github.com/bullettrain-sh/bullettrain-go-core/car_date"
	"github.com/bullettrain-sh/bullettrain-go-core/car_directory"
	"github.com/bullettrain-sh/bullettrain-go-core/car_os"
	"github.com/bullettrain-sh/bullettrain-go-core/car_status"
	"github.com/bullettrain-sh/bullettrain-go-core/car_time"
	"github.com/bullettrain-sh/bullettrain-go-git"
	"github.com/bullettrain-sh/bullettrain-go-golang"
	"github.com/bullettrain-sh/bullettrain-go-nodejs"
	"github.com/bullettrain-sh/bullettrain-go-php"
	"github.com/bullettrain-sh/bullettrain-go-python"
	"github.com/bullettrain-sh/bullettrain-go-ruby"
)

const defaultCarOrder = "os time date context dir python go ruby nodejs php git status"
const separatorSymbol = " "

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	if d := os.Getenv("BULLETTRAIN_NO_PAINT"); d == "true" {
		ansi.DisableColors(true)
	}

	buildAndPrintCars()
}

func buildAndPrintCars() {
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
		go sep.Render(chans[j+1], newPaint, trailers[k].GetSeparatorSymbol())
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
	n.WriteRune(' ')

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
		o = strings.Split(strings.TrimSpace(defaultCarOrder), " ")
	} else {
		o = strings.Split(strings.TrimSpace(envOrder), " ")
	}

	return
}

func carsToRender() []carRenderer {
	d := pwd()
	// List of cars to be available for use.
	trailers := map[string]carRenderer{
		"context": &carContext.Car{},
		"date":    &carDate.Car{},
		"dir":     &carDirectory.Car{Pwd: d},
		"git":     &carGit.Car{Pwd: d},
		"go":      &carGo.Car{Pwd: d},
		"nodejs":  &carNodejs.Car{Pwd: d},
		"os":      &carOs.Car{},
		"php":     &carPhp.Car{Pwd: d},
		"python":  &carPython.Car{Pwd: d},
		"ruby":    &carRuby.Car{Pwd: d},
		"status":  &carStatus.Car{},
		"time":    &carTime.Car{},
	}

	var carsToRender []carRenderer
	for _, car := range carsOrder() {
		if trailers[car].CanShow() {
			carsToRender = append(carsToRender, trailers[car])
		}
	}

	return carsToRender
}

func lineEnding() string {
	u, e := user.Current()
	if e != nil {
		panic("Can't figure out current username.")
	}

	var c, l string
	if u.Username == "root" {
		if l = os.Getenv("BULLETTRAIN_PROMPT_CHAR_ROOT"); l == "" {
			l = "#"
		}

		if c = os.Getenv("BULLETTRAIN_PROMPT_CHAR_ROOT_PAINT"); c == "" {
			c = "red"
		}
	} else {
		if l = os.Getenv("BULLETTRAIN_PROMPT_CHAR"); l == "" {
			l = "$"
		}

		if c = os.Getenv("BULLETTRAIN_PROMPT_CHAR_PAINT"); c == "" {
			c = "green"
		}
	}

	return ansi.Color(l, c)
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

type carRenderer interface {
	// Render builds and passes the end product of a completely composed car onto
	// the channel.
	Render(out chan<- string)

	// GetPaint returns the calculated end paint string for the car.
	GetPaint() string

	// CanShow decides if this car needs to be displayed.
	CanShow() bool

	// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
	// separator through ENV variables.
	GetSeparatorPaint() string

	// GetSeparatorSymbol overrides the symbol of the right hand side
	// separator through ENV variables.
	GetSeparatorSymbol() string
}

type separator string

func (s *separator) Render(out chan<- string, paint, symbolOverride string) {
	defer close(out)

	var symbol string
	if symbolOverride != "" {
		symbol = symbolOverride
	} else if symbol = os.Getenv("BULLETTRAIN_SEPARATOR_ICON"); symbol == "" {
		symbol = separatorSymbol
	}

	out <- ansi.Color(symbol, paint)
}
