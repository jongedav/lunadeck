package main

import (
	"fmt"
	"os"

	"github.com/jongedav/rlayed/internal/display"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: rlayed stop")
		os.Exit(2)
	}

	switch os.Args[1] {
	case "stop":
		if err := display.Stop("/tmp/rlayed-dev"); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case "start":
		if err := display.Start("/tmp/rlayed-dev"); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %q\n", os.Args[1])
		fmt.Fprintln(os.Stderr, "usage: rlayed stop")
		os.Exit(2)
	}
}
