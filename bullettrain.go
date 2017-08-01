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
	"github.com/fatih/color"
)

// Here we must check if ENVIRONMENT variables have the appropriate value
// we may expect them to hold.
func init() {

}

func main() {
	color.NoColor = false // force terminal to use colours
	const lineEnding = "$"

	var segmentList []renderer = getSegments()

	// Create a channel for each segment.
	chans := make([]chan string, len(segmentList))
	for i := range chans {
		chans[i] = make(chan string)
	}

	// Spin off a goroutine for each segment.
	for i, segment := range segmentList {
		go segment.Render(chans[i])
	}

	// Gather each goroutine's response through their channels,
	// keeping their order.
	for _, c := range chans {
		fmt.Print(<-c)
	}

	fmt.Printf("\n%s x", color.HiGreenString(lineEnding))
}

type renderer interface {
	Render(c chan<- string)
}

// Configure the segments and store them in the right order.
func getSegments() []renderer {
	return []renderer{
		&timeSegment{color.FgHiWhite, color.BgBlack},
		&separator{color.FgBlack, color.BgYellow},
		&python.Segment{color.FgHiWhite, color.BgYellow},
		&separator{color.FgYellow, color.BgRed},
		&ruby.Segment{color.FgHiWhite, color.BgRed},
		&separator{color.FgRed, color.BgBlue},
		&golang.Segment{color.FgHiWhite, color.BgBlue},
		&separator{color.FgBlue, color.BgGreen},
		&nodejs.Segment{color.FgHiWhite, color.BgGreen},
		&separator{color.FgGreen, nil},
	}
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
	// let's have a space at the end to make sure it will leave enough space in
	// terminals to render the char correclty
	const segmentSeparator string = " "
	defer close(ch)

	col := color.New(s.fg)
	switch s.bg.(type) {
	case color.Attribute:
		col.Add(s.bg.(color.Attribute))
	}

	ch <- col.Sprint(segmentSeparator)
}

//  _____ _
// |_   _(_)
//   | |  _ _ __ ___   ___
//   | | | | '_ ` _ \ / _ \
//   | | | | | | | | |  __/
//   \_/ |_|_| |_| |_|\___|

type timeSegment struct {
	fg, bg color.Attribute
}

func (s *timeSegment) Render(ch chan<- string) {
	defer close(ch)

	col := color.New(s.fg, s.bg)
	t := time.Now()
	ch <- col.Sprintf(" %02d:%02d:%02d ", t.Hour(), t.Minute(), t.Second())
}
