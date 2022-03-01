#listing 1
echo "----------------"
echo "Demo of loop index variable capture (Listing 1)"
echo "----------------"
go run --race loop_index_variable_capture.go

#listing 2
echo "----------------"
echo "Demo of error variable capture (Listing 2)"
echo "----------------"
go run --race err_variable_capture.go

#listing 3
echo "----------------"
echo "Demo of named return variable capture (Listing 3)"
echo "----------------"
go run --race named_return_variable_capture.go

#listing 4
echo "----------------"
echo "Demo of named return variable capture with deferred return (Listing 4)"
echo "----------------"
go run --race named_return_variable_capture_with_defer_return.go

#listing 5
echo "----------------"
echo "Demo of races in slices (Listing 5)"
echo "----------------"
go run --race races_in_slices.go

#listing 6
echo "----------------"
echo "Demo of races in concurrent access too hashmap (Listing 6)"
echo "----------------"
go run --race concurrent_access_to_hashmap.go

#listing 7
echo "----------------"
echo "Demo of races due to confusion on whether method invocation is by value or pointer (Listing 7)"
echo "----------------"
go run --race method_invocation_by_value_or_pointer.go
