// oscer.go by aike
// licenced under MIT License. 

package main

import (
	"fmt"
	"os"
	"./osc"
)

var version string = "1.1"

func main() {
	err := osc.CheckArg(os.Args)
	if err != nil {
		if err.Error() == "args error" {
			fmt.Fprintf(os.Stderr, "oscer ver %s\n", version)
			fmt.Fprintf(os.Stderr, "usage: oscer host port /osc/address [args ...]\n")
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(2)
		}
	}
	osc.Send()
	os.Exit(0)
}

