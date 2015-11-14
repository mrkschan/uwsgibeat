package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
)

// Parser parses stats from uWSGI server.
type Parser interface {
	// Parse stats from the given URL.
	Parse(u url.URL) (map[string]interface{}, error)
}

// StatsParser is a Parser that parses uWSGI stats server.
type StatsParser struct {
}

// NewStatsParser constructs a StatsParser.
func NewStatsParser() Parser {
	return &StatsParser{}
}

// Parse uWSGI stats from the given URL.
func (p *StatsParser) Parse(u url.URL) (map[string]interface{}, error) {
	var reader io.Reader

	switch u.Scheme {
	case "tcp":
		conn, err := net.Dial(u.Scheme, u.Host)
		if err != nil {
			return nil, err
		}
		defer conn.Close()
		reader = conn
	case "unix":
		path := strings.Replace(u.String(), "unix://", "", -1)
		conn, err := net.Dial(u.Scheme, path)
		if err != nil {
			return nil, err
		}
		defer conn.Close()
		reader = conn
	case "http":
		res, err := http.Get(u.String())
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("HTTP%s", res.Status)
		}
		reader = res.Body
	default:
		return nil, fmt.Errorf("%v is a unsupported protocol", u.Scheme)
	}

	// uWSGI stats is expected to be a JSON document.
	// Ref - http://uwsgi-docs.readthedocs.org/en/latest/StatsServer.html
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var payload map[string]interface{}
	if err := json.Unmarshal([]byte(body), &payload); err != nil {
		return nil, err
	}

	// Convert float to int values.
	stats := ftoi(payload)

	return stats, nil
}

// ftoi returns a copy of in where float values are casted to int values.
func ftoi(in map[string]interface{}) map[string]interface{} {
	out := map[string]interface{}{}

	for k, v := range in {
		switch v.(type) {
		case float64:
			vt := v.(float64)
			out[k] = int(vt)
		case map[string]interface{}:
			vt := v.(map[string]interface{})
			out[k] = ftoi(vt)
		case []interface{}:
			vt := v.([]interface{})
			l := len(vt)
			a := make([]interface{}, l)
			for i := 0; i < l; i++ {
				e := vt[i]
				switch e.(type) {
				case float64:
					et := e.(float64)
					a[i] = int(et)
				case map[string]interface{}:
					et := e.(map[string]interface{})
					a[i] = ftoi(et)
				default:
					a[i] = e
				}
			}
			out[k] = a
		default:
			out[k] = v
		}
	}

	return out
}
