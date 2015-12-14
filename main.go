package main

import (
	"fmt"
)

var (
	build     string
	buildDate string
)

func main() {
	fmt.Printf("Drone Azure Storage Plugin built at %s\n", buildDate)
}
