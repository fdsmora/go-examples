package main

import "encoding/json"

type MyStruct struct {
	Foo        string `json:"foo_field"`
	Bar        string `json:"bar_field"`
	AnotherFoo string `json:"-"`
}

func (s *MyStruct) UnmarshalJSON(data []byte) error {
	type MS MyStruct
	if err := json.Unmarshal(data, (*MS)(s)); err != nil {
		return err
	}
	s.AnotherFoo = s.Foo
	return nil
}

func MyFunc(s string) MyStruct {
	var result MyStruct
	if err := json.Unmarshal([]byte(s), &result); err != nil {
		panic(err)
	}

	return result
}
