package main

func Foo()(string, error) {
   return "",nil
}

func Bar()(string, error) {
   return "",nil
}

func Baz()(string, error) {
   return "",nil
}


func main() {
  _,err := Foo()
  if err != nil {
    // do something
  }

  go func () {
    _,err = Bar()
    if err != nil {
     // do something
    }
  }()

  _,err = Baz()
  if err != nil {
    // do something
  }
}
