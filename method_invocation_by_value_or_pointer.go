package main
import "sync"

var a int
// CriticalSection receives a copy of mutex .
func CriticalSection(m sync.Mutex)() {
   m.Lock ()
   a++
   m.Unlock ()
}

func main () {
   a = 0
   mutex := sync.Mutex{}
   // passes a copy of m to A .
   go CriticalSection(mutex)
   go CriticalSection(mutex)
}
