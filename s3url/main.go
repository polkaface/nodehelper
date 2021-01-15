package main

import (
	"sync"

	"github.com/cz-theng/czkit-go/log"
)

func main() {
	log.Init(log.WithLogPath("./"), log.WithLogName("s3url.log"), log.WithConsole(false))
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go watchdog()
	server.Init()
	server.Start()

	return
}
