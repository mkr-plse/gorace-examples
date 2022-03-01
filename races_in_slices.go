package main
import "sync"

func Foo(x string)(string) {
   return x
}

func ProcessAll(uuids []string)(int) {
  var myResults []string
  var mutex sync.Mutex
  safeAppend := func(res string) {
     mutex.Lock()
     myResults = append(myResults, res)
     mutex.Unlock()
  }

  for _, uuid := range uuids {
     go func (id string, results []string) {
        res := Foo(id)
        safeAppend(res)
     }(uuid, myResults) // slice read without holding lock
  }
  return 0
}

func main() {
  uuids := [2]string{"abcd", "efgh"}
  _ = ProcessAll(uuids[:])
}
