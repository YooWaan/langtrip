package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	value := make(chan bool)
	done := make(chan bool)

	go func() {
		defer func() { done <- true }()
		num := 0
		for {
			select {
			case <-value:
				num++
				if num >= 10 {
					return
				} else {
					fmt.Println("wait...[", num, "]")
				}
			}
		}
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t.Format("15:04:05"))
			value <- true
		}
	}
}
