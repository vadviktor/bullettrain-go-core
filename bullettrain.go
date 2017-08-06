package main

import (
	"fmt"
	"os"
	"os/user"
	"regexp"
	"strings"

	"github.com/bullettrain-sh/bullettrain-go-core/cars"
	"github.com/bullettrain-sh/bullettrain-go-python"
	"github.com/mgutz/ansi"
)

func main() {
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

	// Gather each goroutine's response through their channels,
	// keeping their order.
	for _, c := range chans {
		fmt.Print(<-c)
	}

	newLine := "false"
	if newLine = os.Getenv("BULLETTRAIN_CARS_SEPARATE_LINE"); newLine == "true" {
		newLine = "\n"
	}

	fmt.Printf("%s%s ", newLine, lineEnding())
}

func lineEnding() string {
	u, e := user.Current()
	if e != nil {
		panic("Can't figure out current username.")
	}

	var l, c string
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

func carsOrderByTrigger() []carRenderer {
	var o []string
	if envOrder := os.Getenv("BULLETTRAIN_CAR_ORDER"); envOrder == "" {
		o = append(o, "time", "python")
	} else {
		o = strings.Split(strings.TrimSpace(envOrder), " ")
	}

	// List of cars to be available for use.
	trailers := map[string]carRenderer{
		"time":   &cars.Time{},
		"python": &python.Car{},
	}

	var carsToRender []carRenderer
	for _, car := range o {
		if trailers[car].CanShow() {
			carsToRender = append(carsToRender, trailers[car])
		}
	}

	return carsToRender
}

// Flip the FG and BG setup in colour strings of cars for a separator.
func flipPaint() func(string, string) string {
	// foregroundColor+attributes:backgroundColor+attributes
	colourExp := regexp.MustCompile(`\w*\+?\w*:?(\w*)\+?\w?`)

	flipped := func(currentPaint, nextPaint string) string {
		currentParts := colourExp.FindStringSubmatch(currentPaint)
		nextParts := colourExp.FindStringSubmatch(nextPaint)

		var newFg string = "default"
		if len(currentParts) == 2 && currentParts[1] != "" {
			newFg = currentParts[1]
		}

		var newBg string = "default"
		if len(nextParts) == 2 && nextParts[1] != "" {
			newBg = nextParts[1]
		}

		return fmt.Sprintf("%s:%s", newFg, newBg)
	}

	return flipped
}

type carRenderer interface {
	// The end product of a competely composed car.
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
