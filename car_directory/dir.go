package car_directory

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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

	cmd := exec.Command("pwd", "-P")
	pwd, err := cmd.Output()
	var d string
	if err == nil {
		ps := string(os.PathSeparator)
		d = strings.Trim(string(pwd), "\n")
		e := strings.Split(d, ps)
		if len(e) > 4 {
			p := e[len(e)-3:]
			d = fmt.Sprintf("...%s", strings.Join(p, ps))
		}
	} else {
		d = "---"
	}

	out <- ansi.Color(fmt.Sprintf(" %s ", d), t.GetPaint())
}
