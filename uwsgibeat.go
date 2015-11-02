package main

import (
	"github.com/elastic/libbeat/beat"
)

// UWSGIbeat implements Beater interface and sends uWSGI status using libbeat.
type UWSGIbeat struct {
	config ConfigSettings
}

// Config UWSGIbeat according to uwsgibeat.yml.
func (ub *UWSGIbeat) Config(b *beat.Beat) error {
	return nil
}

// Setup UWSGIbeat.
func (ub *UWSGIbeat) Setup(b *beat.Beat) error {
	return nil
}

// Run UWSGIbeat.
func (ub *UWSGIbeat) Run(b *beat.Beat) error {
	return nil
}

// Cleanup UWSGIbeat.
func (ub *UWSGIbeat) Cleanup(b *beat.Beat) error {
	return nil
}

// Stop UWSGIbeat.
func (ub *UWSGIbeat) Stop() {
}
