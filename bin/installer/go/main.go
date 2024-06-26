package main

import (
	"flag"
	"goblitz-installer/src"
	"log"
	"os"
	"runtime"
)

var version string
var install bool
var update bool

func init() {
	flag.StringVar(&version, "version", "", "Specify the version to install or update.")
	flag.BoolVar(&install, "install", false, "Install GoBlitz.")
	flag.BoolVar(&update, "update", false, "Update GoBlitz.")
	flag.Parse()
}

func main() {
	if version == "" {
		log.Println("No version specified. Please specify the version to install or update.")
		log.Println("Available versions can be tag or branch name. Example: v1.0.0 or master")
		os.Exit(1)
	}

	switch {
	case install && runtime.GOOS == "darwin":
		src.InstallMacOSUtilities()
		src.InstallGoBlitz(version)
	case update && runtime.GOOS == "darwin":
		// Implement the update functionality here
	default:
		log.Println("Invalid option. Please use --install or --update.")
		os.Exit(1)
	}
}
