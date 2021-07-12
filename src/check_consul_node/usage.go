package main

import (
	"fmt"
	"runtime"
)

func showVersion() {
	fmt.Printf(versionText, name, version, name, runtime.Version())
}

func showUsage() {
	showVersion()
	fmt.Printf(helpText, name, defaultConsulURL)
}
