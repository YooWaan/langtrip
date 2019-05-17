package main

import (
	"fmt"
	"time"
)

var (
	ErrEndOfInterator = fmt.Errorf("EndOfIterator")
)

type (
	// Event Handlers
	Next  func(interface{})
	Done  func()
	Error func(error)

	// observable
	observable <-chan interface{}

	// iterable
	iterable <-chan interface{}
)

func defaultError(e error) { fmt.Println("onerr:", e) }

// subscription
type subscription struct {
	EndAt time.Time
	Error error
}

func (s subscription) Err() error {
	return s.Error
}

func (s subscription) End() subscription {
	s.EndAt = time.Now()
	return s
}

// Event Handlers
type observer struct {
	next Next
	done Done
	err  Error
}

func handler(n Next, d Done) observer {
	return observer{next: n, done: d, err: defaultError}
}

func (o observable) subscribe(ob observer) <-chan subscription {
	done := make(chan subscription)
	sub := subscription{}

	// parallel unsupported
	go func() {
	OuterLoop:
		for item := range o {
			switch item := item.(type) {
			case error:
				ob.err(item)
				sub.Error = item
				break OuterLoop
			default:
				ob.next(item)
			}
		}

		// OnDone only gets executed if there's no error.
		if sub.Error == nil {
			ob.done()
		}

		done <- sub.End()
		return
	}()

	return done
}

func toItr(any interface{}) (iterable, error) {
	switch ch := any.(type){
	case chan interface{}:
		return iterable(ch), nil
	case <- chan interface{}:
		return iterable(ch), nil
	default:
		return nil, fmt.Errorf("unsupported any")
	}
}

func (it iterable) Next() (interface{}, error) {
	if next, ok := <-it; ok {
		return next, nil
	}
	return nil, ErrEndOfInterator
}

func just(items ...interface{}) observable {
	source := make(chan interface{})
	go func() {
		for _, item := range items {
			source <- item
		}
		close(source)
	}()

	return observable(source)
}

func justSubscribe() {
	score := 0

	observer := handler(
		// next
		func(n interface{}) {
			if num, ok := n.(int); ok {
				score += num
			}
		},
		// done
		func() {
			score *= 2
		},
	)
	data := []interface{}{1, 2, 3, 4, 5}
	sub := just(data...).subscribe(observer)
	<-sub

	fmt.Println("score:", score)
}

func from(it iterable) observable {
	source := make(chan interface{})
	go func() {
		for {
			val, err := it.Next()
			if err != nil {
				break
			}
			source <- val
		}
		close(source)
	}()
	return observable(source)
}

func ticker() {
	score := 0
	observer := handler(
		// next
		func(n interface{}) {
			if num, ok := n.(int); ok {
				score += num
			}
		},
		// done
		func() {
			score *= 2
		},
	)

	value := make(chan interface{})

	itr, err := toItr(value)
	if err != nil {
		fmt.Println("error:", )
		return
	}
	sub := from(itr).subscribe(observer)

	ticker := time.NewTicker(time.Second)
	go func() {
		defer ticker.Stop()
		num := 0
		for {
			select {
			case t := <-ticker.C:
				num++
				fmt.Println("Current time: ", t.Format("15:04:05"), ",num:", num)
				if num > 5 {
					close(value)
					return
				}
				value <- num
			}
		}
	}()

	<- sub
	fmt.Println("score:", score)
}

func main() {
	justSubscribe()
	ticker()
}
