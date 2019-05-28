package carDirectory

import (
	"os"
	"strconv"
	"strings"

	"github.com/bullettrain-sh/bullettrain-go-core/src/ansi"
)

const (
	carPaint        = "white:blue"
	separatorSymbol = "/"
	ellipsisSymbol  = "*"
	depthIndicator  = "..."
)

// Directory car
type Car struct {
	paint string
	// Current directory
	Pwd string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_DIRECTORY_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	s := true
	if e := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_SHOW"); e == "false" {
		s = false
	}

	return s
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out)

	out <- ansi.Color(rebuildDirForRender(c.Pwd), c.GetPaint())
}

// rebuildDirForRender contains the main logic to rebuild the current path
// according to the env variable settings, returning the final string version
// that is ready to by printed to the screen.
func rebuildDirForRender(directory string) string {
	var sep string
	if sep = os.Getenv("BULLETTRAIN_CAR_DIRECTORY_PATH_SEPARATOR"); sep == "" {
		sep = separatorSymbol
	}

	if strings.HasPrefix(directory, os.Getenv("HOME")) {
		directory = strings.Replace(directory, os.Getenv("HOME"), "~", 1)
	}

	directoryParts := strings.Split(directory, string(os.PathSeparator))
	l := len(directoryParts)

	frontMaxLength := 2
	if e := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_FRONT_MAX_LENGTH"); e != "" {
		if ml, err := strconv.ParseInt(e, 10, 32); err == nil {
			frontMaxLength = int(ml)
		}
	}

	tailMaxLength := 2
	if e := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_TAIL_MAX_LENGTH"); e != "" {
		if ml, err := strconv.ParseInt(e, 10, 32); err == nil {
			tailMaxLength = int(ml)
		}
	}

	if l > frontMaxLength+tailMaxLength && frontMaxLength > 0 && tailMaxLength > 0 {
		var partsReconstruct []string

		acronymMode := "acronym"
		if e := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_ACRONYM_MODE"); e != "" {
			acronymMode = e
		}

		partsReconstruct = append(partsReconstruct, directoryParts[0:frontMaxLength]...)

		switch acronymMode {
		case "merge":
			var di string
			if di = os.Getenv("BULLETTRAIN_CAR_DIRECTORY_DEPTH_INDICATOR"); di == "" {
				di = depthIndicator
			}
			partsReconstruct = append(partsReconstruct, di)
		default:
			var es string
			if es = os.Getenv("BULLETTRAIN_CAR_DIRECTORY_ELLIPSIS"); es == "" {
				es = ellipsisSymbol
			}

			for _, part := range directoryParts[frontMaxLength : len(directoryParts)-tailMaxLength] {
				if len(part) > 1 {
					partsReconstruct = append(partsReconstruct, part[:1]+es)
				} else {
					partsReconstruct = append(partsReconstruct, part)
				}
			}
		}

		partsReconstruct = append(partsReconstruct, directoryParts[len(directoryParts)-tailMaxLength:]...)

		return strings.Join(partsReconstruct, sep)
	} else {
		return strings.Join(directoryParts, sep)
	}
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_DIRECTORY_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_DIRECTORY_SEPARATOR_SYMBOL")
}

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
	return os.Getenv("BULLETTRAIN_CAR_DIRECTORY_SEPARATOR_TEMPLATE")
}
