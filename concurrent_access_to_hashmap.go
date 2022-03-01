package main
import "errors"

func Foo(x string)(string,error) {
   return x,errors.New("random error")
}

func combineErrors(errMap map[string] error)(int) {
   return 10
}

func processOrders(uuids []string)(int) {
   var errMap = make(map[string] error)
   for _,uuid := range uuids {
     go func (uuid string) {
       _,err := Foo(uuid)
       if err != nil {
          errMap[uuid] = err
          return
       }
     }(uuid)
   }
   return combineErrors(errMap)
}


func main() {
  uuids := [2]string{"abcd", "efgh"}
  _ = processOrders(uuids[:])
}
