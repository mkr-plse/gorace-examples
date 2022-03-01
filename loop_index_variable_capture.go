package main

func ProcessJob(job int) {
  // do something
}

func main() {
  jobs := [2]int{1,2}
  for job := range jobs {
    go func () {
      ProcessJob(job)
    }()
  }
}
