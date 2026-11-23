package main

import (
	"errors"
	"fmt"
	"strconv"
)

// errors are just values in go, there is no try / catch
// error is a built in interface ->  type error interface { Error() string }
// so anything with an Error() string method is an error, that's why MyError in error.go works

// err_not_found -> a sentinel error, a fixed error value we can compare against later
var err_not_found = errors.New("item not found")

// divide -> the most common go pattern, return (result, error), error is nil when everything is fine
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero") // simplest way to make an error
	}
	return a / b, nil
}

// find_item -> returns the sentinel error when the item is missing
func find_item(items map[string]int, name string) (int, error) {
	qty, ok := items[name]
	if !ok {
		return 0, err_not_found
	}
	return qty, nil
}

// buy -> wraps the lower error with extra context using fmt.Errorf and %w
// %w keeps the original error inside, so callers can still check it with errors.Is
func buy(items map[string]int, name string) error {
	if _, err := find_item(items, name); err != nil {
		return fmt.Errorf("buy %q: %w", name, err)
	}
	return nil
}

// errors_patterns -> runs all the error patterns one by one
func errors_patterns() {
	// 1. (value, error) pattern
	if result, err := divide(10, 0); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", result)
	}

	// 2. standard library functions also return errors, e.g. converting text to a number
	if n, err := strconv.Atoi("12a"); err != nil {
		fmt.Println("Atoi error:", err)
	} else {
		fmt.Println("number:", n)
	}

	// 3. wrapping + errors.Is -> checks the error AND everything wrapped inside it
	err := buy(map[string]int{"apple": 3}, "kiwi")
	fmt.Println("wrapped error:", err)
	fmt.Println("errors.Is(err, err_not_found) ?", errors.Is(err, err_not_found))

	// 4. errors.As -> pulls out our custom error type (MyError from error.go) so we can read its fields
	var my_err *MyError
	if errors.As(run(), &my_err) {
		fmt.Println("errors.As found MyError, what =", my_err.what)
	}
}
