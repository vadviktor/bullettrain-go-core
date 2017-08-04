package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/bullettrain-sh/bullettrain-go-core/cars"
	"github.com/bullettrain-sh/bullettrain-go-python"
	"github.com/mgutz/ansi"
)

func main() {
	// List of cars to be available for use.
	trailers := map[string]carRenderer{
		"time":   &cars.Time{},
		"python": &python.Car{},
	}

	var carOrderLists []string = carOrder()
	// Create a channel for each car.
	noOfCars := len(carOrderLists) * 2
	chans := make([]chan string, noOfCars)
	for i := range chans {
		chans[i] = make(chan string)
	}

	// Spin off a goroutine for each car.
	var lastSeparator bool
	paintFlipper := flipPaint()
	for j, k := 0, 0; j < noOfCars; j, k = j+2, k+1 {
		go trailers[carOrderLists[k]].Render(chans[j])
		lastSeparator = j+2 == noOfCars

		sep := &separator{}
		var newPaint string
		if lastSeparator {
			newPaint = paintFlipper(
				trailers[carOrderLists[k]].GetPaint(),
				"default:default")
		} else {
			newPaint = paintFlipper(
				trailers[carOrderLists[k]].GetPaint(),
				trailers[carOrderLists[k+1]].GetPaint())
		}
		go sep.Render(chans[j+1], newPaint)
	}

	// Gather each goroutine's response through their channels,
	// keeping their order.
	for _, c := range chans {
		fmt.Print(<-c)
	}

	lineEnding := "$"
	fmt.Printf("\n%s x", ansi.Color(lineEnding, "green"))
}

// Defining the order of the cars in which they must be printed,
// also defining the list of cars which are actually used.
func carOrder() []string {
	if carOrder := os.Getenv("BULLETTRAIN_CAR_ORDER"); carOrder == "" {
		return []string{
			"time",
			"python",
		}
	} else {
		return strings.Split(strings.TrimSpace(carOrder), " ")
	}
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
	Render(out chan<- string)
	GetPaint() string
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

	var symbolPaint string
	if symbolPaint = os.Getenv("BULLETTRAIN_SEPARATOR_PAINT"); symbolPaint == "" {
		symbolPaint = paintOverride
	}

	out <- ansi.Color(symbol, symbolPaint)
}
