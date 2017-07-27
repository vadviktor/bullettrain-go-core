package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	color.NoColor = false // force terminal to use colours
	const line_ending = "$"

	fmt.Printf("%s%s%s%s\n%s x",
		time_segment(color.FgHiWhite, color.BgBlack),
		separator(color.FgBlack, color.BgYellow),
		python_segment(color.FgHiWhite, color.BgYellow),
		separator(color.FgYellow, color.FgYellow),
		color.HiGreenString(line_ending),
	)
}

func separator(fg, bg color.Attribute) string {
	const segment_separator string = ""
	c := color.New(fg, bg)
	return c.Sprint(segment_separator)
}

func time_segment(fg, bg color.Attribute) string {
	c := color.New(fg, bg)
	t := time.Now()
	return c.Sprintf(" %02d:%02d:%02d ", t.Hour(), t.Minute(), t.Second())
}

// Builds the version string of the currently available Python interpreter(s).
// Python version managers can expose multiple versions too.
// Version managers analyzed first, then system Pythons.
// Empty string is returned when no interpreter could be reached.
func python_segment(fg, bg color.Attribute) string {
	const python_symbol string = "🐍"
	c := color.New(fg, bg)

	// ______
	// | ___ \
	// | |_/ /   _  ___ _ ____   __
	// |  __/ | | |/ _ \ '_ \ \ / /
	// | |  | |_| |  __/ | | \ V /
	// \_|   \__, |\___|_| |_|\_/
	//        __/ |
	//       |___/

	pyenvCmd := exec.Command("pyenv", "version")
	pyenvOut, err := pyenvCmd.Output()
	if err == nil {
		re := regexp.MustCompile(`(?m)^([a-zA-Z0-9_\-]+)`)
		versions := re.FindAllStringSubmatch(string(pyenvOut), -1)
		var versions_info string
		for _, i := range versions {
			versions_info = fmt.Sprintf("%s %s", versions_info, i[1])
		}

		return c.Sprintf(" %s%s ", python_symbol, versions_info)
	}

	// ______      _   _
	// | ___ \    | | | |
	// | |_/ /   _| |_| |__   ___  _ __
	// |  __/ | | | __| '_ \ / _ \| '_ \
	// | |  | |_| | |_| | | | (_) | | | |
	// \_|   \__, |\__|_| |_|\___/|_| |_|
	//        __/ |
	//       |___/

	// TODO python 2 and python 3 version info!

	pythonCmd := exec.Command("python", "--version")
	var stderr bytes.Buffer
	pythonCmd.Stderr = &stderr
	pyErr := pythonCmd.Run()
	if pyErr == nil {
		return c.Sprintf(" %s %s ",
			python_symbol, strings.Trim(stderr.String(), "\n"))
	}

	return ""
}
