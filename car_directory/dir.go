package carDirectory

import (
	"os"
	"strconv"
	"strings"

	"github.com/bullettrain-sh/bullettrain-go-core/ansi"
)

const (
	carPaint        = "white:blue"
	separatorSymbol = ""
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

	dir := c.Pwd

	if strings.HasPrefix(dir, os.Getenv("HOME")) {
		dir = strings.Replace(dir, os.Getenv("HOME"), "~", 1)
	}

	// Calculate max directory elements to display.
	maxLength := 3
	if e := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT"); e != "" {
		if ml, err := strconv.Atoi(e); err == nil && ml >= 3 {
			maxLength = ml
		}
	}

	// Compose directory segments.
	ps := string(os.PathSeparator)
	dirs := strings.Split(dir, ps)
	if maxLength > 0 && len(dirs) > maxLength+1 {
		newPath := make([]string, 0)
		lastPathIdx := len(dirs) - maxLength

		if s := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_FIRST_DIR_SHOW"); s != "false" {
			newPath = append(newPath, dirs[1])
			lastPathIdx += 1
		}

		lastPath := dirs[lastPathIdx:]

		depthIndicator := "..."
		// Allow to override the default three dots by some other string.
		if di, ex := os.LookupEnv("BULLETTRAIN_CAR_DIRECTORY_DEPTH_INDICATOR"); ex {
			depthIndicator = di
		}

		newPath = append(newPath, depthIndicator)
		newPath = append(newPath, lastPath...)

		dir = joinDirs(newPath)
	} else {
		dir = joinDirs(dirs)
	}

	out <- ansi.Color(dir, c.GetPaint())
}

func joinDirs(dirs []string) string {
	psSymbol := separatorSymbol
	if s := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_PATH_SEPARATOR"); s != "" {
		psSymbol = s
	}

	var p string
	fds := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_FIRST_DIR_SHOW")
	fss := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_FIRST_SEPARATOR_SHOW")
	if fds != "false" && fss == "true" {
		p += psSymbol
	}

	p += strings.Join(dirs, psSymbol)

	return p
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
