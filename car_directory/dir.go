package carDirectory

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bullettrain-sh/bullettrain-go-core/ansi"
)

const carPaint = "white:blue"

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

	dir := c.Pwd

	if os.Getenv("HOME") == dir {
		dir = "~"
	} else {
		ps := string(os.PathSeparator)

		// Calculate max directory elements to display.
		max_length := 3
		if e := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT"); e != "" {
			ml, err := strconv.Atoi(e)
			if err == nil {
				max_length = ml
			}
		}

		// Allow to override the default three dots by some other string.
		depth_indicator := "..."
		di, di_defined := os.LookupEnv("BULLETTRAIN_CAR_DIRECTORY_DEPTH_INDICATOR")
		if di_defined {
			depth_indicator = di
		}

		// Compose directory segments.
		dirs := strings.Split(dir, ps)
		if max_length > 0 && len(dirs) > max_length+1 {
			f := len(dirs) - max_length
			p := dirs[f:]
			dir = fmt.Sprintf("%s%s", depth_indicator, strings.Join(p, ps))
		}

		if s := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_PATH_SEPARATOR"); s != "" {
			dir = strings.Replace(dir, ps, s, -1)
		}
	}

	out <- ansi.Color(fmt.Sprintf("%s", dir), c.GetPaint())
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
