package main

import (
	"fmt"

	"github.com/rushins/port"
)

func main() {
	fmt.Println("Port Scanning")
	results := port.InitialScan("localhost")
	fmt.Println(results)
}
