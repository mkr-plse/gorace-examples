package main
import "errors"

// TODO: fix this code

func Start(f * Future)() {
  go func () {
    resp,err := f.f() // invoke a registered function
    f.response = resp
    f.err = err
    f.ch <- 1 // may block forever !
  }()
}

func (f *Future ) Wait(ctx context.Context) error {
  select {
    case <-f.ch:
     return nil
    case <- ctx.Done() :
      f.err = ErrCancelled
      return ErrCancelled
  }
}
