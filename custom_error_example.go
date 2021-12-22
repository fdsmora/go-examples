package main

import (
	"errors"
	"fmt"
)

var errAny = fmt.Errorf("any err")

type Details struct {
	Code error
	Msg  string
}

type myError struct {
	Details
}

func (e myError) Error() string {
	//return fmt.Sprintf("code: %s, msg: %s", e.code.Error(), e.msg)
	return fmt.Sprintf("%#v", e.Details)
}

func (e myError) Is(tar error) bool {
	return errors.Is(e.Code, tar)
}

func main() {
	err := wrap()

	if errors.Is(err, errAny) {
		fmt.Println("true")
		fmt.Println(err)
	} else {
		fmt.Println("false")
	}

}

func wrap() error {
	return fmt.Errorf("wrapping error: %w", myError{
		Details{
			Code: errAny,
			Msg:  "wing",
		},
	})
}

