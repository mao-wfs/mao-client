package main

import (
	"github.com/mao-wfs/maoctrl/external/waf"
)

func main() {
	server := waf.NewServer()
	server.Run()
}
