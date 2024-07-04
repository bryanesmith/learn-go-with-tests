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
* **Variadic functions** allow you to pass in variable number of parameters: 
    ```go
    func SumAll(numbersToSum ...[]int) []int {
	    return nil
    }
    ```
* Use `reflect.DeepEqual` to compare slices - though warning, it doesn't provide compile-time type safety (e.g., can compare string to slice)
* Using `make` to create a slice:
    ```go
    sums := make([]int, lengthOfNumbers)
    ```
* Slices have fixed capacity, but you can create new slices using `append`