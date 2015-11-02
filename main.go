package main

import (
	"github.com/elastic/libbeat/beat"
)

var Version = "1.0.0-beta1"
var Name = "uwsgibeat"

func main() {
	ub := &UWSGIbeat{}

	b := beat.NewBeat(Name, Version, ub)

	b.CommandLineSetup()

	b.LoadConfig()
	ub.Config(b)

	b.Run()
}
