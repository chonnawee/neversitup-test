package main

import (
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(5 * time.Second)
		}
	}()

	select {}
}
