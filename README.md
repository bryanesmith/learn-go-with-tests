# learn-go-with-tests
Exercises from [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests)

# Useful commands
```
% go run hello.go
% go mod init hello
% go test [-v]
% go test [-v] -bench=.
% go test -cover
% godoc -http=localhost:8000
```

# Notes

* Cannot write tests in package `main`, intended for integrating with other packages only (see "Arrays and slices")
* Arrays include their length as part of type, and are pretty inflexible; you don't see them much in Go code. Instead, use **slices**, which have no specified length
* Slices have fixed capacity, but you can create new slices using `append`
    - Array: `b := [2]string{"Penn", "Teller"}` or `b := [...]string{"Penn", "Teller"}`
    - Slice: `b := []string{"Penn", "Teller"}`
* Three ways to create slices:
    1. Literal `[]string{"foo", "bar"}`
    2. Using `make`
    3. Slicing, e.g., `b[2:]`
* Slices are pointers to an underlying array, not as copies
    - So use `copy` before returning a slice if you want to ensure caller doesn't modify underlying data structure or file!
* Example of using `make([]T, len, cap)` to create a slice, where `cap` is optional capacity:
    ```go
    sums := make([]int, lengthOfNumbers)
    ```
* **Variadic functions** allow you to pass in variable number of parameters: 
    ```go
    func SumAll(numbersToSum ...[]int) []int {
	    return nil
    }
    ```
* Use `reflect.DeepEqual` to compare slices - though warning, it doesn't provide compile-time type safety (e.g., can compare string to slice)

