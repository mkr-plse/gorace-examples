package main

func Foo()(string, error) {
   return "",nil
}

func Bar()(string, error) {
   return "",nil
}

func Baz(x string, val bool)(string, error) {
   return "",nil
}


func Redeem() (resp string, err error) {
 defer func () {
   resp,err = Foo()
 }()
 _,err = Bar()
 // err check but no return

 go func () {
   Baz("abcdef", err != nil )
 }()
 return // the defer function runs after here
}

func main() {
  _,_ = Redeem()
}
