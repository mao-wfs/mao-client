package main

import (
	"github.com/mao-wfs/mao-client/external/waf"
)

func main() {
	server := waf.NewServer()
	server.Run()
}
