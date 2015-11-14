package parser

import (
	"os"
	"path"
	"testing"

	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/stretchr/testify/assert"
)

func NewTCPServer(network, address string, handler func(c net.Conn)) (net.Listener, error) {
	l, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Printf("ERROR: %v", err)
				return
			}
			handler(conn)
			conn.Close()
		}
	}()

	return l, nil
}

func TestStatsParser(t *testing.T) {
	ts1, err := NewTCPServer("tcp", ":7777", func(c net.Conn) {
		payload := `{
        "workers": [{
          "id": 1,
          "pid": 31759,
          "requests": 0,
          "exceptions": 0,
          "status": "idle",
          "rss": 0,
          "vsz": 0,
          "running_time": 0,
          "last_spawn": 1317235041,
          "respawn_count": 1,
          "tx": 0,
          "avg_rt": 0,
          "apps": [{
            "id": 0,
            "modifier1": 0,
            "mountpoint": "",
            "requests": 0,
            "exceptions": 0,
            "chdir": ""
          }]
        }]}`
		fmt.Fprintln(c, payload)
	})
	if err != nil {
		t.Error(err)
	}
	defer ts1.Close()

	p1 := &StatsParser{}
	u1, _ := url.Parse("tcp://127.0.0.1:7777")
	s1, _ := p1.Parse(*u1)

	assert.NotNil(t, s1["workers"])
	assert.IsType(t, []interface{}{}, s1["workers"])
	ww1 := s1["workers"].([]interface{})
	w1 := ww1[0].(map[string]interface{})

	assert.Equal(t, 1, w1["id"])
	assert.Equal(t, 31759, w1["pid"])
	assert.Equal(t, 0, w1["requests"])
	assert.Equal(t, "idle", w1["status"])
	assert.Equal(t, 1317235041, w1["last_spawn"])
	assert.NotNil(t, w1["apps"])

	ts2, err := NewTCPServer("unix", path.Join(os.TempDir(), "uwsgibeat-ts2.sock"), func(c net.Conn) {
		payload := `{
        "workers": [{
          "id": 1,
          "pid": 31759,
          "requests": 0,
          "exceptions": 0,
          "status": "idle",
          "rss": 0,
          "vsz": 0,
          "running_time": 0,
          "last_spawn": 1317235041,
          "respawn_count": 1,
          "tx": 0,
          "avg_rt": 0,
          "apps": [{
            "id": 0,
            "modifier1": 0,
            "mountpoint": "",
            "requests": 0,
            "exceptions": 0,
            "chdir": ""
          }]
        }]}`
		fmt.Fprintln(c, payload)
	})
	if err != nil {
		t.Error(err)
	}
	defer ts2.Close()

	p2 := &StatsParser{}
	u2, _ := url.Parse(fmt.Sprintf("unix://%s", path.Join(os.TempDir(), "uwsgibeat-ts2.sock")))
	s2, _ := p2.Parse(*u2)

	assert.NotNil(t, s2["workers"])
	assert.IsType(t, []interface{}{}, s2["workers"])
	ww2 := s2["workers"].([]interface{})
	w2 := ww2[0].(map[string]interface{})

	assert.Equal(t, 1, w2["id"])
	assert.Equal(t, 31759, w2["pid"])
	assert.Equal(t, 0, w2["requests"])
	assert.Equal(t, "idle", w2["status"])
	assert.Equal(t, 1317235041, w2["last_spawn"])
	assert.NotNil(t, w2["apps"])

	ts3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := `{
        "workers": [{
          "id": 1,
          "pid": 31759,
          "requests": 0,
          "exceptions": 0,
          "status": "idle",
          "rss": 0,
          "vsz": 0,
          "running_time": 0,
          "last_spawn": 1317235041,
          "respawn_count": 1,
          "tx": 0,
          "avg_rt": 0,
          "apps": [{
            "id": 0,
            "modifier1": 0,
            "mountpoint": "",
            "requests": 0,
            "exceptions": 0,
            "chdir": ""
          }]
        }]}`
		fmt.Fprintln(w, payload)
	}))
	defer ts3.Close()

	p3 := &StatsParser{}
	u3, _ := url.Parse(ts3.URL)
	s3, _ := p3.Parse(*u3)

	assert.NotNil(t, s3["workers"])
	assert.IsType(t, []interface{}{}, s3["workers"])
	ww3 := s3["workers"].([]interface{})
	w3 := ww3[0].(map[string]interface{})

	assert.Equal(t, 1, w3["id"])
	assert.Equal(t, 31759, w3["pid"])
	assert.Equal(t, 0, w3["requests"])
	assert.Equal(t, "idle", w3["status"])
	assert.Equal(t, 1317235041, w3["last_spawn"])
	assert.NotNil(t, w3["apps"])
}
