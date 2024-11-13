// oscer.go by aike
// licenced under MIT License. 

package main

import (
	"fmt"
	"os"
	"oscer/osc"
)

var version string = "1.2"

func main() {
	if osc.IsServer(os.Args) {
		// server
		err := osc.CreateServer(os.Args[2])
		checkError(err)
		os.Exit(0)

	} else {
		// client
		err := osc.CheckArg(os.Args)
		checkError(err)
		osc.Send()
		os.Exit(0)
	}
}

func checkError(err error) {
	if err != nil {
		if err.Error() == "args error" {
			fmt.Fprintf(os.Stderr, "oscer ver %s\n", version)
			fmt.Fprintf(os.Stderr, "usage: oscer host port /osc/address [args ...]\n")
			fmt.Fprintf(os.Stderr, "       oscer receive port\n")
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(2)
		}
	}
}
