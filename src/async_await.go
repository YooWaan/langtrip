package main

import (
	"time"
	"sync"
)

func elapsed(fnName string, start time.Time) {
	println("done " + fnName + ": elapsed", time.Since(start) / time.Millisecond, "msec")

}

func asyncFunc(sleep time.Duration, start time.Time) {
	time.Sleep(sleep)
	elapsed("async", start)
}

func awaitFunc(wg *sync.WaitGroup, sleep time.Duration, start time.Time) {
	defer wg.Done()
	// do somethig
	time.Sleep(sleep)
	elapsed("await", start)
}

func main() {
	start := time.Now()
	nums := []time.Duration{ 2 * time.Second, time.Second, 1500 * time.Millisecond }
	
	// async finished
	for _, sec := range nums {
		go asyncFunc(sec, start)
	}

	// await
	await := new(sync.WaitGroup)
	for _, sec := range nums {
		await.Add(1)
		go awaitFunc(await, sec, start)
	}
	await.Wait()

	println("finish")
}
