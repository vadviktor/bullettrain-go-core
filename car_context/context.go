package car_context

import (
	"fmt"
	"os"
	"os/user"

	"github.com/mgutz/ansi"
)

type Context struct {
	paint string
}

func (t *Context) GetPaint() string {
	if t.paint = os.Getenv("BULLETTRAIN_CAR_CONTEXT_PAINT"); t.paint == "" {
		t.paint = "black:white"
	}

	return t.paint
}

func (t *Context) CanShow() bool {
	s := true
	if e := os.Getenv("BULLETTRAIN_CAR_CONTEXT_SHOW"); e == "false" {
		s = false
	}

	return s
}

func (t *Context) Render(out chan<- string) {
	defer close(out)

	var username string
	u, e := user.Current()
	if e == nil {
		username = u.Username
	} else {
		username = "---"
	}

	hostname, e := os.Hostname()
	if e != nil {
		hostname = "---"
	}

	out <- ansi.Color(fmt.Sprintf(" %s@%s ", username, hostname),
		t.GetPaint())
}
