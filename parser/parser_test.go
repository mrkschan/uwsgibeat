package parser

import (
	"testing"

	"fmt"
	"net"
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
}
