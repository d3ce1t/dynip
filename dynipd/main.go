package main

import (
	"log"
	"peeple/dynip/dynip"
	"time"
)

func main() {

	// Load config from file
	cfg, err := loadConfigFromFile("config.yaml")
	if err != nil {
		log.Fatalf("Couldn't load config.yaml: %v", err)
	}

	dynip := &dynip.NameCheap{
		SubDomainName: cfg.SubDomainName,
		DomainName:    cfg.DomainName,
		Password:      cfg.APIPassword,
		UpdatingTime:  time.Duration(cfg.UpdateFreqSec) * time.Second,
		VerifyChange:  cfg.VerifyChange,
	}

	dynip.Execute()
}
