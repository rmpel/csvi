package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/hymkor/csvi/csviapp"
)

func main() {
	banner := fmt.Sprintf("Csvi %s-%s-%s <https://hymkor.github.io/csvi>",
		version, runtime.GOOS, runtime.GOARCH)
	if err := csviapp.RunWithBanner(banner); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
