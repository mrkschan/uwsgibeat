package main

import (
	"net/url"
	"time"

	"github.com/elastic/libbeat/beat"
	"github.com/elastic/libbeat/cfgfile"
	"github.com/elastic/libbeat/common"
	"github.com/elastic/libbeat/logp"
	"github.com/elastic/libbeat/publisher"

	"github.com/mrkschan/uwsgibeat/parser"
)

const selector = "uwsgibeat"

// Uwsgibeat implements Beater interface and sends uWSGI status using libbeat.
type Uwsgibeat struct {
	// UbConfig holds configurations of Uwsgibeat parsed by libbeat.
	UbConfig ConfigSettings

	done   chan uint
	events publisher.Client

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
	ub.events = b.Events
	ub.done = make(chan uint)

	return nil
}

// Run Uwsgibeat.
func (ub *Uwsgibeat) Run(b *beat.Beat) error {
	logp.Debug(selector, "Run uwsgibeat")

	p := parser.NewStatsParser()

	ticker := time.NewTicker(ub.period)
	defer ticker.Stop()

	for {
		select {
		case <-ub.done:
			goto GotoFinish
		case <-ticker.C:
		}

		start := time.Now()

		s, err := p.Parse(*ub.url)
		if err != nil {
			logp.Err("Fail to read uWSGI stats: %v", err)
			goto GotoNext
		}
		ub.events.PublishEvent(common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":        "uwsgi",
			"uwsgi":       s,
		})

	GotoNext:
		end := time.Now()
		duration := end.Sub(start)
		if duration.Nanoseconds() > ub.period.Nanoseconds() {
			logp.Warn("Ignoring tick(s) due to processing taking longer than one period")
		}
	}

GotoFinish:
	return nil
}

// Cleanup Uwsgibeat.
func (ub *Uwsgibeat) Cleanup(b *beat.Beat) error {
	return nil
}

// Stop Uwsgibeat.
func (ub *Uwsgibeat) Stop() {
	logp.Debug(selector, "Stop uwsgibeat")
	close(ub.done)
}
