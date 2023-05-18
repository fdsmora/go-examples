package main

import (
	"encoding/json"
	"fmt"
)

var data = []byte(`{
        "campo1" : null,
        "campoStr" : null
}`)

type A struct {
	Campo    *int    `json:",omitempty"`
	CampoStr *string `json:",omitempty"`
}

func main() {
	a := A{}
	err := json.Unmarshal(data, a)
	if err != nil {
		fmt.Printf("value: %#v", a)
	}
}
