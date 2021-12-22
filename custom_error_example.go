package main

import (
	"errors"
	"fmt"
)

var errAny = fmt.Errorf("any err")

type myError struct {
	code error
	msg  string
}

func (e myError) Error() string {
	return fmt.Sprintf("code: %s, msg: %s", e.code.Error(), e.msg)
}

func (e myError) Is(tar error) bool {
	return errors.Is(e.code, tar)
}

func main() {
	err := wrap()

	if errors.Is(err, errAny) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}

func wrap() error {
	return fmt.Errorf("wrapping error: %w", myError{
		code: errAny,
		msg:  "wing",
	})
}

