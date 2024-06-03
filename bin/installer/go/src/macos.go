package src

import (
	"io"
	"log"
	"os/exec"
)

func InstallMacOSUtilities() {
	// List of utilities to install
	commonUtils := []string{"curl", "git", "go"}
	linterUtils := []string{"golangci-lint", "gosec"}
	migrationUtils := []string{"goose"}

	// Check if Homebrew is installed
	if !isCommandAvailable("brew") {
		log.Println("Homebrew is not installed. Installing Homebrew...")
		runCommand("/bin/bash", "-c", `$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)`)
	}

	// Install common utilities
	for _, util := range commonUtils {
		if !isCommandAvailable(util) {
			log.Printf("%s is not installed. Installing %s...\n", util, util)
			runCommand("brew", "install", util)
		}
	}

	// Install linter utilities
	for _, util := range linterUtils {
		if !isCommandAvailable(util) {
			log.Printf("%s is not installed. Installing %s...\n", util, util)
			runCommand("brew", "install", util)
		}
	}

	// Install migration utilities
	for _, util := range migrationUtils {
		if !isCommandAvailable(util) {
			log.Printf("%s is not installed. Installing %s...\n", util, util)
			runCommand("brew", "install", util)
		}
	}

	// Install air
	if !isCommandAvailable("air") {
		log.Println("air is not installed. Installing air...")
		runCommand("/bin/bash", "-c", `$(curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh)`)
	}

	// Check if Docker is installed
	if !isCommandAvailable("docker") {
		log.Println("Docker is not installed. Please install Docker Desktop.")
		// Open the Docker Desktop download page
		runCommand("open", "https://hub.docker.com/editions/community/docker-ce-desktop-mac/")
	}
}

func isCommandAvailable(name string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command -v "+name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func runCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to execute command: %s", err)
	}
}
