package another

import (
	"fmt"
	"packages-learning/version"
)

// I did this to test that the global variables in an imported package
// are also global with respect to the project scope, so if an importing module A
// updates I, another importing project B sees that update

func another() {
	version.I += 1
	fmt.Printf("another.another() -> version.I = %d\n", version.I)
}

func init() {
	another()
}
