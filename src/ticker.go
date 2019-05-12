package main

import (
	"fmt"
	"time"
)

func main() {
	value := make(chan int)
	score := 0

	// score updater
	go func() {
		for {
			select {
			case n, ok := <-value:
				if ok {
					score += n
				} else {
					score *= 2
					return
				}
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	loop := 0
	loopEnd := 5
outer:
	for {
		select {
		case t := <-ticker.C:
			if loop >= loopEnd {
				if loop > loopEnd {
					break outer
				} else {
					loop++
					close(value)
				}
			} else {
				loop++
				fmt.Println("Current time: ", t.Format("15:04:05"))
				value <- loop
			}
		}
	}

	fmt.Println("score:", score)
}
