package main

import (
	"github.com/oojob/service-profile/src/cmd"
	"github.com/oojob/service-profile/src/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
