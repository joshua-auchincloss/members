package main

import (
	"members/cli"
	"os"
)

// "members/cli"
// "os"

func main() {
	if err := cli.BuildApp().Run(os.Args); err != nil {
		panic(err)
	}

	// err := ast.FromFile("../test/pkg/a/test-a.proto").Walk()
	// if err != nil {
	// 	panic(err)
	// }
}
