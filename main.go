package main

import (
	"time"

	"github.com/johnsudaar/http_loader/scenario"
)

func main() {
	s := scenario.NewSinusScenario(20*time.Minute, 5*time.Minute, 75, 20)

	s.Start()
}
