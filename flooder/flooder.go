package flooder

import (
	"fmt"
	"time"

	"github.com/johnsudaar/http_loader/probe"
)

var totalRequest = 0
var totalTimeout = 0
var totalMs = int64(0)
var totalReq = 0
var lastDisplay = time.Now()

func Launch(url string, count int) {
	done := make(chan bool, count)
	probes := make([]*probe.Probe, count)
	for i := 0; i < count; i++ {
		probes[i] = probe.NewProbe(url, 2, done)
	}

	for i := 0; i < count; i++ {
		go probes[i].Start()
	}

	for i := 0; i < count; i++ {
		if i%(count/10) == 0 {
			//fmt.Println("Received : " + strconv.Itoa(i))
		}
		<-done
	}

	errored := 0
	timeout := 0
	sum := int64(0)
	good := 0

	for i := 0; i < count; i++ {
		if probes[i].Err != probe.NO_ERROR {
			errored++
		} else {
			good++
			sum += probes[i].Duration.Nanoseconds()
		}
		if probes[i].Err == probe.TIMEOUT {
			timeout++
		}
	}

	if good == 0 {
		//fmt.Println("Time : NO DATA AVAILABLE")
	} else {
		mean := sum / int64(good*1000000)
		totalReq++
		totalMs += mean
		//fmt.Println("Time : " + strconv.FormatInt(mean, 10) + " ms")
	}
	totalTimeout += timeout
	totalRequest += count

	if time.Now().Sub(lastDisplay) >= 30*time.Second {
		fmt.Println("ms")
		fmt.Println(totalMs)
		fmt.Println("req")
		fmt.Println(totalReq)
		fmt.Println("request")
		fmt.Println(totalRequest)
		fmt.Println("timeout")
		fmt.Println(totalTimeout)
		totalMs = 0
		totalReq = 0
		totalRequest = 0
		totalTimeout = 0
		lastDisplay = time.Now()
	}
	//fmt.Println("Timeout : " + strconv.Itoa(timeout))
	//fmt.Println("Errored : " + strconv.Itoa(errored))
}
