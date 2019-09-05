package main

import (
	"errors"
	"fmt"
)

var (
	MyError = myError()
)

func myError() error { return errors.New("myErr") }

func simpleError() error {
	return MyError
}

func wrappedError() error {
	err := simpleError()
	return fmt.Errorf("%w", err)
}

func main() {
	err := wrappedError()
	if errors.Is(err, MyError) {
		fmt.Printf(err.Error()) // myErr
	}
}
