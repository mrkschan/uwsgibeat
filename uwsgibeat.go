package main

import (
	"net/url"
	"time"

	"github.com/elastic/libbeat/beat"
	"github.com/elastic/libbeat/cfgfile"
	"github.com/elastic/libbeat/logp"
)

const selector = "uwsgibeat"

// Uwsgibeat implements Beater interface and sends uWSGI status using libbeat.
type Uwsgibeat struct {
	// UbConfig holds configurations of Uwsgibeat parsed by libbeat.
	UbConfig ConfigSettings

	url    *url.URL
	period time.Duration
}

// Config Uwsgibeat according to uwsgibeat.yml.
func (ub *Uwsgibeat) Config(b *beat.Beat) error {
	err := cfgfile.Read(&ub.UbConfig, "")
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return err
	}

	var u string
	if ub.UbConfig.Input.URL != "" {
		u = ub.UbConfig.Input.URL
	} else {
		u = "127.0.0.1:1717"
	}
	ub.url, err = url.Parse(u)
	if err != nil {
		logp.Err("Invalid uWSGI stats server address: %v", err)
		return err
	}

	if ub.UbConfig.Input.Period != nil {
		ub.period = time.Duration(*ub.UbConfig.Input.Period) * time.Second
	} else {
		ub.period = 1 * time.Second
	}

	logp.Debug(selector, "Init uwsgibeat")
	logp.Debug(selector, "Watch %v", ub.url)
	logp.Debug(selector, "Period %v", ub.period)

	return nil
}

// Setup Uwsgibeat.
func (ub *Uwsgibeat) Setup(b *beat.Beat) error {
	return nil
}

// Run Uwsgibeat.
func (ub *Uwsgibeat) Run(b *beat.Beat) error {
	return nil
}

// Cleanup Uwsgibeat.
func (ub *Uwsgibeat) Cleanup(b *beat.Beat) error {
	return nil
}

// Stop Uwsgibeat.
func (ub *Uwsgibeat) Stop() {
}
