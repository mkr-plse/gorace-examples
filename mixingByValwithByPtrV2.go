package main

import "sync"
import "time"

var m sync.Mutex

func main() {
	obj := myObject{a: 42}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		obj.writeA(1) // updates obj.a
		wg.Done()
	}()
	// intentionally sleep to expose the racy interleaving
	time.Sleep(1000 * time.Millisecond)
	_ = obj.readA() // Reads obj.a
	wg.Wait()
}

type myObject struct {
	a int
}

// By-ptr receiver
func (f *myObject) writeA(value int) {
	m.Lock()
	f.a = value
	m.Unlock()
}

// By-value receiver
func (f myObject) readA() int {
	m.Lock()
	defer m.Unlock()
	return f.a
}
