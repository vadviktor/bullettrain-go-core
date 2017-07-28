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

	segmentList := getSegments()
	for _, segment := range segmentList {
		fmt.Print(segment.Render())
	}

	fmt.Printf("\n%s x", color.HiGreenString(lineEnding))
}

type renderer interface {
	Render() string
}

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

type separator struct {
	fg, bg color.Attribute
}

func (s *separator) Render() string {
	const segmentSeparator string = ""
	c := color.New(s.fg, s.bg)
	return c.Sprint(segmentSeparator)
}

type timeSegment struct {
	fg, bg color.Attribute
}

func (s *timeSegment) Render() string {
	c := color.New(s.fg, s.bg)
	t := time.Now()
	return c.Sprintf(" %02d:%02d:%02d ", t.Hour(), t.Minute(), t.Second())
}
