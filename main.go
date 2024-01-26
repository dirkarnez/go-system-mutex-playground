package main

import (
	"fmt"
	"time"

	"github.com/juju/mutex"
)

type fakeClock struct {
	delay time.Duration
}

func (f *fakeClock) After(time.Duration) <-chan time.Time {
	return time.After(f.delay)
}

func (f *fakeClock) Now() time.Time {
	return time.Now()
}

func main() {
	// this function is blocking
	r, err := mutex.Acquire(mutex.Spec{Name: "test", Clock: &fakeClock{time.Millisecond}, Delay: 10 * time.Second})
	if err != nil {
		fmt.Println("d")
	}
	defer r.Release()

	fmt.Scanln()
}
