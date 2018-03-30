# How to create a new car in Go

This document will describe the process to add new car plugin in Go.

## Naming

The car needs to have a fitting package name that matches the current
convention: `carSomeName` The name is prefixed by `car`.

## Template:

You can have a look at the very simple
[timeCar](../src/car/time/time.go), it shows basic implementation of the
`carRenderer` interface.

When you may want to just permanently change symbol or paint colours,
you can simply change the constants on the top level and compile your
custom version:

```go
const (
	carPaint    = "black:white"
	symbolIcon  = ""
	symbolPaint = "black:white"
	carTemplate = `{{.Icon | printf "%s " | cs}}{{.Time | c}}`
)
```

To use your car, this is a checklist to be done in
[bullettrain-go-core/defaults.go](../defaults.go):

* add you package path to the imports
* add your car's instance to the `trailers` function
* add your car's callword to the list in the `defaultCarOrder` constant

## Documenting cars

The `README` should also be as detailed as the ones already existing.
Each and every env variable must be documented with default values and
description.
