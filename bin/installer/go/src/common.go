package src

import (
	"goblitz-installer/helper"
	"log"
	"os"
)

func InstallGoBlitz(version string) {
	exists, err := helper.CheckIfFileExists("GoBlitz")
	if err != nil {
		log.Fatalf("Failed to check if GoBlitz directory exists: %s", err)
	}
	if !exists {
		log.Println("GoBlitz directory does not exist. Pulling GoBlitz...")
		runCommand("git", "config", "--global", "advice.detachedHead", "false")
		runCommand("git", "clone", "--branch", version, "git@github.com:LookinLabs/GoBlitz.git")
		os.RemoveAll("GoBlitz/.git")
	}

	err = os.WriteFile(".version.txt", []byte(version), 0644)
	if err != nil {
		log.Fatalf("Failed to write version to .version.txt: %s", err)
	}
}
