package scenario

import (
	"fmt"
	"math"
	"time"

	"github.com/johnsudaar/http_loader/flooder"
)

type Sinus struct {
	Duration time.Duration
	Period   time.Duration
	Max      int
	Min      int
}

func NewSinusScenario(d, p time.Duration, max, min int) *Sinus {
	return &Sinus{
		Duration: d,
		Period:   p,
		Max:      max,
		Min:      min,
	}
}

func (s *Sinus) Start() {
	deltaA := 2.0 * 3.14 / s.Period.Seconds()
	start := time.Now()

	angle := 0.0

	for time.Now().Sub(start) < s.Duration {
		value := float64(s.Min) + (1+math.Sin(angle))*float64(s.Max-s.Min)/2
		angle += deltaA
		fmt.Print("*")
		go flooder.Launch("http://ensiie-test-1.scalingo.io/?min=2000", int(value))
		time.Sleep(time.Second)
	}
}
