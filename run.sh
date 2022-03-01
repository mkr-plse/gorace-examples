set -e
#listing 1
echo "----------------"
echo "Demo of loop index variable capture (Listing 1)"
echo "----------------"
go run -race loopIndexVariableCapture.go || true

#listing 2
echo "----------------"
echo "Demo of error variable capture (Listing 2)"
echo "----------------"
go run -race errVariableCapture.go || true

#listing 3
echo "----------------"
echo "Demo of named return variable capture (Listing 3)"
echo "----------------"
go run -race namedReturnVariableCapture.go || true

#listing 4
echo "----------------"
echo "Demo of named return variable capture with deferred return (Listing 4)"
echo "----------------"
go run -race namedReturnVariableCaptureWithDefer.go || true

#listing 5
echo "----------------"
echo "Demo of races in slices (Listing 5)"
echo "----------------"
go run -race racesInSlices.go || true

#listing 6
echo "----------------"
echo "Demo of races in concurrent access too hashmap (Listing 6)"
echo "----------------"
go run -race concurrentHashmapAccess.go || true

#listing 7
echo "----------------"
echo "Demo of races due to confusion on whether method invocation is by value or pointer (Listing 7)"
echo "----------------"
go run -race methodInvocationByValueVsByPointer.go || true


#listing 10
echo "----------------"
echo "Demo of races due to incorrect placement of WaitGroup.Add() (Listing 10)"
echo "----------------"
go run -race incorrectGroupAdd.go || true


#listing 10+
echo "----------------"
echo "Demo of races due to incorrect placement of WaitGroup.Done() (not shown but mentioned)"
echo "----------------"
go run -race incorrectGroupDone.go || true


#listing 11
echo "----------------"
echo "Demo of races due to mutating shared variable in a reader lock held in read-only mode "
echo "----------------"
go run -race rlock.go || true


#listing NA
echo "----------------"
echo "Demo of races due to running parallel tests in test suite/table"
echo "----------------"
pushd testTable
go test -race . || true
popd




#mixingMessagePassingWithSharedmem.go
