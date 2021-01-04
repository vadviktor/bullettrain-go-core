package carPython

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/bullettrain-sh/bullettrain-go-core/src/ansi"
)

const (
	carPaint              = "black:220"
	pythonSymbolPaint     = "32:220"
	pythonSymbolIcon      = "\uE235"           // 
	virtualenvSymbolIcon  = "\xf0\x9f\x90\x8d" // 🐍
	virtualenvSymbolPaint = "32:220"
	// language=GoTemplate
	carTemplate = `{{.VersionIcon | printf "%s " | cs}}{{.Version | printf "%s " | c}}{{.VenvIcon | printf "%s " | cvs}}{{.Venv | c}}`
)

// Car for Python and virtualenv
type Car struct {
	paint string
	// Current directory
	Pwd string
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_PYTHON_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	if e := os.Getenv("BULLETTRAIN_CAR_PYTHON_SHOW"); e == "true" {
		return true
	}

	var d string
	if d = c.Pwd; d == "" {
		return false
	}

	// Show when .py files exist in current directory
	pyPattern := fmt.Sprintf("%s%s*.py", d, string(os.PathSeparator))
	pyFiles, _ := filepath.Glob(pyPattern)
	if pyFiles != nil {
		return true
	}

	// Show when .python-version file exist in current directory
	versionFiles, _ := filepath.Glob(fmt.Sprintf("%s%s.python-version",
		d, string(os.PathSeparator)))
	if versionFiles != nil {
		return true
	}

	return false
}

// getPythonVersion gets the available version number for a python executable
//
// Use it to check if python2 responds, python3 responds or only python does.
func getPythonVersion(pythonExecutable string) string {
	cmdPython := exec.Command(pythonExecutable, "--version")
	resultPython, errPython := cmdPython.CombinedOutput()
	if errPython == nil {
		return strings.TrimSpace(strings.TrimLeft(
			string(resultPython), "Python "))
	} else {
		return ""
	}
}

func collectPythonVersions() []string {
	pythonVersions := make([]string, 0)

	var p string
	if p = getPythonVersion("python2"); p != "" {
		pythonVersions = append(pythonVersions, p)
	}
	if p = getPythonVersion("python3"); p != "" {
		pythonVersions = append(pythonVersions, p)
	}
	if len(pythonVersions) == 0 {
		if p = getPythonVersion("python"); p != "" {
			pythonVersions = append(pythonVersions, p)
		}
	}

	return pythonVersions
}

func pythonVirtualenv() string {
	if e := os.Getenv("VIRTUAL_ENV"); e != "" {
		return path.Base(e)
	}

	return ""
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
//
// Car version managers can expose multiple Python versions too.
// Python version managers analyzed first, then system Pythons are looked at.
// Empty string is returned when no interpreter could be reached.
func (c *Car) Render(out chan<- string) {
	defer close(out) // Always close the channel!

	var ps string
	if ps = os.Getenv("BULLETTRAIN_CAR_PYTHON_SYMBOL_ICON"); ps == "" {
		ps = pythonSymbolIcon
	}

	var ssp string
	if ssp = os.Getenv("BULLETTRAIN_CAR_PYTHON_SYMBOL_PAINT"); ssp == "" {
		ssp = pythonSymbolPaint
	}

	var vs string
	if vs = os.Getenv("BULLETTRAIN_CAR_PYTHON_VIRTUALENV_SYMBOL_ICON"); vs == "" {
		vs = virtualenvSymbolIcon
	}

	var vsp string
	if vsp = os.Getenv("BULLETTRAIN_CAR_PYTHON_VIRTUALENV_SYMBOL_PAINT"); vsp == "" {
		vsp = virtualenvSymbolPaint
	}

	var s string
	if s = os.Getenv("BULLETTRAIN_CAR_PYTHON_TEMPLATE"); s == "" {
		s = carTemplate
	}

	funcMap := template.FuncMap{
		// Pipeline functions for colouring.
		"c":   func(t string) string { return ansi.Color(t, c.GetPaint()) },
		"cs":  func(t string) string { return ansi.Color(t, ssp) },
		"cvs": func(t string) string { return ansi.Color(t, vsp) },
	}

	tpl := template.Must(template.New("python").Funcs(funcMap).Parse(s))
	data := struct {
		VersionIcon string
		Version     string
		VenvIcon    string
		Venv        string
	}{
		VersionIcon: pythonSymbolIcon,
		Version:     strings.Join(collectPythonVersions(), " "),
		VenvIcon:    virtualenvSymbolIcon,
		Venv:        pythonVirtualenv(),
	}
	fromTpl := new(bytes.Buffer)
	err := tpl.Execute(fromTpl, data)
	if err != nil {
		log.Fatalf("Can't generate the python template: %s", err.Error())
	}

	out <- fromTpl.String()
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_PYTHON_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_PYTHON_SEPARATOR_SYMBOL")
}

// GetSeparatorTemplate overrides the template of the right hand side
// separator through ENV variable.
func (c *Car) GetSeparatorTemplate() string {
	return os.Getenv("BULLETTRAIN_CAR_PYTHON_SEPARATOR_TEMPLATE")
}
