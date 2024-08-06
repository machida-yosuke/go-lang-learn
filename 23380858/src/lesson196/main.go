package main

import (
	"fmt"
	"time"
)

func or(chans ...<-chan interface{}) <-chan interface{} {
	switch len(chans) {
	case 0:
		return nil
	case 1:
		return chans[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(chans) {
		case 2:
			select {
			case <-chans[0]:
			case <-chans[1]:
			}
		default:
			select {
			case <-chans[0]:
			case <-chans[1]:
			case <-chans[2]:
			case <-or(append(chans[3:], orDone)...):
			}
		}
	}()

	return orDone
}

func signal(after time.Duration) <-chan interface{} {
	done := make(chan interface{})
	go func() {
		defer close(done)
		time.Sleep(after)
	}()
	return done
}

func main() {
	start := time.Now()
	<-or(signal(time.Hour), signal(time.Minute), signal(time.Second))
	fmt.Printf("done after %v\n", time.Since(start))
}
