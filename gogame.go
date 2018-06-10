package main

import (
	"flag"
	"fmt"
)

var (
	Version     string
	BuildCommit string
	BuildTime   string
	GoVersion   string
)

// printVersion print version info from makefile
func printVersion() {
	fmt.Printf(`Version:      %s
Go version:   %s
Git commit:   %s
Built:        %s
`, Version, GoVersion, BuildCommit, BuildTime)
}

func main() {

	version := flag.Bool("v", false, "print version info")
	flag.Parse()

	// print version info only
	if *version {
		printVersion()
		return
	}

	// go on
}
