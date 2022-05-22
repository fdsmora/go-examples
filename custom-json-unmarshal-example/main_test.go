package main

import "testing"

func TestMainFunc(t *testing.T) {
	want := MyStruct{Foo: "foo", Bar: "bar", AnotherFoo: "foo"}
	input := `{
        "foo_field" : "foo",
        "bar_field" : "bar",
        "another_field" : "foo"
    }`
	t.Run("my test main", func(t *testing.T) {
		got := MyFunc(input)
		if want != got {
			t.Errorf("got %#v, want %#v", got, want)
		}
	})
}
