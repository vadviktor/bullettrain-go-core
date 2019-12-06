package carGit

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/bullettrain-sh/bullettrain-go-core/src/ansi"
)

const (
	carPaint       = "black:white"
	gitSymbolPaint = "red:white"
	gitSymbolIcon  = ""
	gitDirtyPaint  = "red:white"
	gitDirtyIcon   = "✘"
	gitCleanPaint  = "green:white"
	gitCleanIcon   = "✔"
	carTemplate    = `{{.Icon | printf "%s " | cs}}{{.Name | c}}{{.StatusIcon | printf " %s" | csi}}`
)

// Car for Git
type Car struct {
	paint string
	// Current directory
	Pwd string
}

func statusIconInfo(pwd string) (symbol, colour string) {
	var dirtyIcon string
	if dirtyIcon = os.Getenv("BULLETTRAIN_CAR_GIT_DIRTY_ICON"); dirtyIcon == "" {
		dirtyIcon = gitDirtyIcon
	}

	var dirtyPaint string
	if dirtyPaint = os.Getenv("BULLETTRAIN_CAR_GIT_DIRTY_PAINT"); dirtyPaint == "" {
		dirtyPaint = gitDirtyPaint
	}

	var cleanIcon string
	if cleanIcon = os.Getenv("BULLETTRAIN_CAR_GIT_CLEAN_ICON"); cleanIcon == "" {
		cleanIcon = gitCleanIcon
	}

	var cleanPaint string
	if cleanPaint = os.Getenv("BULLETTRAIN_CAR_GIT_CLEAN_PAINT"); cleanPaint == "" {
		cleanPaint = gitCleanPaint
	}

	cmd := exec.Command("git", "-C", pwd, "status", "--porcelain")
	out, err := cmd.Output()
	if err != nil {
		return "", ""
	}

	if len(out) > 0 {
		return dirtyIcon, dirtyPaint
	} else {
		return cleanIcon, cleanPaint
	}
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_GIT_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	cmd := exec.Command("git", "-C", c.Pwd, "rev-parse", "--git-dir")
	cmdOut, _ := cmd.Output()
	if string(cmdOut) != "" {
		return true
	}

	return false
}

func currentHeadName(pwd string) string {
	cmd := exec.Command("git", "-C", pwd, "symbolic-ref", "HEAD")
	ref, err := cmd.Output()
	if err != nil {
		cmd := exec.Command("git", "-C", pwd, "describe", "--tags", "--exact-match", "HEAD")
		ref, err = cmd.Output()
	}
	if err != nil {
		cmd := exec.Command("git", "-C", pwd, "rev-parse", "--short", "HEAD")
		ref, err = cmd.Output()
	}
	if err != nil {
		return strings.TrimRight(err.Error(), "\n")
	}

	ref = []byte(strings.Replace(string(ref), "refs/heads/", "", 1))

	if len(ref) == 0 {
		return ""
	}

	return strings.TrimRight(string(ref), "\n")
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out) // Always close the channel!

	var symbolIcon string
	if symbolIcon = os.Getenv("BULLETTRAIN_CAR_GIT_SYMBOL_ICON"); symbolIcon == "" {
		symbolIcon = gitSymbolIcon
	}

	var symbolPaint string
	if symbolPaint = os.Getenv("BULLETTRAIN_CAR_GIT_SYMBOL_PAINT"); symbolPaint == "" {
		symbolPaint = gitSymbolPaint
	}

	statusIcon, statusIconColour := statusIconInfo(c.Pwd)

	var s string
	if s = os.Getenv("BULLETTRAIN_CAR_GIT_TEMPLATE"); s == "" {
		s = carTemplate
	}

	funcMap := template.FuncMap{
		// Pipeline functions for colouring.
		"c":   func(t string) string { return ansi.Color(t, c.GetPaint()) },
		"cs":  func(t string) string { return ansi.Color(t, symbolPaint) },
		"csi": func(t string) string { return ansi.Color(t, statusIconColour) },
	}

	tpl := template.Must(template.New("user").Funcs(funcMap).Parse(s))
	data := struct {
		Icon       string
		Name       string
		StatusIcon string
	}{Icon: symbolIcon,
		Name:       currentHeadName(c.Pwd),
		StatusIcon: statusIcon}
	fromTpl := new(bytes.Buffer)
	err := tpl.Execute(fromTpl, data)
	if err != nil {
		log.Fatalf("Can't generate the user template: %s", err.Error())
	}

	out <- fromTpl.String()
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_GIT_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_GIT_SEPARATOR_SYMBOL")
}

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
	return os.Getenv("BULLETTRAIN_CAR_GIT_SEPARATOR_TEMPLATE")
}
