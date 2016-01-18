package main

import (
	"peeple/dynip/dynip"
	"time"
)

func main() {

	dynip := &dynip.NameCheap{
		DomainName:   "peeple.es",
		Password:     "2cc80a9e6f864ea9973b0d8c521e26e9",
		UpdatingTime: 60 * time.Second,
		VerifyChange: false,
	}

	dynip.Execute()
}
