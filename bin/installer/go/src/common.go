package src

import (
	"goblitz-installer/helper"
	"io"
	"log"
	"net/http"
	"os"
)

func InstallGoBlitz(version string) {
	checkFile, err := helper.CheckIfPathExists("GoBlitz")
	if err != nil {
		log.Fatalf("Failed to check if GoBlitz directory exists: %s", err)
	}

	if !checkFile {
		log.Println("GoBlitz directory does not exist. Pulling GoBlitz...")
		releaseUri := "https://github.com/LookinLabs/GoBlitz/archive/refs/tags/" + version + ".zip"

		uriResponse, err := http.Get(releaseUri)
		if err != nil {
			log.Fatalf("Failed to download the zip file: %s", err)
		}

		defer uriResponse.Body.Close()

		fileOutput, err := os.Create("/tmp/GoBlitz.zip")
		if err != nil {
			log.Fatalf("Failed to create the zip file: %s", err)
		}

		defer fileOutput.Close()

		_, err = io.Copy(fileOutput, uriResponse.Body)
		if err != nil {
			log.Fatalf("Failed to write the zip file: %s", err)
		}

	}

}
