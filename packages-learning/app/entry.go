package main

import (
	"fmt"
	_ "packages-learning/another"
	"packages-learning/version"
)

/*
	https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc

	Execution order (numbered):
	0. app/entry.go // module evaluated first as it's the 'main' module
	1. app/entry.go: var myVersion = fetchVersion() // Global scoped variable
	2. app/fetch-version.go // file evaluated but because it imports "packages-learning/version" then...
	3. version/entry.go: var Version = getLocalVersion() // global scoped vars in the module evaluated first
	4. version/entry.go ==> getLocalVersion()
	5. version/get-version.go ==> getVersion()
	6. version/entry.go ==> init()
	7. app/fetch-version.go ==> fetchVersion()
	8. app/entry.go ==> init()
	9. version ===>  1.0.0
*/

func init() {
	fmt.Println("app/entry.go ==> init()")
}

var myVersion = fetchVersion()

func main() {
	fmt.Println("app/fetch-version.go ==> fetchVersion()")
	fmt.Println("version ===> ", myVersion)
	version.I += 1
	fmt.Printf("entry.main() -> version.I = %d\n", version.I)
}
