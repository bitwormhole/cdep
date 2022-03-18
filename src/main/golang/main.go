package main

import (
	"github.com/bitwormhole/cdep"
	"github.com/bitwormhole/starter"
)

func main() {
	i := starter.InitApp()
	i.Use(cdep.Module())
	i.Run()
}
