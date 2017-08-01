package main

import (
	"fmt"
	"time"

	golang "github.com/bullettrain-sh/bullettrain-go-golang"
	nodejs "github.com/bullettrain-sh/bullettrain-go-nodejs"
	python "github.com/bullettrain-sh/bullettrain-go-python"
	ruby "github.com/bullettrain-sh/bullettrain-go-ruby"
	//git "github.com/bullettrain-sh/bullettrain-go-git"
	//php "github.com/bullettrain-sh/bullettrain-go-php"

	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	color.NoColor = false // Force terminal to use colours.
	const lineEnding = "$"

	// List of cars to be available for use.
	cars := map[string]renderer{
		"time":   &timeCar{},
		"python": &python.Segment{},
		"ruby":   &ruby.Segment{},
		"golang": &golang.Segment{},
		"nodejs": &nodejs.Segment{},
	}

	var car_order_list []string = carOrder()
	// Create a channel for each car.
	chans := make([]chan string, len(car_order_list))
	for i := range chans {
		chans[i] = make(chan string)
	}

	// Spin off a goroutine for each car.
	for i, car := range car_order_list {
		go cars[car].Render(chans[i])
	}

	// Gather each goroutine's response through their channels,
	// keeping their order.
	for _, c := range chans {
		fmt.Print(<-c)
	}

	fmt.Printf("\n%s x", color.HiGreenString(lineEnding))
}

// Defining the order of the cars in which they must be printed,
// also defining the list of cars which are actually used.
func carOrder() []string {
	var car_order string = os.Getenv("BULLETTRAIN_CAR_ORDER")
	if car_order == "" {
		// baked in default car order
		return []string{
			"time",
			"python",
			"ruby",
			"golang",
			"nodejs",
		}
	} else {
		return strings.Split(strings.TrimSpace(car_order), " ")
	}
}

type renderer interface {
	Render(c chan<- string)
}

//  _____                            _
// /  ___|                          | |
// \ `--.  ___ _ __   __ _ _ __ __ _| |_ ___  _ __
//  `--. \/ _ \ '_ \ / _` | '__/ _` | __/ _ \| '__|
// /\__/ /  __/ |_) | (_| | | | (_| | || (_) | |
// \____/ \___| .__/ \__,_|_|  \__,_|\__\___/|_|
//            | |
//            |_|

type separator struct {
	fg color.Attribute
	bg interface{}
}

func (s *separator) Render(ch chan<- string) {
	// Let's have a space at the end to make sure it will leave enough space in
	// terminals to render the char correclty.
	const carSeparator string = ""
	defer close(ch)

	col := color.New(s.fg)
	switch s.bg.(type) {
	case color.Attribute:
		col.Add(s.bg.(color.Attribute))
	}

	ch <- col.Sprint(carSeparator)
}

//  _____ _
// |_   _(_)
//   | |  _ _ __ ___   ___
//   | | | | '_ ` _ \ / _ \
//   | | | | | | | | |  __/
//   \_/ |_|_| |_| |_|\___|

type timeCar struct {
	fg, bg color.Attribute
}

func (s *timeCar) Render(ch chan<- string) {
	const time_symbol = ""
	defer close(ch)

	col := color.New(s.fg, s.bg)
	t := time.Now()
	ch <- col.Sprintf("%s %02d:%02d:%02d ",
		time_symbol, t.Hour(), t.Minute(), t.Second())
}
