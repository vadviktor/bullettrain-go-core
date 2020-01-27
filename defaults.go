package main

import (
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/date"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/directory"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/git"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/golang"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/host"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/kubernetes"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/nodejs"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/openvpn"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/os"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/php"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/python"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/ruby"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/status"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/time"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/user"
)

const (
	defaultCarOrder = "os time date user host dir python go ruby nodejs php git status"
	separatorSymbol = ""
	// language=GoTemplate
	separatorTemplate = `{{.Icon | printf "%s " | c}}`
	// language=GoTemplate
	promptCharTemplate = `{{.Icon | printf "%s " | c}}`
)

// trailers results in the list of cars to be available for use.
func trailers(currentWorkingDir string) map[string]carRenderer {
	return map[string]carRenderer{
		"user":    &carUser.Car{},
		"host":    &carHost.Car{},
		"date":    &carDate.Car{},
		"dir":     &carDirectory.Car{Pwd: currentWorkingDir},
		"git":     &carGit.Car{Pwd: currentWorkingDir},
		"go":      &carGo.Car{Pwd: currentWorkingDir},
		"nodejs":  &carNodejs.Car{Pwd: currentWorkingDir},
		"os":      &carOs.Car{},
		"status":  &carStatus.Car{},
		"openvpn": &carOpenvpn.Car{},
		"time":    &carTime.Car{},
		"php":     &carPhp.Car{Pwd: currentWorkingDir},
		"python":  &carPython.Car{Pwd: currentWorkingDir},
		"ruby":    &carRuby.Car{Pwd: currentWorkingDir},
		"k8s":     &carK8s.Car{},
	}
}
