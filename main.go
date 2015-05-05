package main

import (
	"./dpkg"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Need a version to parse\n")
		return
	}

	versionString := os.Args[1]

	version, err := dpkg.ParseVersion(versionString)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("  %s -> %d // %s // %s\n",
		versionString,
		version.Epoch,
		version.Version,
		version.Revision,
	)
}
