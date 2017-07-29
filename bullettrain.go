package main

import (
	"fmt"
	"time"

	"github.com/bullettrain-sh/bullettrain-go-python"
	"github.com/fatih/color"
)

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
	python := &bullettrain_go_python.PythonSegment{}
	python.SetFg(color.FgHiWhite)
	python.SetBg(color.BgYellow)

	return []renderer{
		&timeSegment{fg: color.FgHiWhite, bg: color.BgBlack},
		&separator{fg: color.FgBlack, bg: color.BgYellow},
		python,
		&separator{fg: color.FgYellow, bg: color.FgYellow},
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
	fg, bg color.Attribute
}

func (s *separator) Render(ch chan<- string) {
	const segmentSeparator string = ""
	col := color.New(s.fg, s.bg)
	ch <- col.Sprint(segmentSeparator)
	close(ch)
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
	col := color.New(s.fg, s.bg)
	t := time.Now()
	ch <- col.Sprintf(" %02d:%02d:%02d ", t.Hour(), t.Minute(), t.Second())
	close(ch)
}
