package main

import (
	"broker_queues/manager"

	"log"
)

func main() {
	th := manager.NewTaskHub()
	err := th.Run()

	if err != nil {
		log.Fatal(err)
	}
}
