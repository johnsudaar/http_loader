package probe

import (
	"net"
	"net/http"
	"time"
)

const (
	NO_ERROR = iota
	TIMEOUT  = iota
	OTHER    = iota
)

type Probe struct {
	Url      string
	Duration time.Duration
	TimeOut  int
	Err      int
	RespChan chan bool
}

func NewProbe(url string, timeout int, channel chan bool) *Probe {
	return &Probe{
		Url:      url,
		Duration: time.Duration(0),
		TimeOut:  timeout,
		Err:      NO_ERROR,
		RespChan: channel,
	}
}

func (p *Probe) Start() {

	client := http.Client{
		Timeout: time.Duration(p.TimeOut) * time.Second,
	}
	start := time.Now()
	r, err := client.Get(p.Url)
	end := time.Now()

	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			p.Err = TIMEOUT
		} else {
			p.Err = OTHER
		}
	} else {
		defer r.Body.Close()
		p.Duration = end.Sub(start)
	}
	p.RespChan <- true
}
