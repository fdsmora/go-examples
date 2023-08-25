package misc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type anyTestType struct {
	A string
}

var rawData = []byte(`{ "A" : "abc" }`)

func TestLoadObjectPointerType(t *testing.T) {
	t.Run("Ok - is a pointer type", func(t *testing.T) {
		want := &anyTestType{
			A: "abc",
		}
		got := LoadObject[*anyTestType]("testdata/any_datafile.json")
		assert.Equal(t, want, got)
	})
}

func TestLoadObjectValueType(t *testing.T) {
	t.Run("Ok - is a value type", func(t *testing.T) {
		want := anyTestType{
			A: "abc",
		}
		got := LoadObject[anyTestType]("testdata/any_datafile.json")
		assert.Equal(t, want, got)
	})
}

