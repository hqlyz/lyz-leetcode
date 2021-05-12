package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	o1 := new(Once)
	for i := 0; i < 1000000; i++ {
		go o1.Do(hello)
	}
}

func hello() {
	fmt.Println("Hello")
}

type Once struct {
	m sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.m.Lock()
		defer o.m.Unlock()
		if o.done == 0 {
			defer atomic.StoreUint32(&o.done, 1)
			f()
		}
	}
}
