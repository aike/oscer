package main

import (
	"fmt"
	"os"
	"./osc"
)

func main() {
	err := osc.CheckArg(os.Args)
	if err != nil {
		if err.Error() == "args error" {
			fmt.Fprintf(os.Stderr, "usage: %s host port path [args ...]\n", os.Args[0])
			os.Exit(1)
		} else {
			fmt.Println(err)
			os.Exit(2)
		}
	}
	os.Exit(0)
}

