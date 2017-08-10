package main

import (
	"bytes"
	"fmt"
	"os"
	"os/user"
	"regexp"
	"strings"

	"github.com/bullettrain-sh/bullettrain-go-core/car_context"
	"github.com/bullettrain-sh/bullettrain-go-core/car_date"
	"github.com/bullettrain-sh/bullettrain-go-core/car_directory"
	"github.com/bullettrain-sh/bullettrain-go-core/car_time"
	"github.com/bullettrain-sh/bullettrain-go-python"
	"github.com/mgutz/ansi"
)

func main() {
	if d := os.Getenv("BULLETTRAIN_NO_PAINT"); d == "true" {
		ansi.DisableColors(true)
	}

	// List of cars available for use.
	trailers := carsOrderByTrigger()

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
		go trailers[k].Render(chans[j])
		lastSeparator = j+2 == noOfCars

		sep := &separator{}
		var newPaint string
		if lastSeparator {
			newPaint = paintFlipper(
				trailers[k].GetPaint(),
				"default:default")
		} else {
			newPaint = paintFlipper(
				trailers[k].GetPaint(),
				trailers[k+1].GetPaint())
		}
		go sep.Render(chans[j+1], newPaint)
	}

	var n bytes.Buffer
	// Gather each goroutine's response through their channels,
	// keeping their order.
	for _, c := range chans {
		n.WriteString(<-c)
	}

	if l := os.Getenv("BULLETTRAIN_CARS_SEPARATE_LINE"); l == "true" {
		n.WriteString("\n")
	}

	fmt.Printf("%s%s", n.String(), lineEnding())
}

func carsOrderByTrigger() []carRenderer {
	// Cars basic, default order.
	var o []string
	if envOrder := os.Getenv("BULLETTRAIN_CARS"); envOrder == "" {
		o = append(o, "time", "date", "context", "dir", "python")
	} else {
		o = strings.Split(strings.TrimSpace(envOrder), " ")
	}

	// List of cars to be available for use.
	trailers := map[string]carRenderer{
		"time":    &car_time.Time{},
		"date":    &car_date.Date{},
		"context": &car_context.Context{},
		"dir":     &car_directory.Directory{},
		"python":  &car_python.Car{},
	}

	var carsToRender []carRenderer
	for _, car := range o {
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

	var l, c string
	if u.Username == "root" {
		if l = os.Getenv("BULLETTRAIN_PROMPT_CHAR_ROOT"); l == "" {
			l = "# "
		}

		if c = os.Getenv("BULLETTRAIN_PROMPT_CHAR_ROOT_PAINT"); c == "" {
			c = "red"
		}
	} else {
		if l = os.Getenv("BULLETTRAIN_PROMPT_CHAR"); l == "" {
			l = "$ "
		}

		if c = os.Getenv("BULLETTRAIN_PROMPT_CHAR_PAINT"); c == "" {
			c = "green"
		}
	}

	return ansi.Color(l, c)
}

// Flip the FG and BG setup in colour strings of cars for a separator.
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
	// The end product of a completely composed car.
	Render(out chan<- string)
	// The calculated end paint string for the car.
	GetPaint() string
	// Decides if this car needs to be displayed.
	CanShow() bool
}

type separator struct {
	paint string
}

func (s *separator) Render(out chan<- string, paintOverride string) {
	defer close(out)

	var symbol string
	if symbol = os.Getenv("BULLETTRAIN_SEPARATOR_ICON"); symbol == "" {
		symbol = " "
	}

	// TODO customisable separator colour for each car

	out <- ansi.Color(symbol, paintOverride)
}
