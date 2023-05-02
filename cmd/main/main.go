package main

import (
	"fmt"

	"github.com/guevarez30/goload/pkg/load"
)

func main() {
	fmt.Println("go-load")
	url := "http://localhost:3001/pie"
	connections := uint32(4)
	requests := uint32(1015)
	load.Hammer(url, connections, requests)
}
