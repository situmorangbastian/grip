package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/situmorangbastian/grip/cmd"
)

func main() {
	if err := cmd.RootCMD.Execute(); err != nil {
		log.Fatal("Fail init Root CMD with error: ", err)
	}
}
