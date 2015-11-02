package main

import (
	"time"
	"net/url"

	"github.com/elastic/libbeat/beat"
	"github.com/elastic/libbeat/cfgfile"
	"github.com/elastic/libbeat/logp"
)

const selector = "uwsgibeat"

// UWSGIbeat implements Beater interface and sends uWSGI status using libbeat.
type UWSGIbeat struct {
	// UbConfig holds configurations of UWSGIbeat parsed by libbeat.
	UbConfig ConfigSettings

	url    *url.URL
	method string
	period time.Duration
}

// Config UWSGIbeat according to uwsgibeat.yml.
func (ub *UWSGIbeat) Config(b *beat.Beat) error {
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

		if ub.UbConfig.Input.Method != "" {
			ub.method = ub.UbConfig.Input.Method
		} else {
			ub.method = "tcp"
		}

		if ub.UbConfig.Input.Period != nil {
			ub.period = time.Duration(*ub.UbConfig.Input.Period) * time.Second
		} else {
			ub.period = 1 * time.Second
		}

		logp.Debug(selector, "Init uwsgibeat")
		logp.Debug(selector, "Watch %v", ub.url)
		logp.Debug(selector, "Method %v", ub.method)
		logp.Debug(selector, "Period %v", ub.period)

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
