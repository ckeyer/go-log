package main

import (
	logging "github.com/ckeyer/go-log"
)

//
func main() {
	log := logging.GetDefaultLogger()

	log.Debug("heelosdf")
	log.Error("heleo")
	log.Info("helelo")
	log.Warning("hhhhh")
}
