package main

import (
	"fmt"
	"time"
)

func main() {
	timer(time.After(2 * time.Second))
}
func timer(c <-chan time.Time) {
	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
	}
}
