package main

import (
	"github.com/mao-wfs/mao-ctrl/external/waf"
)

func main() {
	server := waf.NewServer()
	server.Run()
}
