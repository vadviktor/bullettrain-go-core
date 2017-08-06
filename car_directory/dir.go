package car_directory

import (
	"fmt"
	"os"

	"path/filepath"

	"github.com/mgutz/ansi"
)

type Directory struct {
	paint string
}

func (t *Directory) GetPaint() string {
	if t.paint = os.Getenv("BULLETTRAIN_CAR_DIRECTORY_PAINT"); t.paint == "" {
		t.paint = "white:blue"
	}

	return t.paint
}

func (t *Directory) CanShow() bool {
	s := true
	if e := os.Getenv("BULLETTRAIN_CAR_DIRECTORY_SHOW"); e == "false" {
		s = false
	}

	return s
}

func (t *Directory) Render(out chan<- string) {
	defer close(out)

	d, err := os.Executable()
	if err == nil {
		d = filepath.Dir(d)
	} else {
		d = "---"
	}

	out <- ansi.Color(fmt.Sprintf("%s", d), t.GetPaint())
}
