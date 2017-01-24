package main

import "github.com/johnsudaar/http_loader/flooder"

func main() {
	flooder.Launch("http://ensiie-test-1.scalingo.io/?min=4500", 100)
}
