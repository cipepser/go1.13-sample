package main

import (
	"errors"
	"fmt"
)

type InvalidChar struct {
	err error
}

func (ic *InvalidChar) Error() string {
	ic.err = errors.New("INVALID CHARACTER")
	return fmt.Errorf("%w", ic.err).Error()
}

type EOF struct {
	err error
}

func (e *EOF) Error() string {
	e.err = errors.New("EOF")
	return fmt.Errorf("%w", e.err).Error()
}

func mustFailParse() error {
	return &InvalidChar{}
}

func wrappedError() error {
	err := mustFailParse()
	return fmt.Errorf("%w", err)
}

func main() {
	err := wrappedError()
	////if ierr, ok := err.(*InvalidChar); ok {
	//var ierr *InvalidChar
	//if errors.As(err, &ierr) {
	//	fmt.Println(ierr) // INVALID CHARACTER
	//}
	fmt.Printf("Type:%T\nValue:%v\n", err, err)

	fmt.Println("---")
	err = errors.Unwrap(err)
	fmt.Printf("Type:%T\nValue:%v\n", err, err)

	fmt.Println("---")
	err = errors.Unwrap(err)
	fmt.Printf("Type:%T\nValue:%v\n", err, err)
}
