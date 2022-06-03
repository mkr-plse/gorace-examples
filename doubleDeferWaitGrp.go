package main

import "sync"

var racyAccess int

func cleanup() {
	racyAccess++
}

func work(wg *sync.WaitGroup) {
	defer cleanup()
	defer wg.Done() // LIFO order of defer will prematurely release the wg below.
	// some work
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go work(&wg)
	wg.Wait() // will unblock before work() finishes cleanup()
	_ = racyAccess
}
