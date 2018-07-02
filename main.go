package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func ShowVersion() {
	fmt.Println("v1.0.0")
}

func main() {
	var help, version bool
	var n int

	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&version, "v", false, "show version")

	flag.IntVar(&n, "n", 1, "number of repeat times")

	flag.Usage = func() {
		fmt.Println()
		fmt.Println("Usage: " + os.Args[0] + " [OPTIONS] COMMAND [ARG...]")
		fmt.Println()
		fmt.Println("Repeat command")
		fmt.Println()
		fmt.Println("Options:")
		flag.CommandLine.PrintDefaults()
		fmt.Println()
	}

	flag.Parse()

	args := flag.Args()

	if help {
		flag.Usage()
		return
	}

	if version {
		ShowVersion()
		return
	}

	if len(args) <= 0 {
		flag.Usage()
		return
	}

	for i := 0; i < n; i++ {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			os.Exit(1)
		}
	}
}
