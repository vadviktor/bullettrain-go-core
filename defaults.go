package main

import (
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/date"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/directory"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/host"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/os"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/status"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/time"
	"github.com/bullettrain-sh/bullettrain-go-core/src/car/user"
	"github.com/bullettrain-sh/bullettrain-go-git"
	"github.com/bullettrain-sh/bullettrain-go-golang"
	"github.com/bullettrain-sh/bullettrain-go-nodejs"
	"github.com/bullettrain-sh/bullettrain-go-openvpn"
	"github.com/bullettrain-sh/bullettrain-go-php"
	"github.com/bullettrain-sh/bullettrain-go-python"
	"github.com/bullettrain-sh/bullettrain-go-ruby"
	"github.com/devkanro/bullettrain-go-k8s"
)

const (
	defaultCarOrder = "os k8s time date user host dir python go ruby nodejs php git status"
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
		"php":     &carPhp.Car{Pwd: currentWorkingDir},
		"python":  &carPython.Car{Pwd: currentWorkingDir},
		"ruby":    &carRuby.Car{Pwd: currentWorkingDir},
		"status":  &carStatus.Car{},
		"openvpn": &carOpenvpn.Car{},
		"time":    &carTime.Car{},
		"k8s":     &carK8s.Car{},
	}
}
