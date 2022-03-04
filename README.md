
This is the artifact accompanying the PLDI'22 paper titled "A Study of Real-world Data Races in Golang."

#  Getting Started Guide

There are two ways to exercise the artifact: 

A. Using a docker

B. Without using a docker.

# A. Using the docker
1. Download the artifact from zenodo URL.

   `curl -o gorace-examples.tgz https://zenodo.org/record/6329600/files/gorace-examples.tgz`
2. Uncompress the file.
  
   `tar zxf gorace-examples.tgz`
3. `cd gorace-examples`
4. Build the docker image.
    
    `docker build -t golang_datarace_artifact .` 
5. Run the docker image.
  
   `docker run -it golang_datarace_artifact /bin/bash`
6. Run all examples. A `PASS` message means, the data race was correctly found. `WARNING: DATA RACE` message is to be expected.
   
   `bash run.sh`
7. Exit from the docker.
  
   `exit`

# B. Without using the docker
Follow these steps.
1. Install Go `go1.15.5` matching you distribution from here: https://go.dev/dl/
2. Download the artifact from zenodo URL.
  
   `curl -o gorace-examples.tgz https://zenodo.org/record/6329600/files/gorace-examples.tgz`
3. Uncompress the file.
  
   `tar zxf gorace-examples.tgz`
4. `cd gorace-examples`
5. Run all examples. A `PASS` message means, the data race was correctly found. `WARNING: DATA RACE` message is to be expected.
  
   `bash run.sh`


# Step-by-Step Instructions 
#### A list of claims from the paper supported by the artifact, and how/why.

We propose to evaluate the data race patterns shown in the Section 4 (Observations of Data Races in Go) of the paper.
The artifact includes the sample data race examples corresponding to the patterns presented in the paper. The artifact will show that when run with the data race tool (`-race` flag) these races are detected. The evaluators are encouraged to peruse the examples to convince that go language and data race have a complex interplay (more so compared with other languages). 

#### A list of claims from the paper not supported by the artifact, and how/why.

We do not claim to reproduce how frequently these patterns occur because doing so requires releasing a 50 million lines of proprietary business-critical code. We also don't claim to reproduce the concurrency aiming different languages that we observed and showed in Table-1 and Figure 1. We also don't release the automation used for bug filing, which is internal to the company.

## Details of the artifact
The data race examples are associated with a Listing numbers used in the paper. A few additional data race examples are also included, which are mentioned but not shown as listings in the paper.

The last step above (`bash run.sh`) runs all data race examples/patterns present in the artifact.
There are 12 data race patterns one each in `.go` file. 

1. Listing 1 from the paper is present in `loopIndexVariableCapture.go`. This example shows the data race due to loop index variable capture.

2. Listing 2 from the paper is present in `errVariableCapture.go`. This example shows the data race due to `err` variable capture.

3. Listing 3 from the paper is present in `namedReturnVariableCapture.go`. This example shows the data race due to `named-return` variable capture.

4. Listing 4 from the paper is present in `namedReturnVariableCaptureWithDefer.go`. This example shows the data race due to `named-return` variable capture with deferred return.

5. Listing 5 from the paper is present in `racesInSlices.go`. This example shows the data race due to concurrent access to a slice.

6. Listing 6 from the paper is present in `concurrentHashmapAccess.go`. This example shows the data race due to concurrent access to hashmap.

7. Listing 7-8 from the paper is present in `methodInvocationByValueVsByPointer.go`. This example shows the data race due to ambiguous by-value vs. by-pointer syntax in Go.

8. Listing 9 from the paper is present in `mixingMessagePassingWithSharedmem.go`. This example shows the data race due to mixing message passing with shared memory.

9. Listing 10 from the paper is present in `incorrectGroupAdd.go`. This example shows the data race due to incorrect placement of WaitGroup.Add().

10. `incorrectGroupDone.go` shows the data race due to incorrect placement of WaitGroup.Done(). This example is mentioned in the paper but there is no listing in the paper.

11. `prod.go` and `prod_test.go` in the directory `testTable` show the data race due to running parallel tests in test suite/table. This example is mentioned in the paper but there is no listing in the paper.

12. Listing 11 from the paper is present in `rlock.go`. This example shows the data race due to mutating a shared variable in a reader lock held in read-only mode.


`run.sh` will run all these examples in one go. After running the `run.sh` script, for each `.go` file, a corresponding `.log` file will be generated, which captures the data race found by the tool. 

You may run an individual test with the command: `go run -race <go file>`.

The test table example is different and requires `cd` into `testTable` directory and running `go test -race .`

You may investigate the `*.log` files which show the  data race in each example. Each log file contains two the call stacks where the data race is found. There will also be parent call stack of the goroutines if present.
You can open up the corresponding `.go` file to see the source code of the data race.

As an example, `loopIndexVariableCapture.log` might look like this.
```
==================
WARNING: DATA RACE
Read at 0x00c0000a2000 by goroutine 6:
  main.main.func1()
      /Users/mc29/scratch/gorace-examples/loopIndexVariableCapture.go:11 +0x38

Previous write at 0x00c0000a2000 by main goroutine:
  main.main()
      /Users/mc29/scratch/gorace-examples/loopIndexVariableCapture.go:9 +0x5a

Goroutine 6 (running) created at:
  main.main()
      /Users/mc29/scratch/gorace-examples/loopIndexVariableCapture.go:10 +0x84
==================
Found 1 data race(s)
exit status 66
```
The contents of the source file `loopIndexVariableCapture.go` with line annotation is shown below:
```
  1 package main
  2 
  3 func processJob(job int) {
  4         // do something
  5 }
  6 
  7 func main() {
  8         jobs := [2]int{1, 2}
  9         for job := range jobs {
 10                 go func() {
 11                         processJob(job)
 12                 }()
 13         }
 14 }
```
The shared variable `job` is at address `0x00c0000a2000`.
The first call stack shows the concurrent read from this variable on line `11`. The second call stack shows the concurrent write to this variable on line `9`.  
The call stack staring with `Goroutine 6 (running) created at:` is the parent of the concurrent reader, which was created in the main goroutine at line `10`. The writer goroutine does not have a parent since it is in `main`.

The data race here is due to the concurrent modification of `job` on line `9` as it is being read concurrently on line `11` due to the *transparent capture of the variable* by the gorouting.

You may investigate `run.bash` file to add more data race examples to easily extend the script.


