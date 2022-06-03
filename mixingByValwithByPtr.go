package main

import "sync"

func main() {
        obj := myObject{a: 42, b: 24}
        var wg sync.WaitGroup
        wg.Add(1)
        go func() {
                obj.writeA(1) // updates obj.a
                wg.Done()
        }()
        _ = obj.readB() // reads entire obj including a and b because the method receiver is by value.
        wg.Wait()
}

type myObject struct {
        a int
        b int
}

// By-ptr receiver
func (f *myObject) writeA(value int) {
        f.a = value
}

// By-value receiver
func (f myObject) readB() int {
        return f.b
}
