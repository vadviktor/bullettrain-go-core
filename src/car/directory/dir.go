package carDirectory

import (
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/bullettrain-sh/bullettrain-go-core/src/ansi"
)

const (
	carPaint        = "white:blue"
	separatorSymbol = ""
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

	d := ""
	if r := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_ROOT_SHOW"); r != "false" {
		d = sep
	}

	if strings.HasPrefix(directory, os.Getenv("HOME")) {
		directory = strings.Replace(directory, os.Getenv("HOME"),
			string(os.PathSeparator)+"~", 1)
	}

	directoryParts := strings.Split(directory, string(os.PathSeparator))
	directoryParts = directoryParts[1:]
	l := len(directoryParts)

	maxLength := 3.0
	if e := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT"); e != "" {
		if ml, err := strconv.ParseFloat(e, 32); err == nil {
			maxLength = ml
		}
	}

	if l > int(maxLength) && maxLength > 0 {
		var partsReconstruct []string

		var firstDir string
		if firstDir = os.Getenv("BULLETTRAIN_CAR_DIRECTORY_FIRST_DIR_SHOW"); firstDir != "false" {
			firstDir = "true"
		}

		if firstDir == "true" && maxLength > 1 {
			head := 2 - int(math.Floor(maxLength/2))
			partsReconstruct = directoryParts[0:head]
		} else if firstDir == "true" && maxLength == 1 {
			partsReconstruct = directoryParts[0:1]
		}

		var di string
		if di = os.Getenv("BULLETTRAIN_CAR_DIRECTORY_DEPTH_INDICATOR"); di == "" {
			di = depthIndicator
		}

		partsReconstruct = append(partsReconstruct, di)

		tail := l - int(math.Ceil(maxLength/2))
		if firstDir == "false" && maxLength > 1 {
			tail -= 1
		}
		partsReconstruct = append(partsReconstruct,
			directoryParts[tail:l]...)

		d = d + strings.Join(partsReconstruct, sep)
	} else {
		d = d + strings.Join(directoryParts, sep)
	}

	return d
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
